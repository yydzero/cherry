package controllers
import (
	"github.com/yydzero/cherry/models"
)

type ArticleController struct {
	CherryController
}

// Signup will register new user
// TODO: pg 不支持 byte[] 类型。
func (this *ArticleController) Get() {
	id, err := this.GetId()
	if err != nil {
		this.Fail(err.Error())
		return
	}

	article := models.Articles{Id: id}

	if err = o.Read(&article); err != nil {
		this.Fail(err.Error())
		return
	}

	this.Resource(article)
}

func (this *ArticleController) Delete() {
	id, err := this.GetId()
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
