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

func Must(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// Parse request body with format:
// application/x-www-form-urlencoded
// NOTE: Will only grab the first value for each key
func parseRequestBody(r *http.Request) (map[string]string, error) {
	Must(r.Method == http.MethodPost ||
		r.Method == http.MethodPut ||
		r.Method == http.MethodPatch,
		"Request must be of type Post, Put or Patch",
	)

	data := make(map[string]string)
	err := r.ParseForm()
	for k, v := range r.PostForm {
		if len(v) > 0 {
			data[k] = v[0]
		}
	}
	return data, err

	// contentType := r.Header.Get("Content-type")
	// switch contentType {
	// case "":
	// case "application/x-www-form-urlencoded":
	// case "application/json":
	// default:
	// 	panic("not implemented")
	// }
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
		params := mux.Vars(r)
		id := params["id"]
		Must(id != "",
			"IsPalindromeHandler must be provided with an id in the route")

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
		params := mux.Vars(r)
		id := params["id"]
		Must(id != "",
			"UserGetHandler must be provided with an id in the route")

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
		params, err := parseRequestBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		firstName := params["firstname"]
		lastName := params["lastname"]

		if firstName == "" || lastName == "" {
			msg := fmt.Sprintf("Missing parameters: firstname=%v&lastname=%v", firstName, lastName)
			http.Error(w, msg, http.StatusBadRequest)
			return
		}

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

		id := mux.Vars(r)["id"]
		Must(id != "",
			"UserPutHandler must be provided with an id in the route")

		params, err := parseRequestBody(r)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

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
		params := mux.Vars(r)
		id := params["id"]
		Must(id != "",
			"IsPalindromeHandler must be provided with an id in the route")

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
	apiRouter.HandleFunc("/user", UserPostHandler(us)).Methods(http.MethodPost)
	apiRouter.HandleFunc("/user/{id}", UserPutHandler(us)).Methods(http.MethodPut)
	apiRouter.HandleFunc("/user/{id}", UserDeleteHandler(us)).Methods(http.MethodDelete)

	apiRouter.HandleFunc("/user/{id}/ispalindrome", IsPalindromeHandler(us)).Methods(http.MethodGet)

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "hello there!")
		fmt.Fprint(os.Stdout, "hello there!\n")
	})

	return router
}
