package dao

// data access object，DAO
import (
	"log"

	"github.com/go-xorm/xorm"
	"v5u.win/golearn/iris/project_api/models"
)

type project_apiDao struct {
	engine *xorm.Engine
}

func Newproject_apiDao(engine *xorm.Engine) *project_apiDao {

	return &project_apiDao{
		engine: engine,
	}
}

func (s *project_apiDao) Get(id int) *models.StarInfo {
	data := &models.StarInfo{Id: id}
	ok, err := s.engine.Get(data)
	if ok && err == nil {
		return data
	} else {
		data.Id = 0
		return data
	}
}

func (s *project_apiDao) GetAll() []models.StarInfo {
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

func (s *project_apiDao) Delete(id int) error {
	// 假删除
	data := &models.StarInfo{Id: id, SysStatus: 1}
	_, err := s.engine.Id(data.Id).Update(data)

	return err
}

// columns 判断强制更新
func (s *project_apiDao) Update(data *models.StarInfo, columns []string) error {
	_, err := s.engine.Id(data.Id).MustCols(columns...).Update(data)
	// 用到 MustCols 方法
	return err
}

func (s *project_apiDao) Create(data *models.StarInfo) error {
	_, err := s.engine.Insert(data)
	return err
}

func (s *project_apiDao) Search(country string) []models.StarInfo {
	datalist := []models.StarInfo{}
	err := s.engine.Where("country=?", country).Desc("id").Find(&datalist)
	if err != nil {
		log.Println(err)
		return datalist
	}
	return datalist
}
