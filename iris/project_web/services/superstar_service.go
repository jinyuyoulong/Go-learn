package services

import (
	"v5u.win/golearn/iris/project_web/dao"
	"v5u.win/golearn/iris/project_web/datasource"
	"v5u.win/golearn/iris/project_web/models"
)

type project_webService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string) []models.StarInfo
}
type project_webService struct {
	dao *dao.project_webDao
}

func Newproject_webService() *project_webService {
	return &project_webService{
		dao: dao.Newproject_webDao(datasource.InstanceMaster()),
	}
}
func (s *project_webService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}
func (s *project_webService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *project_webService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *project_webService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *project_webService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
func (s *project_webService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
