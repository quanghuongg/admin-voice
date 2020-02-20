package model

import "github.com/astaxie/beego/orm"

type Media struct {
	Id int `orm:"column(id);pk" json:"id"`
	Duration int `orm:"column(audio_duration)" json:"duration"`
	Created string `orm:"column(created_at)" json:"created"`
}


func (u *Media) GetMediaByFileId(fileId int) (Media, error) {

	o := orm.NewOrm()

	var media Media
	or := o.Raw("select * from media where file_id = ?", fileId)
	err := or.QueryRow(&media)

	return media, err
}
