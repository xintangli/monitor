package msql
/**
	Id string
	HostsId string
	MemId 		string
	NodeId		string
	SvcType string
	QpsSucc string
	QpsFail string
	ProcessNum string
	RedisUsability string
	RedisMem string
	RedisSize string

	CreateTime string
 */

var (
	Services_Data_sql_list = `select * from services_data order by createtime desc limit ?,?`
	Services_Data_sql_count = `select count(*) from services_data`
	Services_Data_sql_insert = `insert into services_data(hostsId, memId, nodeId, svc_type, qps_succ, qps_fail, process_num, redis_usability,redis_mem, redis_size, createTime) values(?,?,?,?,?,?,?,?,?,?,now())`
	Services_Data_sql_MaxId = `select * from services_data where id = (select max(id) from services_data where svc_type = ? )`
)

