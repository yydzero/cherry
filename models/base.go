package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	"os"
)

var o orm.Ormer

func init() {
	orm.Debug = true
	orm.RegisterDriver("postgres", orm.DR_Postgres)

	url := fmt.Sprintf("postgres://%s@localhost:5432/cherry?sslmode=disable", os.Getenv("USER"));

	maxIdle := 30
	maxConn := 30
	orm.RegisterDataBase("default", "postgres", url, maxIdle, maxConn)

	orm.RegisterModelWithPrefix("cherry_", new(User))
	orm.RegisterModelWithPrefix("cherry_", new(Group))
	orm.RegisterModelWithPrefix("cherry_", new(Gzh))

	err := orm.RunSyncdb("default", false, false)
	if err != nil {
		fmt.Println(err)
	}

	o = orm.NewOrm()
	o.Using("default")
}

//
//import (
//	"github.com/jmoiron/sqlx"
//	"fmt"
//	"os"
//	"log"
//
//	_ "github.com/lib/pq"
//)
//
//var db *sqlx.DB = nil
//var err error
//
//// MySQL uses the ? variant shown above
//// PostgreSQL uses an enumerated $1, $2, etc bindvar syntax
//// SQLite accepts both ? and $1 syntax
//// Oracle uses a :name syntax
//func GetDB() *sqlx.DB {
//	if db == nil {
//		url := fmt.Sprintf("postgres://%s@localhost:5432/cherry?sslmode=disable", os.Getenv("USER"));
//		log.Println(url)
//		if db, err = sqlx.Open("postgres", url); err != nil {
//			panic("could not connect to db: " + err.Error())
//		}
//
//
//	}
//	return db
//}