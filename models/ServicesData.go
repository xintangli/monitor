package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type ServicesData struct {
	Id 		int	`orm:"pk"`
	HostsId 	int	`orm:"null" form:"hostsId"`
	MemId 		string	`orm:"null" form:"memId"`
	NodeId		string	`orm:"null" form:"nodeId"`
	SvcType 	int	`orm:"null" form:"svcType"`
	QpsSucc 	int	`orm:"null" form:"qpsSucc"`
	QpsFail 	int	`orm:"null" form:"qpsFail"`
	ProcessNum 	int	`orm:"null" form:"processNum"`
	RedisUsability 	string	`orm:"null" form:"redisUsability"`
	RedisMem 	string	`orm:"null" form:"redisMem"`
	RedisSize 	int	`orm:"null" form:"redisSize"`

	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
}

func (h *ServicesData) Insert() (int64, error)  {
	o := orm.NewOrm()
	return o.Insert(h)
}

func (h *ServicesData) Read() error {
	o := orm.NewOrm()
	return o.Read(h)
}

func (h *ServicesData) Page(limit int, offset int) ([]ServicesData, int64, error) {
	o := orm.NewOrm()
	var list []ServicesData
	_, err := o.QueryTable(h).OrderBy("-create_time").Limit(limit, offset).All(&list)
	o = orm.NewOrm()
	count, err := o.QueryTable(h).Count()
	return list, count, err

}

func (h *ServicesData) Update() (int64, error) {
	o := orm.NewOrm()
	return o.Update(h)
}

func (h *ServicesData) GetNewestData(limit int) ([]ServicesData, error) {
	o := orm.NewOrm()
	var list []ServicesData
	if limit <= 0 {
		limit = 10
	}
	_, err := o.QueryTable(h).OrderBy("-create_time").Limit(limit).All(&list)
	return list, err
}

func (h *ServicesData) List(limit int, offset int) ([]ServicesData, error)  {
	o := orm.NewOrm()
	var list []ServicesData
	qt := o.QueryTable(h)
	if h.HostsId > 0 {
		qt = qt.Filter("HostsId", h.HostsId)
	}

	_, err := qt.OrderBy("create_time").Limit(limit, offset).All(&list)
	return list, err

}

type PageServicesData struct {
	SEcho int			`json:"sEcho"`
	ITotalRecords int64		`json:"iTotalRecords"`
	ITotalDisplayRecords int64	`json:"iTotalDisplayRecords"`
	AAData	[]ServicesData		`json:"aaData"`
}
