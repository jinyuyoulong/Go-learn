package models

import (
	"time"
)

type CronUser struct {
	Id       int       `xorm:"not null pk autoincr INT(11)"`
	Name     string    `xorm:"not null unique VARCHAR(32)"`
	Password string    `xorm:"not null CHAR(32)"`
	Salt     string    `xorm:"not null CHAR(6)"`
	Email    string    `xorm:"not null default '' unique VARCHAR(50)"`
	Created  time.Time `xorm:"not null DATETIME"`
	Updated  time.Time `xorm:"DATETIME"`
	IsAdmin  int       `xorm:"not null default 0 TINYINT(4)"`
	Status   int       `xorm:"not null default 1 TINYINT(4)"`
}
