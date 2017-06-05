package msql
/**
	id string

	CpuTotal 	string	//总量
	CpuUsed 	string	//可用
	CpuLeft 	string	//剩余
	CpuRate 	string	//cpu比率

	MemoryTotal 	string	//内存总量
	MemoryUsed  	string	//使用量
	MemoryFree  	string	//剩余量
	MemoryRate	string	//内存比率
	Cached		string

	DiskTotal	string	//硬盘总量
	DiskUsed 	string	//硬盘使用
	DiskAvail	string	//硬盘剩余
	DiskRate	string	//硬盘使用率

	LoadAvg1	string	//
	LoadAvg5	string	//
	LoadAvg15	string	//

	orgMsg		string	//

	CreateTime	string	//
 */

var (
	Sys_Data_sql_list = `select * from sys_data limit ?,?`
	Sys_Data_sql_count = `select count(*) from sys_data`
	Sys_Data_sql_insert = `insert into sys_data(memId, nodeId, cpu_rate, memory_total, memory_used, memory_free, memory_rate, cached,disk_total,disk_used, disk_avail, disk_rate, load_avg1,load_avg5,load_avg15, org_msg, createTime) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,now())`
	Sys_Data_sql_MaxId = `select * from sys_data where id = (select max(id) from sys_data)`
)

