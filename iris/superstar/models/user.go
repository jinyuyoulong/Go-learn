package models

type User struct {
	Id       int    `xorm:"not null pk autoincr comment('主键ID') INT(10)"`
	Username string `xorm:"not null comment('用户名') VARCHAR(50)"`
	Age      int    `xorm:"not null default 0 comment('年龄') TINYINT(1)"`
	Email    string `xorm:"not null comment('邮箱') VARCHAR(50)"`
	Pwd      string `xorm:"not null comment('密码') VARCHAR(50)"`
}
