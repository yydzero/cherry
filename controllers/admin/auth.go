package admin

import (
	"github.com/yydzero/cherry/controllers"
	"github.com/yydzero/cherry/models"
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
	user := models.User{Email: email, Password: password}

	u1, err := models.HasUserByEmail(& models.User{Email: email})
	if err != nil {
		c.Fail(err.Error())
		return
	}

	if u1 != nil {
		c.Fail(fmt.Sprintf("User '%s' exist already", email))
		return
	}

	if _, err := c.GetORM().Insert(&user); err != nil {
		c.Fail(err.Error())
		return
	}

	c.Ok("user saved")
}

// Login will return token if signin successfully, otherwise report error.
func (c *AuthController) Login() {
	s := c.GetSession("user")
	if s != nil {
		c.Redirect("/home", 302)
	}

	email := c.GetString("email")
	password := c.GetString("password")

	if email == "" || password == "" {
		c.Fail("email and password could not be empty")
		return
	}

	u1, err := models.HasUserByEmail(& models.User{Email: email})
	if err != nil || u1 == nil || u1.Password != password {
		c.Fail("email and password are incorrect")
		return
	}

	c.SetSession("user", u1)

	r := map[string]interface{}{
		"sessionid": c.CruSession.SessionID(),
		"user": &models.User{
			Id: u1.Id,
			Email: u1.Email,
			First: u1.First,
			Last: u1.Last,
		},
	}

	c.Resource(r)
}

func (c *AuthController) Logout() {

}
