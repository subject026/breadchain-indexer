package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/google/uuid"
	"github.com/subject026/breadchain-indexer/internal/database"
)

type VoterState struct {
	Projects []database.Project
	Users    []database.User
}

func startVoter(DB *database.Queries) {

	VOTE_INTERVAL := os.Getenv("VOTE_INTERVAL")
	if VOTE_INTERVAL == "" {
		log.Fatal("VOTE_INTERVAL environment variable not set")
	}

	SLICE_INTERVAL := os.Getenv("SLICE_INTERVAL")
	if SLICE_INTERVAL == "" {
		log.Fatal("SLICE_INTERVAl environment variable not set")
	}

	scheduler := gocron.NewScheduler(time.UTC)

	mainCtx := context.Background()

	lastVote := time.Now()

	voteInterval, err := strconv.ParseInt(VOTE_INTERVAL, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	sliceInterval, err := strconv.ParseInt(SLICE_INTERVAL, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// voters
	scheduler.Every(int(voteInterval)).Second().WaitForSchedule().Do(func() {
		fmt.Println("voter running job!")
		projects, err := DB.GetProjects(mainCtx)
		if err != nil {
			log.Fatal(err)
			return
		}

		users, err := DB.GetUsers(mainCtx)
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println(("lastVote: "), lastVote)

		for _, user := range users {
			// find user that hasn't voted yet
			votes, err := DB.GetVotesInRangeForUser(mainCtx, database.GetVotesInRangeForUserParams{
				UserID:      user.ID,
				CreatedAt:   lastVote,
				CreatedAt_2: time.Now(),
			})
			if err != nil {
				log.Fatal(err)
				return
			}
			if len(votes) > 0 {
				continue
			}
			votesCast := 0

			for i, project := range projects {

				votesRemaining := uint32(100) - uint32(votesCast)

				value := generateValue(votesRemaining)

				if i == len(projects)-1 {
					value = int32(votesRemaining)
				}

				votesCast += int(value)

				fmt.Println(project.Name, " : ", value)

				_, voteErr := DB.CreateVote(mainCtx, database.CreateVoteParams{
					ID:        uuid.New(),
					CreatedAt: time.Now(),
					UserID:    user.ID,
					ProjectID: project.ID,
					Value:     value,
				})

				if voteErr != nil {
					log.Fatal(err)
					return
				}
			}
			break
		}
	})

	// slicer
	scheduler.Every(int(sliceInterval)).Seconds().WaitForSchedule().Do(func() {
		_, err := DB.GetVotesInRange(mainCtx, database.GetVotesInRangeParams{
			CreatedAt:   lastVote,
			CreatedAt_2: time.Now(),
		})
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("lastVote: ", lastVote)

		DB.CreateSlice(mainCtx, database.CreateSliceParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			StartedAt: lastVote,
		})

		lastVote = time.Now()

	})

	scheduler.StartAsync()
}

func generateValue(pointsRemaining uint32) int32 {
	return (int32(rand.Float64() * float64(pointsRemaining)))
}
