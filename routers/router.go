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
}
