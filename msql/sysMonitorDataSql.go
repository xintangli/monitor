package msql
/**
type MonitorData struct {
	id int
	hostsId int
	data string
	createTime string
}
 */
var (
	Sys_MonitorData_sql_list = `select * from monitor_data limit ?,?`
	Sys_MonitorData_sql_count = `select count(*) from monitor_data`
	Sys_MonitorData_sql_insert = `insert into monitor_data(hostsId, data, createTime) values(?,?,now())`
	Sys_MonitorData_sql_MaxId = `select * from monitor_data where id = (select max(id) from monitor_data)`
)
