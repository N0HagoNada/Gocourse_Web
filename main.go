package main

import (
	"github.com/gorilla/mux"
	"gocourse/internal/courses"
	"gocourse/internal/users"
	"gocourse/pkg/bootstrap"
	"log"
	"net/http"
	"time"
)

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
	courseRepo, _ := courses.NewRepo(l, db)
	coursesSrv := courses.NewService(l, courseRepo)
	coursesEnd := courses.MakeEndpoints(coursesSrv)

	router.HandleFunc("/users", userEnd.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userEnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", userEnd.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/users/{id}", userEnd.Update).Methods(http.MethodPatch)
	router.HandleFunc("/users", userEnd.Create).Methods(http.MethodPost)

	router.HandleFunc("/courses", coursesEnd.Create).Methods(http.MethodPost)
	router.HandleFunc("/courses", coursesEnd.GetAll).Methods(http.MethodGet)
	router.HandleFunc("/courses/{id}", coursesEnd.Get).Methods(http.MethodGet)
	router.HandleFunc("/courses/{id}", coursesEnd.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/courses/{id}", coursesEnd.Update).Methods(http.MethodPatch)

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
