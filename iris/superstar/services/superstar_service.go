package services

import (
	"v5u.win/GoLearn/iris/superstar/models"
)

// UserService handles CRUID operations of a user datamodel,
// it depends on a user repository for its actions.
// It's here to decouple the data source from the higher level compoments.
// As a result a different repository type can be used with the same logic without any aditional changes.
// It's an interface and it's used as interface everywhere
// because we may need to change or try an experimental different domain logic at the future.
type SuperstarService interface {
	GetAll() []models.StartInfo
	Get(id int) models.StartInfo
	Delete(id int) bool
	Update(user *models.StartInfo) error
	Create(user *models.StartInfo) error
	Serach(country string) []models.StartInfo
}
