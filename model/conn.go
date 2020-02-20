package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/vtcc/voice-note-admin/config"
)

func init() {

	orm.RegisterDriver("mysql", orm.DRMySQL)
	err := orm.RegisterDataBase("default", "mysql", config.AppConfig().MysqlDsn)

	if err != nil {
		panic("Error init DB Mysql!")
	}
}
