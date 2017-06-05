package main

import (
	"github.com/xintangli/monitor/server"
	"github.com/xintangli/monitor/msql"
	"github.com/xintangli/monitor/utils"
)


func main()  {
	//init

	//注册ORM映射
	msql.InitORM()
	//初始化日志配置
	utils.InitLogging()
	//启动服务器
	server.InitServer()

}
