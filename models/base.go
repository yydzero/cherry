package models

import (
	"github.com/jmoiron/sqlx"
	"fmt"
	"os"
	"log"

	_ "github.com/lib/pq"
)

var db *sqlx.DB = nil
var err error

func GetDB() *sqlx.DB {
	if db == nil {
		url := fmt.Sprintf("postgres://%s@localhost:5432/cherry?sslmode=disable", os.Getenv("USER"));
		log.Println(url)
		if db, err = sqlx.Open("postgres", url); err != nil {
			panic("could not connect to db: " + err.Error())
		}
	}
	return db
}