package main

import (
	"log"
	"net/http"
	"os"
	"palindromee/pkg/rest"
	"palindromee/pkg/storage"
	"palindromee/pkg/userservice"
	"time"
)

func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DATABASE_PATH")
	if dbPath == "" {
		dbPath = "./db"
	}

	repo, err := storage.New(dbPath)
	if err != nil {
		log.Fatalln(err)
	}

	userService := userservice.New(repo)
	router := rest.Handle(userService)

	srv := &http.Server{
		Addr:         ":" + port,
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Serving on:", srv.Addr)
	log.Fatalln(srv.ListenAndServe())
}
