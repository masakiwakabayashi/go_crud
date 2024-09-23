package main

import (
	"go_crud/controllers"
	"go_crud/infra"
	"go_crud/repositories"
	"go_crud/services"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.Initialize()
	db := infra.SetupDB()

	itemRepository := repositories.NewItemRepository(db)
	itemService := services.NewItemService(itemRepository)
	itemController := controllers.NewItemController(itemService)

	r := gin.Default()
	r.POST("/items", itemController.Create)
	r.GET("/items", itemController.FindAll)
	r.GET("/items/:id", itemController.FindById)
	r.PUT("/items/:id", itemController.Update)
	r.DELETE("/items/:id", itemController.Delete)
  r.Run("localhost:8080")
}
