package service

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/willive/scoreservice/data"
	"github.com/willive/scoreservice/types"
	"golang.org/x/net/html"
)

func GetScore(db data.Repository, requestedUrl string) (*[]types.Score, error) {
	if requestedUrl != "" {
		_, err := url.ParseRequestURI(requestedUrl)
		if err != nil {
			return nil, err
		}
		log.Printf("Getting %s", requestedUrl)
		resp, err := http.Get(requestedUrl)
		if err != nil {
			return nil, err
		}

		log.Printf("Calculating score for %s", requestedUrl)
		score := calculateScore(resp.Body)
		err = resp.Body.Close()
		if err != nil {
			return nil, err
		}

		log.Printf("%s scored with a value of %d \n", requestedUrl, score)

		t := time.Now()
		ts := t.Format("2006-01-02 15:04:05")
		record := types.Score{Score: score, FileName: requestedUrl, Time: ts}
		db.InsertScore(&record)
		response := []types.Score{record}
		return &response, nil
	}
	log.Println("Getting all scores")
	response := db.GetAllScores()
	return response, nil
}

func calculateScore(r io.Reader) int {
	var score int
	z := html.NewTokenizer(r)
	for {
		tt := z.Next()

		switch {
		case tt == html.ErrorToken:
			return score
		case tt == html.StartTagToken || tt == html.SelfClosingTagToken:
			t := z.Token()
			score += types.Tags[strings.ToLower(t.Data)]
		}
	}
}
