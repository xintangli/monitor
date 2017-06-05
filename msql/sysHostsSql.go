package msql

var (
	Sys_hosts_sql_list = `select * from sys_hosts limit ?,?`
	Sys_hosts_sql_count = `select count(*) from sys_hosts limit ?,?`
	Sys_hosts_sql_insert = `insert into sys_hosts(id, name, inIp, outIp, cpu, memory, disk, status, createTime) values(?,?,?,?,?,?,?,?,now())`
)
