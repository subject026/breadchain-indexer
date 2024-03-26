package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/subject026/breadchain-indexer/internal/database"
)

func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		WalletAddress string `json:"wallet_address"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:            uuid.New(),
		CreatedAt:     time.Now(),
		WalletAddress: params.WalletAddress,
	})

	if err != nil {
		respondWithError(w, 400, fmt.Sprint("Error creating user: ", err))
		return
	}
	respondWithJSON(w, 201, databaseUserToUser((user)))
}

func (apiConfig *apiConfig) handlerGetUser(w http.ResponseWriter, r *http.Request, user database.User) {

	respondWithJSON(w, 200, databaseUserToUser(user))
}
