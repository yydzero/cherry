package routers

import (
	"github.com/astaxie/beego"
	"github.com/yydzero/cherry/controllers"
	"github.com/yydzero/cherry/controllers/admin"
)

func init() {
	beego.Router("/signup", &admin.AuthController{}, "get,post:Signup")
	beego.Router("/signin", &admin.AuthController{}, "get,post:Signin")
	beego.Router("/dd", &admin.DDController{}, "get,post:Signin")

	beego.Router("/", &controllers.MainController{})
}
