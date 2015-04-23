package controllers

// cherry controllers define base func for all controllers.

import (
	"github.com/astaxie/beego"
	"github.com/ninedata/goat"
)

type CherryController struct {
	beego.Controller
}

type Response struct {
	Status    string
	Message   string
	Resources []interface{}
}

func (c *CherryController) Fail(message string) {
	r := Response{
		Status:  "fail",
		Message: message,
	}

	c.Data["json"] = r
	goat.PrintVarInJson(r)

	c.ServeJson()
}

func (c *CherryController) Ok(message string) {
	r := Response{
		Status:  "ok",
		Message: message,
	}

	c.Data["json"] = r
	c.ServeJson()
}

func (c *CherryController) Resource(v interface{}) {
	resources := make([]interface{}, 0)
	resources = append(resources, v)
	r := Response{
		Status: "ok",
		Resources: resources,
	}
	c.Data["json"] = r
	c.ServeJson()
}

