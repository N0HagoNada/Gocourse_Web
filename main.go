package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"gocourse/internal/users"
	"log"
	"net/http"
	"time"
)

// Respond XML struct
type Response struct {
	Greeting string
	Names    []string `xml:"Names>Name"`
}

func main() {

	router := mux.NewRouter()

	userEnd := users.MakeEndpoints()
	router.HandleFunc("/users", userEnd.GetAll).Methods(http.MethodGet).Methods(http.MethodGet)
	// router.HandleFunc("/users", userEnd.Get).Methods(http.MethodGet).Methods(http.MethodGet)
	// router.HandleFunc("/users", userEnd.Delete).Methods(http.MethodGet).Methods(http.MethodPost)
	// router.HandleFunc("/users", userEnd.Update).Methods(http.MethodGet).Methods(http.MethodPost)
	router.HandleFunc("/users", userEnd.Create).Methods(http.MethodPost)
	router.HandleFunc("/courses", getCourses).Methods(http.MethodGet)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}

}

func getCourses(w http.ResponseWriter, r *http.Request) {
	response := Response{"Hello", []string{"World", "Courses"}}
	// x, err := xml.MarshalIndent(response, "", " ")
	fmt.Println("got /courses")
	w.Header().Set("Content-Type", "text/xml")
	err := xml.NewEncoder(w).Encode(response)
	fmt.Println(response.Names)
	if err != nil {
		log.Fatal(err)
	}

}
