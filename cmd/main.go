package main

import (
	"log"
	"net/http"
	"palindromee/pkg/rest"
	"palindromee/pkg/storage"
	"palindromee/pkg/userservice"
	"time"
)

func main() {

	repo, err := storage.New()
	if err != nil {
		log.Fatalln(err)
	}

	userService := userservice.New(repo)
	router := rest.Handle(userService)

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      router,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Println("Serving on:", srv.Addr)
	log.Fatalln(srv.ListenAndServe())
}
