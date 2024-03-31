package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/subject026/breadchain-indexer/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		type parameters struct {
			UserAddress string `json:"user_address"`
		}
		decoder := json.NewDecoder(r.Body)
		params := parameters{}
		err := decoder.Decode(&params)
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err))
			return
		}

		user, err := apiConfig.DB.GetUserByAddress(r.Context(), params.UserAddress)

		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error getting user: ", err))
			return
		}

		handler(w, r, user)
	}
}
