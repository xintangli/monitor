package service

import (
	"github.com/xintangli/monitor/models"
	/*"github.com/hzwy23/dbobj"
	"github.com/xintangli/monitor/msql"
	"fmt"*/
)

type MonitorDataService struct {
	models *models.MonitorData
}

var MonitorDataSvc = &MonitorDataService{
	new(models.MonitorData),
}


/*func (h *MonitorDataService) Page(offset int, limit int,monitorData *models.MonitorData) ([]models.MonitorData, int64, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	fmt.Println(offset, limit)
	rows, err := dbobj.Query(msql.Sys_MonitorData_sql_list, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	defer rows.Close()
	var dataList []models.MonitorData
	fmt.Println(rows.Columns())
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(msql.Sys_MonitorData_sql_count).Scan(&total)

	return dataList, total, nil

}

func (h *MonitorDataService) GetNewestData() (models.MonitorData, error) {
	var data models.MonitorData
	rows, err :=dbobj.Query(msql.Sys_MonitorData_sql_MaxId)
	if err != nil {
		fmt.Println("query data error", dbobj.GetErrorMsg(err))
		return data, err
	}
	defer rows.Close()
	var dataList []models.MonitorData
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return data, err
	}
	return dataList[0], nil

}

func (h *MonitorDataService) Put(monitorData models.MonitorData) error {
	tx, err := dbobj.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tx.Exec(msql.Sys_MonitorData_sql_insert, monitorData.HostsId, monitorData.Data)
	if err != nil {
		fmt.Println("insert fail err :", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}*/
