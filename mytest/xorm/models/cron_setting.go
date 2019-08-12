package models

type CronSetting struct {
	Id    int    `xorm:"not null pk autoincr INT(11)"`
	Code  string `xorm:"not null VARCHAR(32)"`
	Key   string `xorm:"not null VARCHAR(64)"`
	Value string `xorm:"not null default '' VARCHAR(4096)"`
}
