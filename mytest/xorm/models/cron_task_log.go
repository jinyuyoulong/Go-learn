package models

import (
	"time"
)

type CronTaskLog struct {
	Id         int64     `xorm:"pk autoincr BIGINT(20)"`
	TaskId     int       `xorm:"not null default 0 index INT(11)"`
	Name       string    `xorm:"not null VARCHAR(32)"`
	Spec       string    `xorm:"not null VARCHAR(64)"`
	Protocol   int       `xorm:"not null index TINYINT(4)"`
	Command    string    `xorm:"not null VARCHAR(256)"`
	Timeout    int       `xorm:"not null default 0 MEDIUMINT(9)"`
	RetryTimes int       `xorm:"not null default 0 TINYINT(4)"`
	Hostname   string    `xorm:"not null default '' VARCHAR(128)"`
	StartTime  time.Time `xorm:"DATETIME"`
	EndTime    time.Time `xorm:"DATETIME"`
	Status     int       `xorm:"not null default 1 index TINYINT(4)"`
	Result     string    `xorm:"not null MEDIUMTEXT"`
}
