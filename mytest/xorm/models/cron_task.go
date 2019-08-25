package models

import (
	"time"
)

type CronTask struct {
	Id               int       `xorm:"not null pk autoincr INT(11)"`
	Name             string    `xorm:"not null VARCHAR(32)"`
	Level            int       `xorm:"not null default 1 index TINYINT(4)"`
	DependencyTaskId string    `xorm:"not null default '' VARCHAR(64)"`
	DependencyStatus int       `xorm:"not null default 1 TINYINT(4)"`
	Spec             string    `xorm:"not null VARCHAR(64)"`
	Protocol         int       `xorm:"not null index TINYINT(4)"`
	Command          string    `xorm:"not null VARCHAR(256)"`
	HttpMethod       int       `xorm:"not null default 1 TINYINT(4)"`
	Timeout          int       `xorm:"not null default 0 MEDIUMINT(9)"`
	Multi            int       `xorm:"not null default 1 TINYINT(4)"`
	RetryTimes       int       `xorm:"not null default 0 TINYINT(4)"`
	RetryInterval    int       `xorm:"not null default 0 SMALLINT(6)"`
	NotifyStatus     int       `xorm:"not null default 1 TINYINT(4)"`
	NotifyType       int       `xorm:"not null default 0 TINYINT(4)"`
	NotifyReceiverId string    `xorm:"not null default '' VARCHAR(256)"`
	NotifyKeyword    string    `xorm:"not null default '' VARCHAR(128)"`
	Tag              string    `xorm:"not null default '' VARCHAR(32)"`
	Remark           string    `xorm:"not null default '' VARCHAR(100)"`
	Status           int       `xorm:"not null default 0 index TINYINT(4)"`
	Created          time.Time `xorm:"not null DATETIME"`
	Deleted          time.Time `xorm:"DATETIME"`
}
