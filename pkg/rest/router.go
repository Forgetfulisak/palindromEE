package rest

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"palindromee/pkg/userservice"

	"github.com/gorilla/mux"
)

func parseBody(r *http.Request) (map[string]string, error) {

	urlParams := mux.Vars(r)
	if len(urlParams) != 0 {
		return urlParams, nil
	}

	bodyParams := make(map[string]string)
	err := json.NewDecoder(r.Body).Decode(&bodyParams)
	return bodyParams, err
}

func respond(w http.ResponseWriter, data any) {
	w.Header().Add("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func IsPalindromeHandler(us userservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := params["id"]
		// TODO: handle missing id
		result, err := us.CheckPalindrome(id)
		if err != nil {
			// Check the type of error
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		respond(w, result)
	}
}

// Get user.
// If no ID is provided, returns all users
func UserGetHandler(us userservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := params["id"]
		// TODO: handle missing id
		user, err := us.GetUser(id)
		if err != nil {
			// Check the type of error
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		respond(w, user)
	}
}

// Create user.
func UserPostHandler(us userservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		firstName := params["firstname"]
		lastName := params["lastname"]
		// TODO: handle missing params
		user, err := us.CreateUser(firstName, lastName)
		if err != nil {
			// Check the type of error
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		respond(w, user)
	}
}

// Update user.
// Does not create user if it does not exist.
func UserPutHandler(us userservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := params["id"]
		firstName := params["firstname"]
		lastName := params["lastname"]
		// TODO: handle missing params
		user, err := us.UpdateUser(id, firstName, lastName)
		if err != nil {
			// Check the type of error
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		respond(w, user)
	}

}

// Deletes user.
func UserDeleteHandler(us userservice.Service) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params, err := parseBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		id := params["id"]
		// TODO: handle missing id
		user, err := us.DeleteUser(id)
		if err != nil {
			// Check the type of error
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		respond(w, user)
	}
}

func Handle(us userservice.Service) *mux.Router {

	router := mux.NewRouter()

	router.Use(func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("Incomming request:", r.URL)
			h.ServeHTTP(w, r)
		})
	})

	apiRouter := router.PathPrefix("/api").Subrouter()
	apiRouter.HandleFunc("/user/{id}", UserGetHandler(us)).Methods(http.MethodGet)
	apiRouter.HandleFunc("/user/{id}", UserPostHandler(us)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/user/{id}", UserPutHandler(us)).Methods(http.MethodPut)
	apiRouter.HandleFunc("/user/{id}", UserDeleteHandler(us)).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/user/{id}/ispalindrome", IsPalindromeHandler(us)).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello there!")
		fmt.Fprint(os.Stdout, "hello there!\n")
	})

	return router
}
