package services

import (
	"context"

	"api-rest-with-go/internal/core/domain"
	"api-rest-with-go/internal/core/ports"
)

type itemService struct {
	repo ports.ItemRepository
}

func NewItemService(repo ports.ItemRepository) ports.ItemService {
	return &itemService{
		repo: repo,
	}
}

func (s *itemService) CreateItem(ctx context.Context, item *domain.Item) error {
	if err := item.Validate(); err != nil {
		return err
	}
	item.SetCreateTime()
	return s.repo.Create(ctx, item)
}

func (s *itemService) GetAllItems(ctx context.Context) ([]domain.Item, error) {
	return s.repo.GetAll(ctx)
}

func (s *itemService) GetItem(ctx context.Context, id uint) (*domain.Item, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *itemService) UpdateItem(ctx context.Context, item *domain.Item) error {
	if err := item.Validate(); err != nil {
		return err
	}
	item.SetUpdateTime()
	return s.repo.Update(ctx, item)
}

func (s *itemService) DeleteItem(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
