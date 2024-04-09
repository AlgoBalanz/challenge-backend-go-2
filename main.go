package main

import (
	"challenge/internal/database"
	"challenge/internal/handlers"
	"challenge/internal/models"
	"challenge/internal/router"

	"github.com/gin-gonic/gin"
)

func main() {

	db := database.GetConnectionGORM()
	// Migrar el esquema de la base de datos
	if err := db.AutoMigrate(&models.Course{}); err != nil {
		panic(err)
	}

	h := handlers.NewHandler(db)

	// Iniciar el servidor Gin
	r := gin.Default()

	router.RegisterRoutes(r, h)

	r.Run()
}
