package router

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/subject026/breadchain-indexer/internal/database"
)

type apiConfig struct {
	DB *database.Queries
}

func New(DB *database.Queries) *chi.Mux {

	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	apiCfg := apiConfig{
		DB: DB,
	}

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlerReadiness)
	v1Router.Get("/err", handlerError)

	v1Router.Post("/users", apiCfg.handlerCreateUser)
	v1Router.Get("/users", apiCfg.middlewareAuth(apiCfg.handlerGetUser))

	v1Router.Get("/projects", apiCfg.handlerGetProjects)

	v1Router.Post("/votes", apiCfg.middlewareAuth(apiCfg.handlerCreateVote))
	v1Router.Get("/votes", apiCfg.handlerGetVotes)

	v1Router.Get("/slices", apiCfg.handlerGetSlices)

	router.Mount("/v1", v1Router)

	return router
}
