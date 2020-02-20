package model

import "github.com/astaxie/beego/orm"

type FeedBack struct {
	Id int `orm:"column(user_id);pk" json:"id"`
	Rate int `orm:"column(rating);pk" json:"rate"`
	Description string `orm:"column(description)" json:"desc"`
}

func (u *FeedBack) GetFeedbackByFileId(fileId int) (FeedBack, error) {

	o := orm.NewOrm()

	var feedback FeedBack
	or := o.Raw("select * from feedback where file_id = ?")
	or = or.SetArgs(fileId)
	err := or.QueryRow(&feedback)

	return feedback, err
}
