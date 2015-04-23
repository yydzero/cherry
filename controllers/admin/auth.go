package admin

import (
	"github.com/yydzero/cherry/controllers"
	"github.com/yydzero/cherry/models/user"
	"fmt"
)

type AuthController struct {
	controllers.CherryController
}

// Signup will register new user
func (c *AuthController) Signup() {
	email := c.GetString("email")
	password := c.GetString("password")

	if email == "" || password == "" {
		c.Fail("email and password could not be empty")
		return
	}

	// store user info into datastore
	u := user.NewUser(email, password)

	hasUser, err := user.HasUser(u)
	if err != nil {
		c.Fail(err.Error())
		return
	}

	if hasUser {
		c.Fail(fmt.Sprintf("User '%s' exist already", email))
		return
	}

	if err := u.Save(); err != nil {
		c.Fail(err.Error())
		return
	}

	c.Ok("user saved")
}

// Login will return token if signin successfully, otherwise report error.
func (c *AuthController) Login() {
	email := c.GetString("email")
	password := c.GetString("password")

	if email == "" || password == "" {
		c.Fail("email and password could not be empty")
		return
	}

	// store user info into datastore
	u := user.NewUser(email, password)
	u, err := user.FindByEmailPassword(u)
	if err != nil || u == nil {
		c.Fail("email and password are incorrect")
		return
	}

	c.SetSession("user", u)


	c.Resource(map[string]string{"sessionid" : c.CruSession.SessionID()})
}

func (c *AuthController) Logout() {

}
