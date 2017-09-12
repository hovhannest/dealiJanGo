package API_v1

import (
	"net/http"
	"github.com/go-chi/chi"
)

// A completely separate router for api routes
func ApiRouter() http.Handler {
	r := chi.NewRouter()

	r.Get("/", func (w http.ResponseWriter, r *http.Request){
		w.Write([]byte("help is here"))
	})

	return  r
}