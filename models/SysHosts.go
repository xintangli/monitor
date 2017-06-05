package models

import (
	"time"
	"github.com/astaxie/beego/orm"
)

type SysHosts struct {
	Id 		int	`orm:"pk" json:"id"`
	Name 		string	`orm:"null" form:"name" json:"name"`
	MemId		string	`orm:"null" form:"memId" json:"mem_id"`
	NodeId 		string	`orm:"null" form:"nodeId" json:"node_id"`
	SvcType 	int	`orm:"null" form:"svcType" json:"svc_type"`
	InIp 		string	`orm:"null" form:"inIp" json:"in_ip"`
	OutIp 		string	`orm:"null" form:"outIp" json:"out_ip"`
	Cpu 		string	`orm:"null" form:"cpu" json:"cpu"`
	Memory 		string	`orm:"null" form:"memory" json:"memory"`
	Disk 		string	`orm:"null" form:"disk" json:"disk"`
	Status 		string	`orm:"null" form:"status" json:"status"`
	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)" json:"create_time"`
}

func (h *SysHosts) Insert() (int64, error)  {
	o := orm.NewOrm()
	return o.Insert(h)
}

func (h *SysHosts) Read() error {
	o := orm.NewOrm()
	return o.Read(h)
}

func (h *SysHosts) ReadByNodeIdAndMemId() error {
	o := orm.NewOrm()
	qt := o.QueryTable(h)
	qt = qt.Filter("MemId", h.MemId)
	qt = qt.Filter("NodeId", h.NodeId)
	return qt.One(h)
}

func (h *SysHosts) Page(limit int, offset int) ([]SysHosts, int64, error) {
	o := orm.NewOrm()
	var hostsList []SysHosts
	_, err := o.QueryTable(h).OrderBy("create_time").Limit(limit, offset).All(&hostsList)
	o = orm.NewOrm()
	count, err := o.QueryTable(h).Count()
	return hostsList, count, err

}

func (h *SysHosts) Update() (int64, error) {
	o := orm.NewOrm()
	return o.Update(h)
}

type PageSysHosts struct {
	Draw int			`json:"draw"`
	RecordsTotal int64		`json:"recordsTotal"`
	RecordsFiltered int64		`json:"recordsFiltered"`
	Data	[]SysHosts		`json:"data"`
}
