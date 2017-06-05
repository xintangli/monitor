package service

import (
	/*"github.com/gin-gonic/gin"*/
	"github.com/xintangli/monitor/models"
	/*"github.com/hzwy23/dbobj"
	"github.com/xintangli/monitor/msql"
	"fmt"*/
)

type HostsService struct {
	models *models.SysHosts
}
/*
func (h *HostsService) Page(offset int, limit int,hosts *models.SysHosts) ([]models.SysHosts, int64, error) {
	if offset <= 0 {
		offset = 0
	}
	if limit <= 0 {
		limit = 10
	}
	rows, err := dbobj.Query(msql.Sys_hosts_sql_list, offset, limit)
	if err != nil {
		fmt.Println("query data error :", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	defer rows.Close()
	var hostsList []models.SysHosts
	err = dbobj.Scan(rows, &hostsList)
	if err != nil {
		fmt.Println("query data error.", dbobj.GetErrorMsg(err))
		return nil, 0, err
	}
	var total int64 = 0
	dbobj.QueryRow(msql.Sys_hosts_sql_count).Scan(&total)

	return hostsList, total, nil

}

func (h *HostsService) Get(c *gin.Context)  {

}

func (h *HostsService) Put(name string, inIp string, outIp string, cpu string, memory string, disk string, status string) error {
	tx, err := dbobj.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = tx.Exec(msql.Sys_hosts_sql_insert, name, inIp, outIp, cpu, memory, disk, status)
	if err != nil {
		fmt.Println("insert fail err :", err)
		tx.Rollback()
		return err
	}
	return tx.Commit()
}

func (h *HostsService) Post(c *gin.Context)  {

}*/
