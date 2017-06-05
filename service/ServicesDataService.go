package service

import (
	"github.com/xintangli/monitor/models"
	/*"github.com/hzwy23/dbobj"
	"github.com/xintangli/monitor/msql"
	"fmt"*/
)

type ServicesDataService struct {
	models *models.ServicesData
}

var ServicesDataSvc = &ServicesDataService{
	new(models.ServicesData),
}
/*
func (h *ServicesDataService) List(offset int, limit int,servicesData *models.ServicesData) ([]models.ServicesData, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	fmt.Println(offset, limit)
	rows, err := dbobj.Query(msql.Services_Data_sql_list, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, err
	}
	defer rows.Close()
	var dataList []models.ServicesData
	fmt.Println(rows.Columns())
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, err
	}
	return dataList, nil
}


func (h *ServicesDataService) Page(offset int, limit int,servicesData *models.ServicesData) ([]models.ServicesData, int64, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	fmt.Println(offset, limit)
	rows, err := dbobj.Query(msql.Services_Data_sql_list, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	defer rows.Close()
	var dataList []models.ServicesData
	fmt.Println(rows.Columns())
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(msql.Services_Data_sql_count).Scan(&total)

	return dataList, total, nil

}

func (h *ServicesDataService) GetNewestData() (models.ServicesData, error) {
	var data models.ServicesData
	rows, err :=dbobj.Query(msql.Services_Data_sql_MaxId, "1")
	if err != nil {
		fmt.Println("query data error", dbobj.GetErrorMsg(err))
		return data, err
	}
	defer rows.Close()
	var dataList []models.ServicesData
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return data, err
	}
	if len(dataList) > 0 {
		return dataList[0], nil
	}
	return dataList[0], nil
}

func (h *ServicesDataService) Put(servicesData models.ServicesData) error {
	tx, err := dbobj.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	*//**
	hostsId, svc_type, qps_succ, qps_fail, process_num, redis_usability,redis_mem, redis_size, createTime
	 *//*
	_, err = tx.Exec(msql.Services_Data_sql_insert, servicesData.HostsId, servicesData.MemId, servicesData.NodeId, servicesData.SvcType, servicesData.QpsSucc,
			servicesData.QpsFail, servicesData.ProcessNum, servicesData.RedisUsability, servicesData.RedisMem, servicesData.RedisSize)
	if err != nil {
		fmt.Println("insert fail err :", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}*/
