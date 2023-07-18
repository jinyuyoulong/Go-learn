package dao

// data access object，DAO
import (
	"log"

	models2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/models"

	"gorm.io/gorm"
)

type SuperstarDao struct {
	engine *gorm.Engine
}

func NewSuperstarDao(engine *gorm.Engine) *SuperstarDao {

	return &SuperstarDao{
		engine: engine,
	}
}

func (s *SuperstarDao) Get(id int) *models2.StarInfo {
	data := &models2.StarInfo{Id: id}
	ok, err := s.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (s *SuperstarDao) GetAll() []models2.StarInfo {
	// 集合的两种创建方式
	// datalist := make([]models.StartInfo, 0)
	datalist := []models2.StarInfo{}
	err := s.engine.Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
		// return nil 也可以
	}
	return datalist
}

func (s *SuperstarDao) Delete(id int) error {
	// 假删除
	data := &models2.StarInfo{Id: id, SysStatus: 1}
	_, err := s.engine.Id(data.Id).Update(data)

	return err
}

// columns 判断强制更新
func (s *SuperstarDao) Update(data *models2.StarInfo, columns []string) error {
	_, err := s.engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

func (s *SuperstarDao) Create(data *models2.StarInfo) error {
	_, err := s.engine.Insert(data)
	return err
}

func (s *SuperstarDao) Search(country string) []models2.StarInfo {
	datalist := []models2.StarInfo{}
	err := s.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	}
	return datalist
}
