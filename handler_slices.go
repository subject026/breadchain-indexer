package main

import (
	"fmt"
	"log"
	"net/http"
)

func (apiCfg apiConfig) handlerGetSlices(w http.ResponseWriter, r *http.Request) {

	dbSlices, err := apiCfg.DB.GetSlices(r.Context())

	fmt.Println("db slices length: ", dbSlices)

	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, databaseSlicesToSlices(dbSlices))
}
