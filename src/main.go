package main

import (
	"net/http"
	"github.com/go-chi/chi"
	"./API_v1"
)

func main() {
	r := chi.NewRouter()

	// Mount the api sub-router
	r.Mount("/api/v1", API_v1.ApiRouter())

	http.ListenAndServe(":8080", r)
}
