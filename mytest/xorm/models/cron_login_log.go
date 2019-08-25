package models

import (
	"time"
)

type CronLoginLog struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Username string    `xorm:"not null VARCHAR(32)"`
	Ip       string    `xorm:"not null VARCHAR(15)"`
	Created  time.Time `xorm:"not null DATETIME"`
}
