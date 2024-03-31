package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/subject026/breadchain-indexer/internal/database"
	"github.com/subject026/breadchain-indexer/internal/router"
	"github.com/subject026/breadchain-indexer/internal/voter"

	_ "github.com/lib/pq"
)

func main() {

	godotenv.Load(".env")

	portString := os.Getenv(("PORT"))

	if portString == "" {
		log.Fatal("PORT environment variable not set")
	}

	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL environment variable not set")
	}

	conn, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Println(err)
		log.Fatal("Cant connect to database")
	}

	database := database.New(conn)

	router := router.New(database)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	go voter.Start(database)

	srvErr := srv.ListenAndServe()
	if srvErr != nil {
		log.Fatal(err)
	}

}
