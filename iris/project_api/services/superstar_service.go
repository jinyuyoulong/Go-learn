package services

import (
	"v5u.win/golearn/iris/project_api/dao"
	"v5u.win/golearn/iris/project_api/datasource"
	"v5u.win/golearn/iris/project_api/models"
)

type project_apiService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string) []models.StarInfo
}
type project_apiService struct {
	dao *dao.project_apiDao
}

func Newproject_apiService() *project_apiService {
	return &project_apiService{
		dao: dao.Newproject_apiDao(datasource.InstanceMaster()),
	}
}
func (s *project_apiService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}
func (s *project_apiService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *project_apiService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *project_apiService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *project_apiService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
func (s *project_apiService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
