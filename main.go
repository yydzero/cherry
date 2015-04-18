package main

import (
	_ "Users/yyao/workspace/go/goproject/cherry/routers"
	"github.com/astaxie/beego"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	beego.InsertFilter("*",
		beego.BeforeRouter, cors.Allow(&cors.Options{
			AllowOrigins:     []string{"*"},
			AllowMethods:     []string{"POST", "GET"},
			AllowHeaders:     []string{"Origin"},
			ExposeHeaders:    []string{"Content-Length"},
			AllowCredentials: true,
		}))

	beego.Run()
}
