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

func startVoter(DB *database.Queries) {

	scheduler := gocron.NewScheduler(time.UTC)

	mainCtx := context.Background()

	stamp := time.Now()

	// scheduler.Every(1).Minute().WaitForSchedule().Do(func() {
	// 	votes, err := DB.GetVotesInRange(mainCtx, database.GetVotesInRangeParams{
	// 		CreatedAt:   stamp,
	// 		CreatedAt_2: time.Now(),
	// 	})
	// 	if err != nil {
	// 		log.Fatal(err)
	// 		return
	// 	}

	// 	// diff := time.Now().Sub(stamp)

	// 	stamp = time.Now()

	// })

	scheduler.Every(1).Second().WaitForSchedule().Do(func() {
		fmt.Println("..........................................................")
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

		fmt.Println(("stamp: "), time.Since(stamp))

		for _, user := range users {
			// find user that hasn't voted yet
			votes, err := DB.GetVotesInRangeForUser(mainCtx, database.GetVotesInRangeForUserParams{
				UserID:      user.ID,
				CreatedAt:   time.Now(),
				CreatedAt_2: time.Now(),
			})
			if err != nil {
				log.Fatal(err)
				return
			}
			if len(votes) > 0 {
				continue
			}

			// fmt.Println("***************************")
			// fmt.Println("len votes: ", len(votes))
			// fmt.Println("we have a voter!!")
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
