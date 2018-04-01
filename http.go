package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/willive/scoreservice/data"
	"github.com/willive/scoreservice/service"
)

func HandleGetScore(repo data.Repository) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		url := params.Get("url")

		score, err := service.GetScore(repo, url)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		body, err := json.Marshal(score)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(body)
	}
}
