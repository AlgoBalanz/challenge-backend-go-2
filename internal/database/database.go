package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetConnectionGORM() *gorm.DB {
	// Configurar la conexión a la base de datos SQLite
	db, err := gorm.Open(sqlite.Open("courses.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}

	return db
}
