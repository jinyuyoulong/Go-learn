package conf

const DriverName = "mysql"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "333",
	DbName: "projectweb",
}

// 项目大了用到主从的话 需要用到 slave
var SlaveDbConfig DbConf = DbConf{
	Host:   "127.0.0.1",
	Port:   3306,
	User:   "root",
	Pwd:    "333",
	DbName: "projectweb",
}
