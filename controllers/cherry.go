package controllers

// cherry controllers define base func for all controllers.

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
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

	c.ServeJson()
}

func PrintVarInJson(v interface{}) {
	b, _ := json.Marshal(v)
	var out bytes.Buffer
	json.Indent(&out, b, "", "\t")

	fmt.Printf("%s\n", out.String())
}
