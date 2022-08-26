package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func UserPostHandler() {

}
func UserPutHandler() {

}
func UserGetHandler() {

}
func UserDeleteHandler() {

}

func main() {

	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello there!")
		fmt.Fprint(os.Stdout, "hello there!\n")
	})

	apiRoutes := r.Path("/api").Subrouter()

	_ = apiRoutes
	authRoutes := r.Path("/auth").Subrouter()

	authRoutes.Use(func(h http.Handler) http.Handler {
		return h
	})

	srv := &http.Server{
		Addr:         "localhost:8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("Serving on...", srv.Addr)
	log.Fatalln(srv.ListenAndServe())
}
