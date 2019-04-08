package config

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

// 项目大了用到主从的话 需要用到 slave
var SlaveDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "333",
	DbName: "superstar",
}
