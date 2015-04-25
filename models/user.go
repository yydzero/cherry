package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type User struct {
	Id        int    `form:"id"`
	First     string `form:"first" valid:"Required"`
	Last      string `form:"last"`
	Email     string `form:"email" valid:"Email" orm:"unique"`
	Password  string `form:"password" valid:"MinSize(6);MaxSize(30)"`
	Reg_key   string
	Reg_date  time.Time `orm:"auto_now_add;type(datetime)"`
	Reset_key string
}

func HasUserByEmail(u *User)  (*User, error) {
	err := o.Read(u, "Email")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return nil, nil
	}
	return u, err
}