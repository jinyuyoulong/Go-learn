package models

import (
	"time"
)

type JdcronTasks struct {
	Id         int       `xorm:"not null pk INT(11)"`
	TaskId     int       `xorm:"not null INT(11)"`
	Spec       string    `xorm:"VARCHAR(255)"`
	Command    string    `xorm:"VARCHAR(255)"`
	Status     int       `xorm:"not null default 0 TINYINT(1)"`
	CreateTime time.Time `xorm:"not null DATETIME"`
}
