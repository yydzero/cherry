package controllers
import (
	"github.com/yydzero/cherry/models"
)

type ArticlesController struct {
	CherryController
}

// Signup will register new user
// TODO: pg 不支持 byte[] 类型。
func (this *ArticlesController) Get() {
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