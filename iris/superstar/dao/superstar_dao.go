package repositories

// data access object，DAO
import (
	"log"

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

func (s *SuperstarDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := s.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (s *SuperstarDao) GetAll() []models.StarInfo {
	// 集合的两种创建方式
	// datalist := make([]models.StartInfo, 0)
	datalist := []models.StarInfo{}
	err := s.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		// return nil 也可以
	}
	return datalist
}

func (s *SuperstarDao) Search(country string) []models.StarInfo {
	datalist := []models.StarInfo{}
	err := s.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	}
	return datalist
}

func (s *SuperstarDao) Delete(id int) error {
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := s.engine.Id(data.Id).Update(data)
	// 假删除
	return err
}

// columns 判断强制更新
func (s *SuperstarDao) Update(data *models.StarInfo, columns []string) error {
	_, err := s.engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

func (s *SuperstarDao) Create(data *models.StarInfo) error {
	_, err := s.engine.Insert(data)
	return err
}
