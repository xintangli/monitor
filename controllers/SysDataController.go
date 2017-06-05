package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/xintangli/monitor/models"
	"github.com/xintangli/monitor/params"
	"strings"
	"github.com/xintangli/monitor/utils"
	"runtime/debug"
)

type SysDataController struct {

}

var SysDataCtr = &SysDataController{

}

func (h *SysDataController) Page(c *gin.Context)  {
	var data models.SysData
	draw, _ := strconv.Atoi(c.PostForm("draw"))
	offset, _ := strconv.Atoi(c.PostForm("start"))
	limit, _ := strconv.Atoi(c.PostForm("length"))
	hostsId, _ := strconv.Atoi(c.PostForm("hostsId"))
	if c.Bind(&data) == nil {
		data.HostsId = hostsId
		list, total, err := data.Page(limit, offset)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "page data error"})
			return
		}
		dtParam := models.PageSysData{
			Draw: draw,
			RecordsTotal: total,
			RecordsFiltered: total,
			Data: list,
		}
		c.JSON(http.StatusOK, dtParam)
	}
}

func (h *SysDataController) Insert(c *gin.Context)  {
	var data models.SysData
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

func (s *SysDataController) Msg(c *gin.Context)  {

	defer func() {
		err := recover()
		if err != nil {
			utils.Log.Error(string(debug.Stack()))
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": string(debug.Stack())})
		}

	}()
	buf := make([]byte, 20480)
	n, _ := c.Request.Body.Read(buf)
	msg := string(buf[0:n])
	//msg := c.PostForm("msg")
	if msg == "" {
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "msg is null"})
		return
	}
	msg = strings.Replace(msg, "\t", "", -1)
	//msg = strings.Replace(msg, "\n", "", -1)
	fmt.Println(msg)
	//生成文件
	utils.Log.Info(msg)
	//保存到mysql
	saveToMysql(c, msg)
}

func (s *SysDataController) GetNewestOne(c *gin.Context)  {
	fmt.Println(c.PostForm("hostsId"))
	hostsId, _ := strconv.Atoi(c.PostForm("hostsId"))
	sysData := &models.SysData{
		HostsId: hostsId,
	}
	data, err := sysData.GetNewestOne()
	if err != nil {
		fmt.Println("GetNewest data error", err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "GetNewest data error"})
		return
	}
	jsByte, err := json.Marshal(data)
	fmt.Println(string(jsByte))
	c.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "success","data": string(jsByte)})
}

func saveToMysql(c *gin.Context, msg string)  {
	var monitorParams params.MonitorParams
	json.Unmarshal([]byte(msg), &monitorParams)
	fmt.Println(monitorParams)
	//查询hosts信息
	hosts := &models.SysHosts{
		NodeId: monitorParams.NodeID,
		MemId: monitorParams.MemID,
	}
	hosts.ReadByNodeIdAndMemId()
	if hosts.Id <= 0 {
		fmt.Println("no hosts !", hosts.NodeId, hosts.MemId)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "put sys data error"})
		return
	}

	//系统监控
	var sysData models.SysData
	memRate := monitorParams.Msg.SYSINFO.Mem.MemUsed/monitorParams.Msg.SYSINFO.Mem.MemTotal

	sysData = models.SysData{
		CpuRate: strconv.FormatFloat(monitorParams.Msg.SYSINFO.CPU[0].CPU, 'f', 3, 64),

		HostsId: hosts.Id,
		MemId: monitorParams.MemID,
		NodeId: monitorParams.NodeID,

		MemoryTotal: strconv.FormatFloat(monitorParams.Msg.SYSINFO.Mem.MemTotal, 'f', 3, 64),
		MemoryFree: strconv.FormatFloat(monitorParams.Msg.SYSINFO.Mem.MemFree, 'f', 3, 64),
		MemoryUsed: strconv.FormatFloat(monitorParams.Msg.SYSINFO.Mem.MemUsed, 'f', 3, 64),
		Cached: strconv.FormatFloat(monitorParams.Msg.SYSINFO.Mem.Cached, 'f', 3, 64),
		MemoryRate: strconv.FormatFloat(memRate,'f', 3, 64),

		LoadAvg1: monitorParams.Msg.SYSINFO.Load.Lavg1,
		LoadAvg5: monitorParams.Msg.SYSINFO.Load.Lavg5,
		LoadAvg15: monitorParams.Msg.SYSINFO.Load.Lavg15,

		DiskTotal: monitorParams.Msg.SYSINFO.Disk[0].Size,
		DiskUsed: monitorParams.Msg.SYSINFO.Disk[0].Used,
		DiskAvail: monitorParams.Msg.SYSINFO.Disk[0].Avail,
		DiskRate: monitorParams.Msg.SYSINFO.Disk[0].Use,

		OrgMsg : msg,
	}
	_, err := sysData.Insert()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "put sys data error"})
		return
	}

	//服务监控
	var servicesData models.ServicesData
	qpsSucc, _ := strconv.Atoi(monitorParams.Msg.SERINFO.QPSSucc)
	qpsFail, _ := strconv.Atoi(monitorParams.Msg.SERINFO.QPSFail)
	processNum, _ := strconv.Atoi(monitorParams.Msg.SERINFO.ProcessNum)
	servicesData = models.ServicesData{
		HostsId: hosts.Id,
		MemId: monitorParams.MemID,
		NodeId: monitorParams.NodeID,
		SvcType: hosts.SvcType,
		QpsSucc: qpsSucc,
		QpsFail: qpsFail,
		ProcessNum : processNum,
		RedisUsability : strconv.FormatBool(monitorParams.Msg.SERINFO.RedisUsability),
		RedisMem: monitorParams.Msg.SERINFO.RedisMem,
		RedisSize:monitorParams.Msg.SERINFO.RedisSize,
	}
	_, err = servicesData.Insert()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "put services data error"})
		return
	}

	//业务监控--供方数据
	fmt.Println("TaskIds :", monitorParams.Msg.SUPINFO.SupData.QuerySuccess.TaskID)
	for i, taskId := range monitorParams.Msg.SUPINFO.SupData.QuerySuccess.TaskID {
		busiData := models.BusiData{
			HostsId: hosts.Id,
			SvcType: hosts.SvcType,

			MemId: monitorParams.MemID,
			NodeId: monitorParams.NodeID,
			OrderId: taskId,

			SupQueryTotal: monitorParams.Msg.SUPINFO.SupData.QueryTotal,
			SupQuerySuccess: monitorParams.Msg.SUPINFO.SupData.QuerySuccess.Count[i],

		}

		_, err = busiData.Insert()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "put busi data error"})
			return
		}
	}
	//业务监控--需方数据
	for i, orderId := range monitorParams.Msg.DEMINFO.OrderIds {
		demSuccNum := monitorParams.Msg.DEMINFO.RequestOk[i]
		reqNum := monitorParams.Msg.DEMINFO.RequestOk[i] + monitorParams.Msg.DEMINFO.RequestFail[i]
		busiData := models.BusiData{
			HostsId: hosts.Id,
			SvcType: hosts.SvcType,

			MemId: monitorParams.MemID,
			NodeId: monitorParams.NodeID,
			OrderId: orderId,

			DemSuccNum: demSuccNum ,
			DemReqNum: reqNum,
			//DemCacheNum: cacheNum,
		}

		_, err = busiData.Insert()
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusOK, gin.H{"code": "1099", "msg": "put busi data error"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{"code": "0000", "msg": "success"})
}


