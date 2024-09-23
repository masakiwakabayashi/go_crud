package controllers

import (
	"go_crud/dto"
	"go_crud/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IItemController interface {
	Create(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type ItemController struct {
	service services.IItemService
}

func NewItemController(service services.IItemService) IItemController {
	return &ItemController{service: service}
}

func (c *ItemController) Create(ctx *gin.Context) {
	var input dto.CreateItemInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newItem, err := c.service.Create(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newItem})
}

func (c ItemController) FindAll(ctx *gin.Context) {
	items, err := c.service.FindAll()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Unexpected error"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": items})
}
