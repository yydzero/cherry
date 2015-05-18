package models
import "github.com/astaxie/beego/orm"


type Group struct {
	Id 		int
	Name    string
	Categories string
}

func HasGroupByName(this *Group)  (*Group, error) {
	err := o.Read(this, "Name")
	if err == orm.ErrNoRows || err == orm.ErrMissPK {
		return nil, nil
	}
	return this, err
}