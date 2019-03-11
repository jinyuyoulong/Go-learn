package datasource

import (
	"sync"

	"github.com/go-xorm/xorm"
)

var (
	masterEngine *xorm.Engine
	slaveEngine  *xorm.Engine
	lock         sync.Mutex
)

func instanceMaster() *xorm.Engine {
	// TODO: 单例的实现
	return masterEngine
}

func instanceSlave() *xorm.Engine {
	// TODO: 单例的实现
	return slaveEngine
}
