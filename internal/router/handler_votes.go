package router

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/subject026/breadchain-indexer/internal/database"
)

type VoteParams struct {
	UserAddress    string `json:"userAddress"`
	ProjectAddress string `json:"projectAddress"`
	Value          int32  `json:"value"`
}

func (apiCfg *apiConfig) handlerCreateVote(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Votes []VoteParams `json:"votes"`
	}
	// b, err := io.ReadAll(r.Body)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// raw := json.RawMessage(string(b))
	// params := parameters{}

	// err2 := json.Unmarshal(raw, &params)
	// if err2 != nil {
	// 	fmt.Println("error unmarshalling in handlerCreateVote", err2)
	// 	respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err2))
	// 	return
	// }
	fmt.Println("CREATE VOTE HANDLER !!!")

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		fmt.Println("error decoding in handlerCreateVote", err)
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err))
		return
	}

	fmt.Println("whats going on ", params.Votes)

	for _, vote := range params.Votes {
		// value, err := strconv.ParseInt(vote.Value, 10, 32)
		// if err != nil {
		// 	respondWithError(w, 400, fmt.Sprint("Error parsing value: ", err))
		// 	return
		// }
		dbProject, err := apiCfg.DB.GetProjectByAddress(r.Context(), vote.ProjectAddress)
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error getting project: ", err))
			return
		}
		dbUser, err := apiCfg.DB.GetUserByAddress(r.Context(), user.WalletAddress)

		fmt.Println("whats going on ")
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error getting user: ", err))
			return
		}

		woo, voteErr := apiCfg.DB.CreateVote(r.Context(), database.CreateVoteParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UserID:    dbUser.ID,
			ProjectID: dbProject.ID,
			Value:     vote.Value,
		})

		fmt.Println("vote created", woo)
		if voteErr != nil {
			log.Fatal(err)
			return
		}

	}

	respondWithJSON(w, http.StatusCreated, nil)
}

func (apiCfg apiConfig) handlerGetVotes(w http.ResponseWriter, r *http.Request) {

	searchFrom := getLastSliceSearchFrom(r, apiCfg)

	dbVotes, err := apiCfg.DB.GetVotesInRange(r.Context(), database.GetVotesInRangeParams{
		CreatedAt:   searchFrom,
		CreatedAt_2: time.Now(),
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, databaseVotesToVotes(dbVotes))
}

func getLastSliceSearchFrom(r *http.Request, apiCfg apiConfig) time.Time {
	lastSlice, err := apiCfg.DB.GetLastSlice(r.Context())
	if err != nil {
		return time.Now().AddDate(-10, 0, 0)
	}
	return lastSlice.CreatedAt
}
