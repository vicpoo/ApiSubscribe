// dependencies.go
package infrastructure

import (
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/application"
)

func InitializeDependencies() *OrdenController {
	repo := NewMySQLOrdenRepository()
	updateOrdenUseCase := application.NewUpdateOrdenUseCase(repo)
	getAllOrdenUseCase := application.NewGetAllOrdenUseCase(repo)
	return NewOrdenController(updateOrdenUseCase, getAllOrdenUseCase)
}
