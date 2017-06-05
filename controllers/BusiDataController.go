package controllers

import (
	"strconv"
	"fmt"
	"net/http"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/xintangli/monitor/models"
	"github.com/xintangli/monitor/params"
	"github.com/xintangli/monitor/utils"
)

type BusiDataController struct {

}

var BusiDataCtr = &BusiDataController{

}

func (b *BusiDataController) CharList(c *gin.Context)  {
	c.Request.ParseForm()
	offset, _ := strconv.Atoi(c.PostForm("offset"))
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	hostsId, _ := strconv.Atoi(c.PostForm("hostsId"))

	hosts := &models.SysHosts{
		Id: hostsId,
	}
	if hostsId <= 0 {
		fmt.Println("hosts id is null")
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "hosts id is null"})
		return
	}
	hosts.Read();

	data := &models.BusiData{
		NodeId: c.PostForm("nodeId"),
	}
	//根据节点ID查询orderid及taskid
	mo := models.MemberOrder{
		SvcType: hosts.SvcType,
		MemId: hosts.MemId,
	}
	moList, err := mo.GetListByMemIdAndSvcType();
	if err != nil {
		fmt.Println("GetListByMemIdAndSvcType error ", err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "GetListByMemIdAndSvcType error"})
		return
	}
	data.SvcType = hosts.SvcType
	dataMap := make(map[string][]models.BusiData)
	for _, m := range moList {
		if m.SvcType == 1 {
			tempList, _ := data.List(limit, offset, m.SvcType, m.OrderId)
			if tempList != nil {
				dataMap[m.OrderId] = tempList
			}
		}else{
			tempList, _ := data.List(limit, offset, m.SvcType, m.TaskId)
			if tempList != nil {
				dataMap[m.TaskId] = tempList
			}
		}
	}
	if err != nil {
		fmt.Println("query BusiData error :", err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "query list data error"})
		return
	}
	jsByte, err := json.Marshal(dataMap)
	c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "data": string(jsByte)})
}

func (b *BusiDataController) Page(c *gin.Context)  {
	var data models.BusiData
	draw, _ := strconv.Atoi(c.PostForm("draw"))
	offset, _ := strconv.Atoi(c.PostForm("start"))
	limit, _ := strconv.Atoi(c.PostForm("length"))
	hostsId, _ := strconv.Atoi(c.PostForm("hostsId"))
	fmt.Println(offset,limit,draw, hostsId)
	data.HostsId = hostsId
	list, total, err := data.Page(limit, offset)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "page data error"})
		return
	}
	dtParam := models.PageBusiData{
		Draw: draw,
		RecordsTotal: total,
		RecordsFiltered: total,
		Data: list,
	}
	/*json, err := json.Marshal(&dtParam)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "marshal data error"})
		return
	}*/
	c.JSON(http.StatusOK, &dtParam)
	return

}

func (b *BusiDataController) Insert(c *gin.Context)  {
	var data models.BusiData
	if c.Bind(&data) == nil {
		fmt.Println(data)
		_, err := data.Insert()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "insert data error"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"code": http.StatusOK, "msg": "success"})
	}
}

func (b *BusiDataController) GetNewest(c *gin.Context)  {
	data := &models.BusiData{}
	limit, _ := strconv.Atoi(c.PostForm("limit"))
	dataList, err := data.GetNewestData(limit)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "GetNewest data error"})
		return
	}
	jsByte, err := json.Marshal(dataList)
	fmt.Println(string(jsByte))
	c.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "success","data": string(jsByte)})
}
/**
	外部查询日志接口
 */
func (b *BusiDataController) ListOuter(c *gin.Context)  {
	var data models.BusiData
	var req params.OuterBusiReq
	err := c.Bind(&req)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "binding data error"})
		return
	}
	key := ""
	//登录验证
	utils.LoginVerify(req.MemId, req.Timestamp, req.Sign, key)
	//参数验证
	err = req.ParamVerify()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": err})
		return
	}
	list, err := data.ListOuter(req.StartTime, req.EndTime, req.MemId, req.OrderId)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "page data error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": "0000", "msg": list})
	return

}



