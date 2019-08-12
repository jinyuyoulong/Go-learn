package models

type CronHost struct {
	Id     int    `xorm:"not null pk autoincr SMALLINT(6)"`
	Name   string `xorm:"not null VARCHAR(64)"`
	Alias  string `xorm:"not null default '' VARCHAR(32)"`
	Port   int    `xorm:"not null default 5921 INT(11)"`
	Remark string `xorm:"not null default '' VARCHAR(100)"`
}
