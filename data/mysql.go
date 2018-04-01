package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/willive/scoreservice/types"
)

type Repository interface {
	InsertScore(*types.Score)
	GetAllScores() *[]types.Score
	CloseDb()
}

type dbhandler struct {
	db *sql.DB
}

// CreateMySQLInstance creates a connection to the database
func CreateMySQLInstance(user, password, host, port, dbname string) Repository {
	h := &dbhandler{}
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbname))
	if err != nil {
		log.Fatal(err)
	}
	h.db = db
	return h
}

func (h *dbhandler) InsertScore(score *types.Score) {
	_, err := h.db.Query("INSERT INTO scores (file_name, seo_score, time_stamp) VALUES (?, ?, ?);", score.FileName, score.Score, score.Time)
	if err != nil {
		log.Fatal(err)
	}
}

func (h *dbhandler) GetAllScores() *[]types.Score {
	results := []types.Score{}
	result := types.Score{}
	rows, err := h.db.Query("SELECT file_name, seo_score, time_stamp FROM scores")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&result.FileName, &result.Score, &result.Time)
		if err != nil {
			log.Fatal(err)
		}
		results = append(results, result)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return &results
}

func (h *dbhandler) CloseDb() {
	h.db.Close()
}
