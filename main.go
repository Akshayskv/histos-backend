package main

import (
	"histos-backend/database"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	psql := database.DatabaseConnection{
		ConnectionString: "string",
	}

	err := psql.Connect()

	if err != nil {
		log.Fatal("database connection failed")
	}

	log.Println("server started and listening on port 8080")
	server.ListenAndServe()

}
