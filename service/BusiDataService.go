package service

import (
	"github.com/xintangli/monitor/models"
	/*"github.com/hzwy23/dbobj"
	"github.com/xintangli/monitor/msql"
	"fmt"*/
)

type BusiDataService struct {
	models *models.BusiData
}

var BusiDataSvc = &BusiDataService{
	new(models.BusiData),
}


/*func (h *BusiDataService) Page(offset int, limit int,busiData *models.BusiData) ([]models.BusiData, int64, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	fmt.Println(busiData)
	fmt.Println(offset, limit, busiData.SvcType)
	rows, err := dbobj.Query(msql.Busi_Data_sql_list, busiData.SvcType, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	defer rows.Close()
	var dataList []models.BusiData
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(msql.Busi_Data_sql_count).Scan(&total)

	return dataList, total, nil

}

func (h *BusiDataService) GetNewestData() (models.BusiData, error) {
	var data models.BusiData
	rows, err :=dbobj.Query(msql.Busi_Data_sql_MaxId, "1")
	if err != nil {
		fmt.Println("query data error", dbobj.GetErrorMsg(err))
		return data, err
	}
	defer rows.Close()
	var dataList []models.BusiData
	err = dbobj.Scan(rows, &dataList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return data, err
	}
	return dataList[0], nil

}

func (h *BusiDataService) Put(busiLogData models.BusiData) error {
	tx, err := dbobj.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tx.Exec(msql.Busi_Data_sql_insert, busiLogData.HostsId, busiLogData.MemId, busiLogData.NodeId, busiLogData.SvcType, busiLogData.DemSuccNum,
			busiLogData.DemReqNum, busiLogData.DemCacheNum, busiLogData.SupQueryTotal, busiLogData.SupQuerySuccess)
	if err != nil {
		fmt.Println("insert fail err :", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}*/
