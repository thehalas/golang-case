package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vatansoft-golang-case/models"
)

var DB *gorm.DB

func Connect() {
	connection, err := gorm.Open(mysql.Open("root:root_password@/vatansoft-case"), &gorm.Config{})

	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(
		&models.Product{},
		&models.Property{},
		&models.ProductProperty{},
		&models.Category{},
		&models.Invoice{},
		)
}
