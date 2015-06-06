package models

type Articles struct {
	DocId string		`xml:"docid" orm:"unique"`
	TplId int			`xml:"tplid"`
	Title string		`xml:"title"`
	Url string			`xml:"url"`
	Title1 string		`xml:"title1"`
	ImgLink string		`xml:"imglink"`
	HeadImage string	`xml:"headimage"`
	SourceName string	`xml:"sourcename"`
	Content168 string	`xml:"content168"`
	IsV string			`xml:"isV"`
	OpenId string		`xml:"openid" orm:"index"`
	Content string		`xml:"content"`
	ShowUrl string		`xml:"showurl"`
	Date string			`xml:"date"`
	PageSize string		`xml:"pagesize"`
	LastModified uint64	`xml:"lastModified"`

	Id int				`orm:"pk;auto"`		// need to use 'auto' tag, otherwise, will get 'no LastInsertId available' error.
	FullContent string	`orm:"size(10485760)"`
	CrawledAt uint64
}
