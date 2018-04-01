package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/willive/scoreservice/data"
)

var (
	route        = "/scorer"
	dbUser       = os.Getenv("DB_USER")
	userPassword = os.Getenv("DB_PASSWORD")
	dbhost       = os.Getenv("DB_HOST")
	dbport       = os.Getenv("DB_PORT")
	databaseName = os.Getenv("DB_NAME")
	serviceAddr  = "0.0.0.0:3001"
)

func main() {
	log.Println("Starting scorer service")
	db := data.CreateMySQLInstance(dbUser, userPassword, dbhost, dbport, databaseName)
	defer db.CloseDb()

	r := mux.NewRouter()
	r.HandleFunc(route, HandleGetScore(db)).Methods("GET")

	srv := &http.Server{
		Addr:         serviceAddr,
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      r,
	}
	log.Printf("Scorer service started, listening on %s ", serviceAddr)
	log.Fatal(srv.ListenAndServe())
}
