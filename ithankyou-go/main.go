package main

import (
	"github.com/gorilla/mux"
	"iThankYou/ithankyou-go/service"
	"log"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	r.Handle("/favicon.ico", http.NotFoundHandler())
	r.HandleFunc("/user/{email}", service.GetUserByEmail).Methods("GET")
	r.HandleFunc("/user", service.CreateUser).Methods("POST")

	err := http.ListenAndServe(":8081", r)
	if err != nil {
		log.Fatalln(err)
	}
}
