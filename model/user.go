package model

import "github.com/astaxie/beego/orm"

type User struct {
	Id int `orm:"column(user_id);pk" json:"id"`
	Avatar string `orm:"column(avatar)" json:"avatar"`
	FullName string `orm:"column(full_name)" json:"full_name"`
	Email string `orm:"column(email)" json:"email"`
	Phone string `orm:"column(phone)" json:"phone"`
	Password string `orm:"column(password)" json:"-"`
	Created string `orm:"column(created_at)" json:"created"`
	RoleId int `orm:"column(role_id)" json:"-"`

	TotalRecord int `orm:"-" json:"totalRecords"`
}

func (r *User) GetAllUsers(pos int, size int, query string) ([]User, int, error) {

	o := orm.NewOrm()

	count := "select count(*) from user"
	sql := "select * from user"
	where := ""
	order := "order by user_id desc"
	limit := "limit ?,?"

	rawSql := ""
	params := make([]interface{}, 0)
	rawSql = BuildSql(where, order)
	if len(query) > 0 {
		query = "%" + query + "%"
		where := "where full_name like ? or email like ? or phone like ?"
		rawSql = BuildSql(where, order)
		params = append(params, query, query, query)
	}

	or := o.Raw(BuildSql(count, rawSql), params)
	var total int
	err := or.QueryRow(&total)

	params = append(params, pos, size)
	or = o.Raw(BuildSql(sql, rawSql, limit), params)
	var users []User
	_, err = or.QueryRows(&users)

	return users, total, err
}

func (r *User) Count() (int, error) {

	o := orm.NewOrm()

	var total int
	or := o.Raw("select count(*) from user")
	err := or.QueryRow(&total)

	return total, err
}

func (u *User) GetUserById(userId int) (User, error) {

	o := orm.NewOrm()

	var user User
	or := o.Raw("select * from user where user_id = ?")
	or = or.SetArgs(userId)
	err := or.QueryRow(&user)

	return user, err
}

func (u *User) GetUserByPhoneOrEmail(email string) (User, error) {

	o := orm.NewOrm()

	var user User
	or := o.Raw("select * from user where email = ? or phone = ?")
	or = or.SetArgs(email, email)
	err := or.QueryRow(&user)

	return user, err
}
