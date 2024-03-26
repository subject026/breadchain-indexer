package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
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

	scheduler := gocron.NewScheduler(time.UTC)

	mainCtx := context.Background()

	lastVote := time.Now()

	scheduler.Every(10).Seconds().WaitForSchedule().Do(func() {
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

		// diff := time.Now().Sub(stamp)

		lastVote = time.Now()

	})

	scheduler.Every(1).Second().WaitForSchedule().Do(func() {
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

	scheduler.StartAsync()
}

func generateValue(pointsRemaining uint32) int32 {
	return (int32(rand.Float64() * float64(pointsRemaining)))
}

// func roundToFullMinute(fullTime time.Time) {
// 	fmt.Println("Original time:", fullTime)

// 	// Round to nearest minute
// 	seconds := fullTime.Second()
// 	nanoseconds := fullTime.Nanosecond()
// 	var round time.Duration
// 	if seconds < 30 {
// 		// Round down
// 		round = time.Duration(-seconds)*time.Second - time.Duration(nanoseconds)*time.Nanosecond
// 	} else {
// 		// Round up
// 		round = time.Duration(60-seconds)*time.Second - time.Duration(nanoseconds)*time.Nanosecond
// 	}
// 	roundedTime := time.Add(round)
// 	fmt.Println("Rounded time:", roundedTime)
// }
