package main

import (
	"encoding/xml"
	"fmt"
	"github.com/gorilla/mux"
	"gocourse/internal/users"
	"gocourse/pkg/bootstrap"
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

	l := bootstrap.InitLogger()
	db, err := bootstrap.DBConnection()
	if err != nil {
		l.Println(err)
	}
	userRepo, err := users.NewRepo(l, db)
	if err != nil {
		log.Fatal(err)
	}
	userSrv := users.NewService(l, userRepo)
	userEnd := users.MakeEndpoints(userSrv)
	router.HandleFunc("/users", userEnd.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userEnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", userEnd.Update).Methods(http.MethodPatch)
	router.HandleFunc("/users", userEnd.Create).Methods(http.MethodPost)
	router.HandleFunc("/courses", getCourses).Methods(http.MethodGet)

	server := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:3333",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	if err := server.ListenAndServe(); err != nil {
		l.Fatal(err)
	}

}

func getCourses(w http.ResponseWriter, r *http.Request) {
	response := Response{"Hello", []string{"World", "Courses"}}
	// x, err := xml.MarshalIndent(response, "", " ")
	fmt.Println("got /courses")
	w.Header().Set("Content-Type", "application/xml")
	err := xml.NewEncoder(w).Encode(response)
	fmt.Println(response.Names)
	if err != nil {
		log.Fatal(err)
	}

}
