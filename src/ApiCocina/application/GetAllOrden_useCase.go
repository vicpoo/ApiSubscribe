// GetAllOrden_useCase.go
package application

import (
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"
)

type GetAllOrdenUseCase struct {
	repo domain.OrdenRepository
}

func NewGetAllOrdenUseCase(repo domain.OrdenRepository) *GetAllOrdenUseCase {
	return &GetAllOrdenUseCase{repo: repo}
}

func (uc *GetAllOrdenUseCase) Execute() ([]entities.Orden, error) {
	return uc.repo.GetAll()
}
