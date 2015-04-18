package routers

import (
	"Users/yyao/workspace/go/goproject/cherry/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
