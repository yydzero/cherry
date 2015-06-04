package controllers
import (
	"github.com/yydzero/cherry/models"
	"strconv"
)

type ArticleController struct {
	CherryController
}

// Signup will register new user
// TODO: pg 不支持 byte[] 类型。
func (this *ArticleController) Get() {
	openId := this.GetString("openId")
	docId := this.GetString("docId")

	var articles []*models.Articles
	qs := o.QueryTable("cherry_articles")
	if openId != "" {
		qs = qs.Filter("OpenId", openId)
	}
	if docId != "" {
		qs = qs.Filter("DocId", docId)
	}

	_, err := qs.All(&articles, "Id", "OpenId", "SourceName", "Title", "Url", "Content", "PageSize", "LastModified", "Date")
	if err != nil {
		this.Fail(err.Error())
		return
	}

	this.Resource(articles)
}

func (this *ArticleController) Delete() {
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

	if _, err := o.Delete(&models.Articles{Id: id}); err != nil {
		this.Fail(err.Error())
		return
	}

	this.Ok("group deleted")
}
