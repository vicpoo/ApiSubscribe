// main.go
package main

import (
	"encoding/json" // Importar el paquete json
	"fmt"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/vicpoo/ApiSubscribe/src/ApiCocina/domain/entities"
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
	ms, err := infrastructure.NewMessagingService()
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %s", err)
	}
	defer ms.Close()

	// Escuchar eventos de pedidos creados
	go func() {
		msgs, err := ms.ConsumeOrderCreated()
		if err != nil {
			log.Fatalf("Failed to consume order created events: %s", err)
		}

		for msg := range msgs {
			var pedido entities.Orden
			if err := json.Unmarshal(msg, &pedido); err != nil {
				log.Printf("Failed to unmarshal order: %s", err)
				continue
			}

			// Guardar el pedido en la base de datos
			repo := infrastructure.NewMySQLOrdenRepository()
			if err := repo.Save(pedido); err != nil { // Cambio: pasar pedido, no &pedido
				log.Printf("Failed to save order: %s", err)
			}
		}
	}()

	// Inicializar dependencias
	ordenController := infrastructure.InitializeDependencies()

	// Configuración de rutas para Ordenes
	infrastructure.SetupOrdenRoutes(r, ordenController)

	// Mensaje de inicio
	fmt.Println("¡API en Funcionamiento :D!")

	// Iniciar el servidor en el puerto 8000
	err = r.Run(":8080")
	if err != nil {
		fmt.Println("Error al iniciar el servidor:", err)
	}
}
