package main

import (
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/infrastructure"
	"github.com/vicpoo/ApiSubscribe/src/core"
)

func main() {
	// Inicializar la base de datos
	core.InitDB()

	// Crear un router con Gin
	r := gin.Default()

	// Configuración de CORS
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowCredentials = true
	corsConfig.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	corsConfig.ExposeHeaders = []string{"Content-Length"}

	r.Use(cors.New(corsConfig))

	// Inicializar dependencias
	ordenController := infrastructure.InitializeDependencies()

	// Configuración de rutas para Ordenes
	infrastructure.SetupOrdenRoutes(r, ordenController)

	// Mensaje de inicio
	fmt.Println("¡API en Funcionamiento :D!")

	// Iniciar el servidor en el puerto 8000
	err := r.Run(":8000")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
