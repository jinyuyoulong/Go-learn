package repositories

// data access object，DAO
import (
	"github.com/go-xorm/xorm"
	"v5u.win/GoLearn/iris/superstar/models"
)

type SuperstarDao struct {
	engine *xorm.Engine
}

func NewSuperstarDao(engine *xorm.Engine) *SuperstarDao {

	return &SuperstarDao{
		engine: engine,
	}
}

func (s *SuperstarDao) Get(id int) *models.StartInfo {
	return nil
}

func (s *SuperstarDao) GetAll() []*models.StartInfo {
	return nil
}

func (s *SuperstarDao) Search(country string) []*models.StartInfo {
	return nil
}

func (s *SuperstarDao) Delete(id int) error {
	return nil
}

// columns 判断强制更新
func (s *SuperstarDao) Update(data *models.StartInfo, columns []string) error {
	return nil
}

func (s *SuperstarDao) Create(data *models.StartInfo) error {
	return nil
}
