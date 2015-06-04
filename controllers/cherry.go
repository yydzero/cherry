package controllers

// cherry controllers define base func for all controllers.

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/ninedata/goat"
	"strconv"
)

var o orm.Ormer

func init() {
	o = orm.NewOrm()
	o.Using("default")
}

type CherryController struct {
	beego.Controller
}

type Response struct {
	Status    string
	Message   string
	Resource  interface{}
	Resources []interface{}
}

func (c *CherryController) GetId() (int, error) {
	idstr := c.Ctx.Input.Param(":id")

	return strconv.Atoi(idstr)
}

func (c *CherryController) GetORM() orm.Ormer {
	return o
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
	r := Response{
		Status: "ok",
		Resource: v,
	}
	c.Data["json"] = r
	c.ServeJson()
}

func (c *CherryController) Resources(v interface{}) {
	r := Response{
		Status: "ok",
		Resources: v.([]interface{}),
	}
	c.Data["json"] = r
	c.ServeJson()
}

