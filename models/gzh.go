package models

import (
	_ "github.com/yydzero/surf/weichat"
	"github.com/astaxie/beego/orm"
)



type Gzh struct {
//	weichat.PublicNo
	Name string
	WXNo string	`orm:"column(wx_no)"`
	Introduction string
	WXCert string `orm:"column(wx_cert)"`
	OpenId string `orm:"pk;column(open_id)"`
}

func HasGzhByName(this *Gzh)  (*Gzh, error) {
	err := o.Read(this, "Name")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return nil, nil
	}
	return this, err
}