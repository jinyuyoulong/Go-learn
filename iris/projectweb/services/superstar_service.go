package services

import (
	"v5u.win/golearn/iris/projectweb/dao"
	"v5u.win/golearn/iris/projectweb/datasource"
	"v5u.win/golearn/iris/projectweb/models"
)

type projectwebService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string) []models.StarInfo
}
type projectwebService struct {
	dao *dao.projectwebDao
}

func NewprojectwebService() *projectwebService {
	return &projectwebService{
		dao: dao.NewprojectwebDao(datasource.InstanceMaster()),
	}
}
func (s *projectwebService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}
func (s *projectwebService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *projectwebService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *projectwebService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *projectwebService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
func (s *projectwebService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
