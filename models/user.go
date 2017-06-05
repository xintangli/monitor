package models

import "github.com/astaxie/beego/orm"

const (
	USER = iota
	ADMIN
)

type User struct {
	ID       int    `orm:"pk" json:"id"`
	Username string `orm:"null" json:"user_name"`
	Password string `orm:"null" json:"password"`
	Role     int    `orm:"null" json:"role"`
	Phone    string `orm:"null" json:"phone"`
	Mail     string `orm:"null" json:"mail"`
	Weixin   string `orm:"null" json:"weixin"`
	Status   int    `orm:"null" json:"status"`
}

func (u *User) Read() error {
	o := orm.NewOrm()
	return o.Read(u);
}

func (u *User) ReadByUserName() error {
	o := orm.NewOrm()
	return o.QueryTable(u).Filter("Username", u.Username).One(o)
}

