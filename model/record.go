package model

import (
	"github.com/astaxie/beego/orm"
)

type Record struct {
	Id int `orm:"column(file_id);pk" json:"id"`
	UserId int `orm:"column(user_id)" json:"user_id"`
	Title string `orm:"column(file_name)" json:"title"`
	Content string `orm:"column(file_content)" json:"content"`
	Created string `orm:"column(created_at)" json:"created"`
	Duration int `orm:"column(file_duration)" json:"duration"`

	User User `orm:"-" json:"user"`
	FeedBack FeedBack `orm:"-" json:"feedback"`
	HashId string `orm:"-" json:"hashid"`
	AudioUrl string `orm:"-" json:"audioUrl"`
}

func (r *Record) GetAllRecords(pos int, size int, query string, userId string) ([]Record, int, error) {

	o := orm.NewOrm()

	count := "select count(*) from file"
	sql := "select * from file"
	where := ""
	order := "order by file_id desc"
	limit := "limit ?,?"

	rawSql := ""
	params := []interface{}{}
	rawSql = BuildSql(where, order)
	if len(query) > 0 {
		query = "%" + query + "%"
		where := "where file_name like ?"
		rawSql = BuildSql(where, order)
		params = append(params, query)
	}
	if len(userId) > 0  {
		where := "where user_id = ?"
		rawSql = BuildSql(where, order)
		params = append(params, userId)
	}

	or := o.Raw(BuildSql(count, rawSql), params)
	var total int
	err := or.QueryRow(&total)

	params = append(params, pos, size)
	or = o.Raw(BuildSql(sql, rawSql, limit), params)
	var records []Record
	_, err = or.QueryRows(&records)

	return records, total, err
}

func (r *Record) Count() (int, error) {

	o := orm.NewOrm()

	var total int
	or := o.Raw("select count(*) from file")
	err := or.QueryRow(&total)

	return total, err
}

func (r *Record) CountByUserId(userId int) (int, error) {

	o := orm.NewOrm()

	var total int
	or := o.Raw("select count(*) from file where user_id = ?", userId)
	err := or.QueryRow(&total)

	return total, err
}

func (r *Record) GetByFileId(fileId int) (Record, error) {

	o := orm.NewOrm()

	var record Record
	or := o.Raw("select * from file where file_id = ?", fileId)
	err := or.QueryRow(&record)

	return record, err
}
