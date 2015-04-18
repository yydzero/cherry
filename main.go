package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"log"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	// Enable CORS
	opts := &cors.Options{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}
	filterFunc := cors.Allow(opts)

	beego.InsertFilter("*", beego.BeforeRouter, filterFunc)

	beego.Run()
}
