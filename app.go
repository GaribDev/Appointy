package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func SheduleMeeteingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func GetMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func ListMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func ParticipantMeetingEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Not implemented yet")
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/meetings", SheduleMeeteingEndPoint).Methods("GET")
	r.HandleFunc("/meetings", GetMeetingEndPoint).Methods("POST")
	r.HandleFunc("/meetings", ListMeetingEndPoint).Methods("PUT")
	r.HandleFunc("/meetings/{id}", ParticipantMeetingEndPoint).Methods("GET")
	if err := http.ListenAndServe(":27017", r); err != nil {
		log.Fatal(err)
	}
}
