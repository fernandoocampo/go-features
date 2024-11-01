package routes

import (
	"encoding/json"
	"io"
	"net/http"
	"strconv"
	"strings"
)

type Pet struct {
	ID   string `json:"pet_id"`
	Name string `json:"pet_name"`
}

func NewRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /pets", func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var newPet Pet
		err = json.Unmarshal(bodyBytes, &newPet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		newPet.ID = strconv.Itoa(len(newPet.Name))

		err = json.NewEncoder(w).Encode(newPet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("PUT /pets", func(w http.ResponseWriter, r *http.Request) {
		bodyBytes, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		var updatedPet Pet
		err = json.Unmarshal(bodyBytes, &updatedPet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = json.NewEncoder(w).Encode(updatedPet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	mux.HandleFunc("/pets", func(w http.ResponseWriter, r *http.Request) {
		pet := Pet{
			ID:   "5",
			Name: "Drila",
		}

		err := json.NewEncoder(w).Encode(pet)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	mux.HandleFunc("/pets/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := r.PathValue("id")
		if id == "404" {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		idNumber, err := strconv.Atoi(id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		pet := Pet{
			ID:   id,
			Name: strings.Repeat("d", idNumber),
		}

		if err := json.NewEncoder(w).Encode(pet); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}
	})

	return mux
}
