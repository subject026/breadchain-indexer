package router

import (
	"log"
	"net/http"
)

func (apiCfg apiConfig) handlerGetProjects(w http.ResponseWriter, r *http.Request) {
	dbProjects, err := apiCfg.DB.GetProjects(r.Context())
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	respondWithJSON(w, http.StatusOK, databaseProjectsToProjects(dbProjects))
}
