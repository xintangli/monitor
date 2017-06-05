package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/xintangli/monitor/models"
)

type ServicesDataController struct {

}

var ServicesDataCtr = &ServicesDataController{

}

func (s *ServicesDataController) List(c *gin.Context)  {
	c.Request.ParseForm()
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	hostsId, _ := strconv.Atoi(c.PostForm("HostsId"))
	data := &models.ServicesData{
		HostsId: hostsId,
	}
	dataList, err := data.List(limit, offset)
	if err != nil {
		fmt.Println("query ServicesDataPage error :", err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "query list data error"})
		return
	}
	jsByte, err := json.Marshal(dataList)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": string(jsByte)})
}

func (s *ServicesDataController) Page(c *gin.Context)  {
	c.Request.ParseForm()
	offset, _ := strconv.Atoi(c.Param("offset"))
	limit, _ := strconv.Atoi(c.Param("limit"))
	data := &models.ServicesData{}
	dataList, total, err := data.Page(offset, limit)
	if err != nil {
		fmt.Println("query ServicesDataPage error :", err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "query data error"})
		return
	}
	jsByte, err := json.Marshal(dataList)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": string(jsByte), "total": total})
}


func (s *ServicesDataController) Put(c *gin.Context)  {
	var servicesData models.ServicesData
	if c.Bind(&servicesData) == nil {
		fmt.Println(servicesData)
		_, err := servicesData.Insert()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "put data error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "success"})
	}
}


func (s *ServicesDataController) GetNewest(c *gin.Context)  {
	servicesData := &models.ServicesData{}
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	data, err := servicesData.GetNewestData(limit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "GetNewest data error"})
		return
	}
	jsByte, err := json.Marshal(data)
	fmt.Println(string(jsByte))
	c.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "success","data": string(jsByte)})
}


