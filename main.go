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
	route        = "/"
	dbUser       = os.Getenv("DB_USER")
	userPassword = os.Getenv("DB_PASSWORD")
	dbhost       = os.Getenv("DB_HOST")
	dbport       = os.Getenv("DB_PORT")
	dbName       = os.Getenv("DB_NAME")
	serviceAddr  = "0.0.0.0:3001"
)

func main() {
	log.Println("Starting scorer service")
	log.Printf("CONNECTING TO DATABASE %s @ %s:%s AS %s", dbName, dbhost, dbport, dbUser)

	db := data.CreateMySQLInstance(dbUser, userPassword, dbhost, dbport, dbName)
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
