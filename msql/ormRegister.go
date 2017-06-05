package msql

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xintangli/monitor/models"
)

func InitORM()  {
	orm.Debug = true
	orm.RegisterDriver("mysql", orm.DRMySQL)

	orm.RegisterDataBase("default", "mysql", "root:abcd_123@/monitor?charset=utf8", 30)

	// register model
	orm.RegisterModel(new(models.SysHosts))
	orm.RegisterModel(new(models.SysData))
	orm.RegisterModel(new(models.ServicesData))
	orm.RegisterModel(new(models.BusiData))
	orm.RegisterModel(new(models.MemberOrder))
	orm.RegisterModel(new(models.User))

}



