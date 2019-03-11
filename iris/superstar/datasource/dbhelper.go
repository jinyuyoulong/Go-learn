package datasource

import (
	"fmt"
	"log"
	"sync"

	"github.com/go-xorm/xorm"
	"v5u.win/GoLearn/iris/superstar/conf"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func instanceMaster() *xorm.Engine {

	if masterEngine != nil {
		return masterEngine
	}
	// WARNNING:
	// 互斥锁，如果有多个线程同时访问，且masterEngine == nil 的时候，第一线程创建完后，后续线程依然认为 masterEngine == nil,此时需要在多一次判断
	lock.Lock()
	defer lock.Unlock()

	if masterEngine != nil {
		return masterEngine
	}

	c := conf.MasterDbConfig
	connet := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, connet)
	if err != nil {
		log.Fatal("dbhelper.instanceMaster error=%s", err)
	}
	masterEngine = engine
	return masterEngine
}

func instanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf.SlaveDbConfig
	connet := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf.DriverName, connet)
	if err != nil {
		log.Fatal("dbhelper.instanceSlace error=%s", err)
	}
	slaveEngine = engine
	return slaveEngine
}
