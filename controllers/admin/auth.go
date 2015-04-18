package admin

import (
	"github.com/astaxie/beego"
	"log"
)

type AuthController struct {
	beego.Controller
}

func (c *AuthController) Signup() {
	log.Println("now called.")
}

func (c *AuthController) Signin() {
	log.Println("now called.")
}
