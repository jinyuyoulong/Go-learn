package services

import (
	dao2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/dao"
	datasource2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/datasource"
	models2 "github.com/jinyuyoulong/Go-learn/src/iris/superstar/models"
)

// interface 协议 接口
type SuperstarService interface {
	GetAll() []models2.StarInfo
	Get(id int) *models2.StarInfo
	Delete(id int) error
	Update(user *models2.StarInfo, columns []string) error
	Create(user *models2.StarInfo) error
	Search(country string) []models2.StarInfo
}
type superstarService struct {
	dao *dao2.SuperstarDao
}

func NewSuperstarService() *superstarService {
	return &superstarService{
		dao: dao2.NewSuperstarDao(datasource2.InstanceMaster()),
	}
}
func (s *superstarService) GetAll() []models2.StarInfo {
	return s.dao.GetAll()
}
func (s *superstarService) Get(id int) *models2.StarInfo {
	return s.dao.Get(id)
}
func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *superstarService) Update(user *models2.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *superstarService) Create(user *models2.StarInfo) error {
	return s.dao.Create(user)
}
func (s *superstarService) Search(country string) []models2.StarInfo {
	return s.dao.Search(country)
}
