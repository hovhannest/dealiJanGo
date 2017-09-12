package main

import (
	"net/http"
	"github.com/go-chi/chi"
)

// A completely separate router for api routes
func adminRouter() http.Handler {
	r := chi.NewRouter()

	return  r
}