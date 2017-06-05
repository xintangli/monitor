package models

import (
	"time"
	"github.com/astaxie/beego/orm"
	"fmt"
)

type BusiData struct {
	Id 		int		`orm:"pk" json:"id"`

	MemId 		string		`orm:"null" json:"mem_id"`
	NodeId		string		`orm:"null" json:"node_id"`
	OrderId		string		`orm:"null" json:"order_id"`

	HostsId 	int		`orm:"null" json:"hosts_id"`
	SvcType 	int		`orm:"null" json:"svc_type"`

	DemSuccNum 	int		`orm:"null" json:"dem_succ_num"`
	DemReqNum 	int		`orm:"null" json:"dem_req_num"`
	DemCacheNum 	int		`orm:"null" json:"dem_cache_num"`

	SupQueryTotal 	int		`orm:"null" json:"sup_query_total"`
	SupQuerySuccess int		`orm:"null" json:"sup_query_success"`

	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)" json:"create_time"`

}

func (b *BusiData) Insert() (int64, error)  {
	o := orm.NewOrm()
	return o.Insert(b)
}

func (b *BusiData) Read() error {
	o := orm.NewOrm()
	return o.Read(b)
}

func (b *BusiData) Page(limit int, offset int) ([]BusiData, int64, error) {
	if limit <= 0 {
		limit = 10
	}
	if offset <= 0 {
		offset = 0
	}
	o := orm.NewOrm()
	var list []BusiData
	_, err := o.QueryTable(b).Filter("HostsId", b.HostsId).OrderBy("-create_time").Limit(limit, offset).All(&list)
	o = orm.NewOrm()
	count, err := o.QueryTable(b).Filter("HostsId", b.HostsId).Count()
	return list, count, err

}

func (b *BusiData) List(limit int, offset int, svcType int, orderId string) ([]BusiData, error)  {
	o := orm.NewOrm()
	var list []BusiData
	qt := o.QueryTable(b)
	/*var orderIds []string
	var taskIds []string*/
	/*for _, mo := range mos {
		orderIds = append(orderIds, mo.OrderId)
		taskIds = append(taskIds, mo.TaskId)
	}*/
	/*if b.SvcType == 1 {
		qt = qt.Filter("OrderId__in", orderIds)
	}else {
		qt = qt.Filter("TaskId__in", taskIds)
	}*/
	if orderId != "" {
		qt = qt.Filter("OrderId", orderId)
	}
	if svcType >= 0 {
		qt = qt.Filter("SvcType", svcType)
	}
	fmt.Println(orderId, svcType)
	_, err := qt.OrderBy("create_time").Limit(limit, offset).All(&list)
	fmt.Println(list)
	return list, err

}

func (b *BusiData) ListOuter(startTime string, endTime string, memId string, orderId string) ([]BusiData, error) {
	o := orm.NewOrm()
	var list []BusiData
	qs := o.QueryTable(b)
	qs.Filter("MemId", memId)
	qs.Filter("OrderId", orderId)
	qs.Filter("CreateTime__gte", startTime)
	qs.Filter("CreateTime__lte", endTime)
	_, err := qs.All(&list)
	return list, err

}

func (b *BusiData) Update() (int64, error) {
	o := orm.NewOrm()
	return o.Update(b)
}

func (b *BusiData) GetNewestData(limit int) ([]BusiData, error) {
	o := orm.NewOrm()
	var list []BusiData
	if limit <= 0 {
		limit = 10
	}
	_, err := o.QueryTable(b).OrderBy("-create_time").Limit(limit).All(&list)
	return list, err
}

type PageBusiData struct {
	Draw int			`json:"draw"`
	RecordsTotal int64		`json:"recordsTotal"`
	RecordsFiltered int64		`json:"recordsFiltered"`
	Data	[]BusiData		`json:"data"`
}
