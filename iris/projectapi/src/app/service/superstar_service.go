package service

import (
	"v5u.win/golearn/iris/projectapi/src/app/dao"
	"v5u.win/golearn/iris/projectapi/src/app/datasource"
	"v5u.win/golearn/iris/projectapi/src/app/model"
)

type ProjectapiService interface {
	GetAll() []model.StarInfo
	Get(id int) *model.StarInfo
	Delete(id int) error
	Update(user *model.StarInfo, columns []string) error
	Create(user *model.StarInfo) error
	Search(country string) []model.StarInfo
}
type projectapiService struct {
	dao *dao.ProjectapiDao
}

func NewprojectapiService() *projectapiService {
	return &projectapiService{
		dao: dao.NewProjectapiDao(datasource.InstanceMaster()),
	}
}
func (s *projectapiService) GetAll() []model.StarInfo {
	return s.dao.GetAll()
}
func (s *projectapiService) Get(id int) *model.StarInfo {
	return s.dao.Get(id)
}
func (s *projectapiService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *projectapiService) Update(user *model.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *projectapiService) Create(user *model.StarInfo) error {
	return s.dao.Create(user)
}
func (s *projectapiService) Search(country string) []model.StarInfo {
	return s.dao.Search(country)
}
