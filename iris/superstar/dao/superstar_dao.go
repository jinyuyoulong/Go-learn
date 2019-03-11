package repositories

import (
	"github.com/go-xorm/xorm"
)

type SuperstarDao struct {
	engine *xorm.Engine
}
