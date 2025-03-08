package infrastructure

import (
	"github.com/gin-gonic/gin"
)

func SetupOrdenRoutes(r *gin.Engine, ordenController *OrdenController) {
	r.PUT("/ordenes/:id", ordenController.UpdateOrden)
	r.GET("/ordenes", ordenController.GetAllOrdenes)
}
