package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/subject026/breadchain-indexer/internal/database"
)

func (apiCfg *apiConfig) handlerCreateVote(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		ProjectAddress string `json:"project_address"`
		Value          string `json:"value"`
	}
	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err))
		return
	}
	value, err := strconv.ParseInt(params.Value, 10, 32)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing value: ", err))
		return
	}

	project, err := apiCfg.DB.GetProjectByAddress(r.Context(), params.ProjectAddress)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error getting project: ", err))
		return
	}

	vote, err := apiCfg.DB.CreateVote(r.Context(), database.CreateVoteParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UserID:    user.ID,
		ProjectID: project.ID,
		Value:     int32(value),
	})
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, databaseVoteToVote(vote))
}

func (apiCfg apiConfig) handlerGetVotes(w http.ResponseWriter, r *http.Request) {
	dbVotes, err := apiCfg.DB.GetVotesInRange(r.Context(), database.GetVotesInRangeParams{
		CreatedAt:   time.Now().Add(-time.Second * 20),
		CreatedAt_2: time.Now(),
	})

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, databaseVotesToVotes(dbVotes))
}
