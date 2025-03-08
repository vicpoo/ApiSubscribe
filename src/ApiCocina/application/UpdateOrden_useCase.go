package application

import (
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"
)

type UpdateOrdenUseCase struct {
	repo domain.OrdenRepository
}

func NewUpdateOrdenUseCase(repo domain.OrdenRepository) *UpdateOrdenUseCase {
	return &UpdateOrdenUseCase{repo: repo}
}

func (uc *UpdateOrdenUseCase) Execute(id int, orden entities.Orden) error {
	return uc.repo.Update(id, orden)
}
