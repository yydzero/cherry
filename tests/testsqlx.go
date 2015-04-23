package main

import (
	"github.com/jmoiron/sqlx"
	"log"
	"os"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	url := fmt.Sprintf("postgres://%s@localhost:5432/cherry?sslmode=disable", os.Getenv("USER"));
	log.Println(url)
	db, err := sqlx.Open("postgres", url)
	if err != nil {
		panic("could not connect to db: " + err.Error())
	}

	var id int
	db.QueryRow(db.Rebind("SELECT id FROM users WHERE email = ?"), "test@example.com").Scan(&id)

	log.Println(id)

	log.Println("finished.")
}
