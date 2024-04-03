package voter

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

func Start(DB *database.Queries) {

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
		projects, err := DB.GetProjects(mainCtx)
		if err != nil {
			log.Fatal(err)
			return
		}

		votes, err := DB.GetVotesInRange(mainCtx, database.GetVotesInRangeParams{
			CreatedAt:   lastVote,
			CreatedAt_2: time.Now(),
		})
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Creating slice...")
		slice, err := DB.CreateSlice(mainCtx, database.CreateSliceParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			StartedAt: lastVote,
		})
		if err != nil {
			log.Fatal(err)
			return
		}

		fmt.Println("Slice created!!!: ", slice.ID)

		for _, project := range projects {
			projectVotes := filter(votes, func(vote database.GetVotesInRangeRow) bool {
				return vote.ProjectID == project.ID
			})

			value := int32(0)

			for _, vote := range projectVotes {
				value += vote.Value
			}

			fmt.Println("vote count: ", value)

			sliceProject, err := DB.CreateSliceProject(mainCtx, database.CreateSliceProjectParams{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				SliceID:   slice.ID,
				ProjectID: project.ID,
				Value:     value,
			})
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println("created sliceProject: ", sliceProject.ID)
		}

		lastVote = time.Now()

	})

	scheduler.StartAsync()
}

func generateValue(pointsRemaining uint32) int32 {
	return (int32(rand.Float64() * float64(pointsRemaining)))
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
