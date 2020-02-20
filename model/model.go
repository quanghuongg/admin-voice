package model

import "github.com/astaxie/beego/orm"

func init() {

	orm.RegisterModel(new(User))
	orm.RegisterModel(new(Record))
}

func BuildSql(sql ...string) string {

	str := ""
	for _, v := range sql {
		str += v + " "
	}

	return str
}