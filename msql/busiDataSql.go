package msql
/**
	Id string
	HostsId string
	SvcType string

	DemSuccNum string
	DemReqNum string
	DemCacheNum string

	SupQueryTotal string
	SupQuerySuccess string

	createTime string
 */

var (
	Busi_Data_sql_list = `select * from busi_data where svc_type = ? order by createtime desc limit ?,?`
	Busi_Data_sql_count = `select count(*) from busi_data`
	Busi_Data_sql_insert = `insert into busi_data(hostsId, memId, nodeId, svc_type, dem_succ_num, dem_req_num, dem_cache_num, sup_query_total,sup_query_success, createTime) values(?,?,?,?,?,?,?,?,?,now())`
	Busi_Data_sql_MaxId = `select * from busi_data where id = (select max(id) from busi_data where svc_type = ? )`
)

