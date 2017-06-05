package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type SysData struct {

	Id int		`orm:"pk"`

	HostsId		int	`orm:"null" from:"hostsId" json:"hosts_id"`
	MemId 		string	`orm:"null" form:"memId" json:"mem_id"`
	NodeId		string	`orm:"null" form:"nodeId" json:"node_id"`

	CpuTotal 	string	`orm:"null" form:"cpuTotal" json:"cpu_total"`//总量
	CpuUsed 	string	`orm:"null" form:"cpuUsed" json:"cpu_used"`//可用
	CpuLeft 	string	`orm:"null" form:"cpuLeft" json:"cpu_left"`//剩余
	CpuRate 	string	`orm:"null" form:"cpuRate" json:"cpu_rate"`//cpu比率

	MemoryTotal 	string	`orm:"null" form:"memoryTotal" json:"memory_total"`//内存总量
	MemoryUsed  	string	`orm:"null" form:"memoryUsed" json:"memory_used"`//使用量
	MemoryFree  	string	`orm:"null" form:"memoryFree" json:"memory_free"`//剩余量
	MemoryRate	string	`orm:"null" form:"memoryRate" json:"memory_rate"`//内存比率
	Cached		string	`orm:"null" form:"cached" json:"cached"`

	DiskTotal	string	`orm:"null" form:"diskTotal" json:"disk_total"`//硬盘总量
	DiskUsed 	string	`orm:"null" form:"diskUsed" json:"disk_used"`//硬盘使用
	DiskAvail	string	`orm:"null" form:"diskAvail" json:"disk_avail"`//硬盘剩余
	DiskRate	string	`orm:"null" form:"diskRate" json:"disk_rate"`//硬盘使用率

	LoadAvg1	string	`orm:"null" form:"loadAvg1" json:"load_avg_1"`//
	LoadAvg5	string	`orm:"null" form:"loadAvg5" json:"load_avg_5"`//
	LoadAvg15	string	`orm:"null" form:"loadAvg15" json:"load_avg_15"`//

	OrgMsg		string	`orm:"null" form:"orgMsg" json:"org_msg"`//

	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)" json:"create_time"`
}


func (h *SysData) Insert() (int64, error)  {
	o := orm.NewOrm()
	return o.Insert(h)
}

func (h *SysData) Read() error {
	o := orm.NewOrm()
	return o.Read(h)
}

func (h *SysData) Page(limit int, offset int) ([]SysData, int64, error) {
	fmt.Printf("sysDatas page, limit : %d, offset : %d, req : %d ", limit, offset, h.HostsId)
	o := orm.NewOrm()
	var list []SysData
	qt := o.QueryTable(h)
	if h.HostsId > 0 {
		qt = qt.Filter("HostsId", h.HostsId)
	}
	qt.OrderBy("-create_time").Limit(limit, offset).All(&list)

	o = orm.NewOrm()
	cqt := o.QueryTable(h)
	if h.HostsId > 0 {
		cqt = cqt.Filter("HostsId", h.HostsId)
	}
	count, err := cqt.Count()
	return list, count, err

}

func (h *SysData) Update() (int64, error) {
	o := orm.NewOrm()
	return o.Update(h)
}

func (h *SysData) GetNewestData(limit int) ([]SysData, error) {
	o := orm.NewOrm()
	var list []SysData
	if limit <= 0 {
		limit = 10
	}
	_, err := o.QueryTable(h).OrderBy("-create_time").Limit(limit).All(&list)
	return list, err
}

func (h *SysData) GetNewestOne() (*SysData, error) {
	o := orm.NewOrm()
	qt := o.QueryTable(h)
	fmt.Println("h.hostsId =", h.HostsId)
	if h.HostsId > 0 {
		qt = qt.Filter("HostsId", h.HostsId)
	}
	err := qt.OrderBy("-create_time").One(h)
	return h, err
}

type PageSysData struct {
	Draw int			`json:"draw"`
	RecordsTotal int64		`json:"recordsTotal"`
	RecordsFiltered int64		`json:"recordsFiltered"`
	Data	[]SysData		`json:"data"`
}


