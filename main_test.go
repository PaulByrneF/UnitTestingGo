package main

import (
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAdd(t *testing.T) {
	total := Add(1, 2)
	assert.NotNil(t, total, "Expecting `total` to not be nil")
	assert.Equal(t, 3, total, "Expected the total to be 3")
}

func TestSubract(t *testing.T) {
	total := Subtract(1, 3)
	assert.NotNil(t, total, "The total not be nil")
	assert.Equal(t, -2, total, "Expected value -2")
}

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/", RootEndpoint).Methods("GET")
	return router
}

func TestRootEndpoint(t *testing.T) {
	request, _ := http.NewRequest("GET", "/", nil)
	response := httptest.NewRecorder()
	Router().ServeHTTP(response, request)
	assert.Equal(t, 200, response.Code, "200 OK - response is expected")
}
