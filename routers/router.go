package routers

import (
	"github.com/astaxie/beego"
	"github.com/yydzero/cherry/controllers/admin"
)

func init() {
	beego.Router("/signup", &admin.AuthController{}, "get,post:Signup")
	beego.Router("/signin", &admin.AuthController{}, "get,post:Signin")
}
