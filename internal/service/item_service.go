package service

import (
	"api-rest-with-go/internal/models"
	"api-rest-with-go/internal/repository"
)

type ItemService struct {
	repo *repository.ItemRepository
}

func NewItemService(repo *repository.ItemRepository) *ItemService {
	return &ItemService{
		repo: repo,
	}
}

func (s *ItemService) CreateItem(item *models.Item) error {
	return s.repo.Create(item)
}

func (s *ItemService) GetAllItems() ([]models.Item, error) {
	return s.repo.GetAll()
}

func (s *ItemService) GetItem(id uint) (*models.Item, error) {
	return s.repo.GetByID(id)
}

func (s *ItemService) UpdateItem(item *models.Item) error {
	return s.repo.Update(item)
}

func (s *ItemService) DeleteItem(id uint) error {
	return s.repo.Delete(id)
} 