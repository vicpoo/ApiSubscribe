package domain

import "github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"

type OrdenRepository interface {
	Save(orden entities.Orden) error
	Update(id int, orden entities.Orden) error
	GetAll() ([]entities.Orden, error)
	FindByID(id int) (entities.Orden, error)
}
