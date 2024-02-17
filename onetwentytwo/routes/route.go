package routes

import (
	"fmt"
	"net/http"
)

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /pets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "post a pet")
	})

	mux.HandleFunc("PUT /pets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "put a pet")
	})

	mux.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "get pets")
	})

	mux.HandleFunc("/pets/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		fmt.Fprintf(w, "get pet: %s", id)
	})

	return mux
}
