package controllers

import (
	"github.com/gin-gonic/gin"
	/*"net/http"
	"github.com/xintangli/monitor/service"*/
	"fmt"
	/*"encoding/json"
	"strconv"*/
	"github.com/xintangli/monitor/models"
	"net/http"
	"strconv"
	"github.com/xintangli/monitor/utils"
)

type HostsController struct {

}

var HostCtr = &HostsController{}

/*func (h *HostsController) QueryPage(c *gin.Context)  {
	c.Request.ParseForm()
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	svc := &service.HostsService{}
	hosts := &models.SysHosts{}
	hostsList, total, err := svc.Page(offset, limit, hosts)
	if err != nil {
		fmt.Println("query hostPage error :", err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "query data error"})
		return
	}
	jsByte, err := json.Marshal(hostsList)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": string(jsByte), "total": total})
}*/

func (h *HostsController) Insert(c *gin.Context)  {
	var hosts models.SysHosts
	if c.Bind(&hosts) == nil {
		fmt.Println(hosts)
		_, err := hosts.Insert()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "query data error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
	}
}

func (h *HostsController) Read(c *gin.Context)  {
	var hosts models.SysHosts
	if c.Bind(&hosts) == nil {
		fmt.Println(hosts)
		err := hosts.Read()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "read data error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
	}
}

func (h *HostsController) Update(c *gin.Context)  {
	var hosts models.SysHosts
	if c.Bind(&hosts) == nil {
		if hosts.Id == 0 {
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "update, id must"})
			return
		}
		fmt.Println(hosts)
		_, err := hosts.Update()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "update data error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
	}
}


func (h *HostsController) Page(c *gin.Context)  {
	var hosts models.SysHosts
	draw, _ := strconv.Atoi(c.PostForm("draw"))
	offset, _ := strconv.Atoi(c.PostForm("start"))
	limit, _ := strconv.Atoi(c.PostForm("length"))
	fmt.Println(offset,limit,draw)
	err := c.Bind(&hosts)
	if err == nil {
		list, total, err := hosts.Page(limit, offset)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "update data error"})
			return
		}
		dtParam := models.PageSysHosts{
			Draw: draw,
			RecordsTotal: total,
			RecordsFiltered: total,
			Data: list,
		}
		c.JSON(http.StatusOK, dtParam)
	}else {
		utils.Log.Error()
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "binding data error"})
		return
	}
}


