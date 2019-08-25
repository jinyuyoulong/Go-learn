package models

type CronTaskHost struct {
	Id     int `xorm:"not null pk autoincr INT(11)"`
	TaskId int `xorm:"not null index INT(11)"`
	HostId int `xorm:"not null index SMALLINT(6)"`
}
