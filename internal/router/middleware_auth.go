package router

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/subject026/breadchain-indexer/internal/database"
)

type authedHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiConfig *apiConfig) middlewareAuth(handler authedHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		address, err := getUserWallet(r.Header)
		if err != nil {
			fmt.Println("error decoding in middlewareAuth", err)
			respondWithError(w, 400, fmt.Sprint("Error parsing JSON: ", err))
			return
		}
		fmt.Println("\naddress", address)
		user, err := apiConfig.DB.GetUserByAddress(r.Context(), address)

		fmt.Println("rr", err)
		if err != nil {
			respondWithError(w, 400, fmt.Sprint("Error getting user: ", err))
			return
		}

		handler(w, r, user)
	}
}

func getUserWallet(headers http.Header) (string, error) {
	val := headers.Get("Authorization")
	if val == "" {
		return "", errors.New("nothing found in authorization header")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("invalid authorization header")
	}

	if vals[0] != "Address" {
		return "", errors.New("invalid authorization header key")
	}

	return vals[1], nil
}
