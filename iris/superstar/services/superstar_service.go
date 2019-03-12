package services

import (
	"v5u.win/golearn/iris/superstar/dao"
	"v5u.win/golearn/iris/superstar/datasource"
	"v5u.win/golearn/iris/superstar/models"
)

type SuperstarService interface {
	GetAll() []models.StarInfo
	Get(id int) *models.StarInfo
	Delete(id int) error
	Update(user *models.StarInfo, columns []string) error
	Create(user *models.StarInfo) error
	Search(country string) []models.StarInfo
}
type superstarService struct {
	dao *dao.SuperstarDao
}

func NewSuperstarService() *superstarService {
	return &superstarService{
		dao: dao.NewSuperstarDao(datasource.InstanceMaster()),
	}
}
func (s *superstarService) GetAll() []models.StarInfo {
	return s.dao.GetAll()
}
func (s *superstarService) Get(id int) *models.StarInfo {
	return s.dao.Get(id)
}
func (s *superstarService) Delete(id int) error {
	return s.dao.Delete(id)
}
func (s *superstarService) Update(user *models.StarInfo, columns []string) error {
	return s.dao.Update(user, columns)
}
func (s *superstarService) Create(user *models.StarInfo) error {
	return s.dao.Create(user)
}
func (s *superstarService) Search(country string) []models.StarInfo {
	return s.dao.Search(country)
}
