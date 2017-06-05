package models

import "time"

type MonitorData struct {
	Id 		string
	HostsId 	string	`form:"hostsid" binding:"required"`
	Data 		string	`form:"data" binding:"required"`
	CreateTime 	time.Time	`orm:"auto_now_add;type(datetime)"`
}
