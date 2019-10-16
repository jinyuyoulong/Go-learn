package datasource

// 初始化 engine ，创建单例

import (
	"fmt"
	conf2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/conf"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" // 使用MySQL的隐式驱动
	"github.com/go-xorm/xorm"
	//  import cycle is not allowed
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func InstanceMaster() *xorm.Engine {

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

	c := conf2.MasterDbConfig
	connet := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf2.DriverName, connet)
	if err != nil {
		log.Fatal("dbhelper.instanceMaster error=%s", err)
	}
	masterEngine = engine
	return masterEngine
}

func InstanceSlave() *xorm.Engine {
	if slaveEngine != nil {
		return slaveEngine
	}
	lock.Lock()
	defer lock.Unlock()

	if slaveEngine != nil {
		return slaveEngine
	}

	c := conf2.SlaveDbConfig
	connet := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	engine, err := xorm.NewEngine(conf2.DriverName, connet)
	// 增加缓存，减少数据库依赖，影响性能
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	if err != nil {
		log.Fatal("dbhelper.instanceSlace error=%s", err)
	}
	slaveEngine = engine
	return slaveEngine
}
