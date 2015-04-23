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

// MySQL uses the ? variant shown above
// PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
// SQLite accepts both ? and $1 syntax
// Oracle uses a :name syntax
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