package controllers

import (
	"github.com/yydzero/cherry/models"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/ninedata/goat"
)

type GroupController struct {
	CherryController
}

// Signup will register new user
func (this *GroupController) Get() {
	var groups []*models.Group
	_, err := o.QueryTable("cherry_group").All(&groups)

	if (err != nil) {
		this.Fail(err.Error())
		return
	}

	this.Resource(groups)
}

func (this *GroupController) Post() {
	var group models.Group
	json.Unmarshal(this.Ctx.Input.RequestBody, &group)

	if group.Name == "" {
		this.Fail("group name could not be empty")
		return
	}

	u1, err := models.HasGroupByName(& group)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	if u1 != nil {
		this.Fail(fmt.Sprintf("Group '%s' exist already", group.Name))
		return
	}

	if _, err := this.GetORM().Insert(&group); err != nil {
		this.Fail(err.Error())
		return
	}

	this.Ok("group saved")
}

func (this *GroupController) Delete() {
	idstr := this.Ctx.Input.Param(":id")

	if idstr == "" {
		this.Fail("delete resource without id")
		return
	}

	id, err := strconv.Atoi(idstr)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	goat.PrintVarInJson(o)

	if _, err := o.Delete(&models.Group{Id: id}); err != nil {
		this.Fail(err.Error())
		return
	}

	this.Ok("group deleted")
}