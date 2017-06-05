package service

import (
	"github.com/xintangli/monitor/models"
	/*"github.com/hzwy23/dbobj"
	"github.com/xintangli/monitor/msql"
	"fmt"*/
)

type SysDataService struct {
	models *models.SysData
}

var SysDataSvc = &SysDataService{
	new(models.SysData),
}

/*

func (h *SysDataService) Page(offset int, limit int,SysData *models.SysData) ([]models.SysData, int64, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	fmt.Println(offset, limit)
	rows, err := dbobj.Query(msql.Sys_Data_sql_list, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	defer rows.Close()
	var dataList []models.SysData
	fmt.Println(rows.Columns())
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(msql.Sys_Data_sql_count).Scan(&total)

	return dataList, total, nil

}

func (h *SysDataService) PageAll(offset int, limit int,SysData *models.SysData) ([]models.SysData, int64, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	fmt.Println(offset, limit)
	rows, err := dbobj.Query(msql.Sys_Data_sql_list, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	defer rows.Close()
	var dataList []models.SysData
	fmt.Println(rows.Columns())
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(msql.Sys_Data_sql_count).Scan(&total)

	return dataList, total, nil

}

func (h *SysDataService) GetNewestData() (models.SysData, error) {
	var data models.SysData
	rows, err :=dbobj.Query(msql.Sys_Data_sql_MaxId)
	if err != nil {
		fmt.Println("query data error", dbobj.GetErrorMsg(err))
		return data, err
	}
	defer rows.Close()
	var dataList []models.SysData
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return data, err
	}
	return dataList[0], nil

}

func (h *SysDataService) Put(sysData models.SysData) error {
	tx, err := dbobj.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}
	*/
/**
	cpu_rate, memory_total, memory_used, memory_free, memory_rate, cached,disk_total,
	 			,disk_used, disk_avail, disk_rate, load_avg1,load_avg5,load_avg15, org_msg
	 *//*

	_, err = tx.Exec(msql.Sys_Data_sql_insert, sysData.MemId, sysData.NodeId, sysData.CpuRate, sysData.MemoryTotal,sysData.MemoryUsed,sysData.MemoryFree,sysData.MemoryRate,
			sysData.Cached, sysData.DiskTotal, sysData.DiskUsed, sysData.DiskAvail, sysData.DiskRate,
			sysData.LoadAvg1, sysData.LoadAvg5,sysData.LoadAvg15, sysData.OrgMsg)
	if err != nil {
		fmt.Println("insert fail err :", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}
*/
