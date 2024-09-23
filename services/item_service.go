package services

import (
	"go_crud/dto"
	"go_crud/models"
	"go_crud/repositories"
)

type IItemService interface {
	Create(createItemInput dto.CreateItemInput) (*models.Item, error)
	FindAll() (*[]models.Item, error)
}

type ItemService struct {
	repository repositories.IItemRepository
}

func NewItemService(repository repositories.IItemRepository) IItemService {
	return &ItemService{repository: repository}
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}
