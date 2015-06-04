package controllers

import (
	"github.com/astaxie/beego"
	"github.com/yydzero/surf/weichat"
	"github.com/yydzero/surf/weichat/api"
	"sync"
	"time"
	"github.com/yydzero/cherry/models"
	"github.com/astaxie/beego/orm"
	"log"
)

type CrawlController struct {
	CherryController
}

// Signup will register new user
// TODO: pg 不支持 byte[] 类型。
// TODO: 更健壮的 chan 模式来处理任何一端可能出错。
// "no LastInsertId available"???
func (this *CrawlController) Get() {
	openId := this.Ctx.Input.Param(":id")

	var stat Stats
	err := o.Raw(STATS_QUERY_ROW, openId).QueryRow(&stat)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	beego.Notice("crawling articles for Id: ", openId)

	now := time.Now()
	secs := now.Unix()

	c := make(chan *weichat.Article, 10)
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()

		for article := range c {
			old := models.Articles{DocId: article.DocId}

			err := o.Read(&old, "DocId")
			if err == nil || err != orm.ErrNoRows {
				log.Printf("article %s ('%s') has been crawled already. err = %v\n", article.Title, article.DocId, err)
				continue
			}

			var a models.Articles

			a.DocId = article.DocId
			a.Content = article.Content
			a.Date = article.Date
			a.HeadImage = article.HeadImage
			a.LastModified  = article.LastModified
			a.OpenId = article.OpenId
			a.SourceName = article.SourceName
			a.Title = article.Title
			a.Url = article.Url
			a.ImgLink = article.ImgLink
			a.ShowUrl = article.ShowUrl
			a.PageSize = article.PageSize
			a.CrawledAt = uint64(secs)

//			log.Printf("article: %v\n", a)

			a.FullContent = article.FullContent


			if _, err := this.GetORM().Insert(&a); err != nil {
				log.Printf("failed to insert %v into database\n", a.DocId)
				this.Fail(err.Error())
//				close(c)	// 不能在这里close，否则造成其他goroutine向chan放置数据时异常。
				return
			}
		}
	}()

	err = api.CrawlGzh(openId, "", stat.LastModified, c)
	if err != nil {
		this.Fail(err.Error())
		//				close(c)	// 不能在这里close，否则造成其他goroutine向chan放置数据时异常。
		return
	}
	close(c)

	wg.Wait()

	// 返回该公众号的爬去状态，例如文章数、最新文章日期等。
	err = o.Raw(STATS_QUERY_ROW, openId).QueryRow(&stat)
	if err != nil {
		this.Fail(err.Error())
		return
	}

	beego.Notice(stat)

	this.Resource(stat)
}
