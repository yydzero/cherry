package main

import (
	"log"
	"github.com/yydzero/cherry/models"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

var schema string = `
	CREATE TABLE IF NOT EXISTS users (
	  id serial,
	  name text,
	  email text not null,
	  password text not null
	);

	CREATE TABLE IF NOT EXISTS wechat (
	  id serial,
	  name text not null,
	  wechat_id text not null
	);
`

func main() {
	db := models.GetDB()

	db.MustExec(schema)

	log.Println("load schema finished.")
}