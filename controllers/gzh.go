package controllers

import (
	"github.com/yydzero/cherry/models"

	"encoding/json"
	"fmt"
	"log"
	weichat "github.com/yydzero/surf/weichat/api"
)

type GzhController struct {
	CherryController
}

// Signup will register new user
func (this *GzhController) Get() {
	var groups []*models.Group
	_, err := o.QueryTable("cherry_group").All(&groups)

	if (err != nil) {
		this.Fail(err.Error())
		return
	}

	this.Resource(groups)
}


func (this *GzhController) Post() {
	var gzh models.Gzh
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &gzh)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	if gzh.Name == "" {
		this.Fail("gzh name could not be empty")
		return
	}

	// 是否已经关注过次公众号
	g, err := models.HasGzhByName(& gzh)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	// 是否已经存在
	if g != nil {
		this.Fail(fmt.Sprintf("gzh '%s' exist already", gzh.Name))
		return
	}

	log.Printf("判断公众号 '%s' 是否存在.", gzh.Name)

	o, err := weichat.GetGzhByName(gzh.Name)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	gzh.OpenId = o.OpenID
	gzh.WXNo = o.WXNo
	gzh.WXCert = o.WXCert
	gzh.Introduction = o.Introduction

	if gzh.OpenId == "" {
		this.Fail(fmt.Sprintf("找不到公众号 '%s' 的OpenId", gzh.Name))
		return
	}


	// 插入这个公众号
	if _, err := this.GetORM().Insert(&gzh); err != nil && err.Error() != "no LastInsertId available" {
		this.Fail(err.Error())
		return
	}

	this.Resource(gzh)
}
