package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	neon "github.com/stationapi/station-search/db"
)

type response struct {
	Websites []neon.Website
}

func main() {
	db, err := neon.Connect()

	if err != nil {
		log.Fatal(err)
	}

	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.RealIP)

	r.Get("/search", func(w http.ResponseWriter, r *http.Request) {
		search := r.URL.Query().Get("search")

		websites := neon.SearchWebsites(search, db)

		res := response{
			Websites: websites,
		}

		stringified, err := json.Marshal(res)

		if err != nil {
			http.Error(w, "there was an error searching for websites", http.StatusInternalServerError)

			return
		}

		w.WriteHeader(200)
		w.Write(stringified)
	})
}
