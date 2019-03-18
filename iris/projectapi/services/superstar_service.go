package services

import (
	"v5u.win/golearn/iris/projectapi/dao"
	"v5u.win/golearn/iris/projectapi/datasource"
	"v5u.win/golearn/iris/projectapi/models"
)

type ProjectapiService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string) []models.StarInfo
}
type projectapiService struct {
	dao *dao.ProjectapiDao
}

func NewprojectapiService() *projectapiService {
	return &projectapiService{
		dao: dao.NewProjectapiDao(datasource.InstanceMaster()),
	}
}
func (s *projectapiService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}
func (s *projectapiService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *projectapiService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *projectapiService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *projectapiService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
func (s *projectapiService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
