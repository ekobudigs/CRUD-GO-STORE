package main

import (
	"rest-api-gostore/models"
	"rest-api-gostore/routes"
)

func main() {

	db := models.SetupDB()
	db.AutoMigrate(&models.Product{})
	db.AutoMigrate(&models.Category{})

	r := routes.SetupRoutes(db)
	r.Run()
}