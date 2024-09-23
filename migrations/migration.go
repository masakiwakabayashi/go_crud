package main

import (
	"go_crud/infra"
	"go_crud/models"
)

func main() {
    infra.Initialize()
		db := infra.SetupDB()

	if err := db.AutoMigrate(&models.Item{}); err != nil {
		panic("Failed to migrate database")
	}
}

