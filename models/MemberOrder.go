package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	derrors "github.com/xintangli/monitor/error"
)

type MemberOrder struct {

	Id int		`orm:"pk"`

	MemId		string	`orm:"null" from:"hostId" json:"mem_id"`
	MemName 	string	`orm:"null" form:"memName" json:"mem_name"`
	SvcType 	int	`orm:"null" form:"svcType" json:"svc_type"`

	OrderId 	string	`orm:"null" form:"OrderId" json:"order_id"`//总量
	TaskId 		string	`orm:"null" form:"cpuUsed" json:"task_id"`//可用

	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)" json:"create_time"`
}


func (m *MemberOrder) Insert() (int64, error)  {
	o := orm.NewOrm()
	return o.Insert(m)
}

func (m *MemberOrder) Read() error {
	o := orm.NewOrm()
	return o.Read(m)
}

func (m *MemberOrder) Page(limit int, offset int) ([]MemberOrder, int64, error) {
	o := orm.NewOrm()
	var list []MemberOrder
	_, err := o.QueryTable(m).OrderBy("-create_time").Limit(limit, offset).All(&list)
	o = orm.NewOrm()
	count, err := o.QueryTable(m).Count()
	return list, count, err

}

func (m *MemberOrder) Update() (int64, error) {
	o := orm.NewOrm()
	return o.Update(m)
}

func (m *MemberOrder) GetNewestData(limit int) ([]MemberOrder, error) {
	o := orm.NewOrm()
	var list []MemberOrder
	if limit <= 0 {
		limit = 10
	}
	_, err := o.QueryTable(m).OrderBy("-create_time").Limit(limit).All(&list)
	return list, err
}

func (m *MemberOrder) GetListByMemId(memId string) (*MemberOrder, error) {
	o := orm.NewOrm()
	qt := o.QueryTable(m)
	if m.MemId == "" {
		return m, derrors.New("1099", "memId is null")
	}
	qt = qt.Filter("MemId", m.MemId)
	err := qt.OrderBy("-create_time").One(m)
	return m, err
}

func (m *MemberOrder) GetListByMemIdAndSvcType() ([]MemberOrder, error) {
	o := orm.NewOrm()
	qt := o.QueryTable(m)
	var list []MemberOrder
	qt = qt.Filter("MemId", m.MemId)
	qt = qt.Filter("SvcType", m.SvcType)
	_, err := qt.OrderBy("-create_time").All(&list)
	return list, err
}

type PageMemberOrder struct {
	SEcho int			`json:"sEcho"`
	ITotalRecords int64		`json:"iTotalRecords"`
	ITotalDisplayRecords int64	`json:"iTotalDisplayRecords"`
	AAData	[]MemberOrder		`json:"aaData"`
}


