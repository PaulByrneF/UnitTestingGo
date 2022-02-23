package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func Add(val1, val2 int) int {
	return val1 + val2
}

func Subtract(val1, val2 int) int {
	return val1 - val2
}

func RootEndpoint(response http.ResponseWriter, request *http.Request) {
	response.WriteHeader(200)
	response.Write([]byte("Hello World"))
}
func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	log.Fatal(http.ListenAndServe(":9001", router))
}
