package datasource

// 初始化 engine ，创建单例

import (
	"fmt"
	"log"
	"sync"

	_ "github.com/go-sql-driver/mysql" // 使用MySQL的隐式驱动
	"github.com/go-xorm/xorm"
	"github.com/pelletier/go-toml"
	"v5u.win/golearn/iris/projectapi/src/app/config"
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

	driver := config.Conf.Get("database.dirver").(string)
	configTree := config.Conf.Get(driver).(*toml.Tree)
	userName := configTree.Get("databaseUsername").(string)
	password := configTree.Get("databasePassword").(string)
	databaseName := configTree.Get("databaseName").(string)
	dbHost := configTree.Get("databaseHost").(string)
	dbPort := configTree.Get("databasePort").(string)

	connet := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", userName, password, dbHost, dbPort, databaseName)

	fmt.Println(connet)
	var dbDirver string
	dbDirver = config.Conf.Get("database.dirver").(string)
	engine, err := xorm.NewEngine(dbDirver, connet)
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

	c := config.SlaveDbConfig
	connet := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8", c.User, c.Pwd, c.Host, c.Port, c.DbName)
	var dbDirver string
	dbDirver = config.Conf.Get("database.dirver").(string)
	engine, err := xorm.NewEngine(dbDirver, connet)
	// 增加缓存，减少数据库依赖，影响性能
	cacher := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
	engine.SetDefaultCacher(cacher)

	if err != nil {
		log.Fatal("dbhelper.instanceSlace error=%s", err)
	}
	slaveEngine = engine
	return slaveEngine
}
