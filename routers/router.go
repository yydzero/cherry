package routers

import (
	"github.com/astaxie/beego"
	"github.com/yydzero/cherry/controllers/admin"
	"github.com/yydzero/cherry/controllers"
)

func init() {
	beego.Router("/signup", &admin.AuthController{}, "get,post:Signup")
	beego.Router("/login", &admin.AuthController{}, "get,post:Login")
	beego.Router("/logout", &admin.AuthController{}, "get,post:Logout")

	beego.Router("/groups/?:id", &controllers.GroupController{})

	// 公众号 CRUD 基本信息
	beego.Router("/gzh/?:id", &controllers.GzhController{})

	// 公众号爬虫爬取情况
	beego.Router("/stats/?:id", &controllers.StatsController{})

	// 爬虫
	beego.Router("/crawl/?:id", &controllers.CrawlController{})

	// 文章列表
	beego.Router("/article/?:id", &controllers.ArticleController{})
}
