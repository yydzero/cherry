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
	var gzhs []*models.Gzh
	_, err := o.QueryTable("cherry_gzh").All(&gzhs)

	if (err != nil) {
		this.Fail(err.Error())
		return
	}

	this.Resource(gzhs)
}

func (this *GzhController) Delete() {
	openId := this.Ctx.Input.Param(":id")

	if num, err := o.Delete(&models.Gzh{OpenId: openId}); err != nil {
		this.Fail(err.Error())
		return
	} else {
		this.Ok(fmt.Sprintf("%d 删除", num))
	}
}


func (this *GzhController) Post() {
	var gzh models.Gzh
	err := json.Unmarshal(this.Ctx.Input.RequestBody, &gzh)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	if gzh.Name == "" {
		this.Fail("公众号名字不能为空")
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
		this.Fail(fmt.Sprintf("公众号 '%s' 已关注", gzh.Name))
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
