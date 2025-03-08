// orden_controller.go
package infrastructure

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/application"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"
)

type OrdenController struct {
	updateOrdenUseCase *application.UpdateOrdenUseCase
	getAllOrdenUseCase *application.GetAllOrdenUseCase
}

func NewOrdenController(updateOrdenUseCase *application.UpdateOrdenUseCase, getAllOrdenUseCase *application.GetAllOrdenUseCase) *OrdenController {
	return &OrdenController{
		updateOrdenUseCase: updateOrdenUseCase,
		getAllOrdenUseCase: getAllOrdenUseCase,
	}
}

func (oc *OrdenController) UpdateOrden(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
		return
	}

	var orden entities.Orden
	if err := c.ShouldBindJSON(&orden); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := oc.updateOrdenUseCase.Execute(id, orden); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Orden updated successfully"})
}

func (oc *OrdenController) GetAllOrdenes(c *gin.Context) {
	ordenes, err := oc.getAllOrdenUseCase.Execute()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, ordenes)
}
