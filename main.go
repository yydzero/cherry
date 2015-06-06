package main

import (
	"log"

	"github.com/astaxie/beego"
	// To init models
	_ "github.com/yydzero/cherry/models"

	// To init routers
	_ "github.com/yydzero/cherry/routers"

	_ "github.com/lib/pq"
	"github.com/astaxie/beego/plugins/cors"
)

func main() {
	log.SetFlags(log.Ltime | log.Lshortfile)

	// Enable CORS
	opts := &cors.Options{
		AllowOrigins:     []string{"*"},
//		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "DELETE", "PUT", "OPTIONS", "POST"},
		AllowHeaders:     []string{"Origin", "X-Requested-With", "Content-Type", "Accept", "X-XSRF-TOKEN", "X-Xsrftoken"},
		ExposeHeaders:    []string{"Content-Length", "Content-Type"},
		AllowCredentials: true,
	}
	filterFunc := cors.Allow(opts)

	beego.DirectoryIndex=true
	beego.StaticDir["/static"] = "/Users/yyao/workspace/go/golang/src/github.com/yydzero/cherry/client/dist"

	beego.InsertFilter("*", beego.BeforeRouter, filterFunc)

	beego.Run()
}
