package server

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/xintangli/monitor/controllers"
)

var Engine *gin.Engine


func InitServer() error {
	//初始化中间件
	//middleware.InitMiddleware()

	Engine = gin.New()

	//进入基础菜单router
	menuRouter()
	//系统内部 router
	innerRouter()
	//外部api router
	outerRouter()
	
	return Engine.Run(":8080")
}

func menuRouter()  {
	
	Engine.LoadHTMLGlob("templates/*")
	Engine.Static("/static", "./static")
	
	Engine.GET("/", index)
	Engine.GET("/busi", busi)
	Engine.GET("/hosts", hosts)
	Engine.GET("/sys", sys)

}

func innerRouter()  {
	
	innerGroup := Engine.Group("/api/v1")
	{
		//主机
		hostGroup := innerGroup.Group("/hosts")
		{
			hostGroup.POST("/i", controllers.HostCtr.Insert)
			hostGroup.POST("/r", controllers.HostCtr.Read)
			hostGroup.POST("/u", controllers.HostCtr.Update)
			hostGroup.POST("/p", controllers.HostCtr.Page)//分页
		}
		//系统
		sysDataGroup := innerGroup.Group("/sysDatas")
		{
			sysDataGroup.POST("/page", controllers.SysDataCtr.Page)
			sysDataGroup.GET("/page", controllers.SysDataCtr.Page)

			sysDataGroup.POST("", controllers.SysDataCtr.Insert)
			sysDataGroup.POST("/newest", controllers.SysDataCtr.GetNewestOne)//获取最近条信息
			sysDataGroup.GET("/newest", controllers.SysDataCtr.GetNewestOne)

			sysDataGroup.POST("/msg", controllers.SysDataCtr.Msg)
		}

		//服务监控
		servicesGroup := innerGroup.Group("/servicesDatas")
		{
			servicesGroup.POST("/list", controllers.ServicesDataCtr.List)
			servicesGroup.GET("/list", controllers.ServicesDataCtr.List)
		}

		//业务日志监控
		busiGroup := innerGroup.Group("/busiDatas")
		{
			busiGroup.POST("", controllers.BusiDataCtr.Page)
			busiGroup.GET("", controllers.BusiDataCtr.Page)
			busiGroup.POST("/charList", controllers.BusiDataCtr.CharList)
			busiGroup.GET("/charList", controllers.BusiDataCtr.CharList)
		}
	}
}

func outerRouter()  {
	outerGroup := Engine.Group("/api/v2")
	busiGroup := outerGroup.Group("/busiDatas")

	busiGroup.POST("", controllers.BusiDataCtr.ListOuter)
}


//实时监控主页
func index(c *gin.Context)  {
	//c.String(http.StatusOK, "TalkingData owl api, use /apidoc to show the detail information")
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "Main website",
	})
}
//日志监控主页
func busi(c *gin.Context)  {
	//c.String(http.StatusOK, "TalkingData owl api, use /apidoc to show the detail information")
	c.HTML(http.StatusOK, "busi.html", gin.H{
		"title": "Main website",
	})
}
//主机管理主页
func hosts(c *gin.Context)  {
	//c.String(http.StatusOK, "TalkingData owl api, use /apidoc to show the detail information")
	c.HTML(http.StatusOK, "hosts.html", gin.H{})
}
//主机管理主页
func sys(c *gin.Context)  {
	//c.String(http.StatusOK, "TalkingData owl api, use /apidoc to show the detail information")
	c.HTML(http.StatusOK, "sys.html", gin.H{})
}
