package admin

import (
	"github.com/yydzero/cherry/controllers"
	"log"
)

type AuthController struct {
	controllers.CherryController
}

// Signup register new user
func (c *AuthController) Signup() {
	email := c.GetString("email")
	password := c.GetString("password")

	if email == "" || password == "" {
		c.Fail("email and password could not be empty")
		return
	}

	// store user info into datastore
}

func (c *AuthController) Signin() {
	log.Println("now called.")
}
