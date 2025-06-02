package ports

import (
	"context"

	"api-rest-with-go/internal/core/domain"
)

// ItemService define el contrato para el servicio de items
type ItemService interface {
	CreateItem(ctx context.Context, item *domain.Item) error
	GetAllItems(ctx context.Context) ([]domain.Item, error)
	GetItem(ctx context.Context, id uint) (*domain.Item, error)
	UpdateItem(ctx context.Context, item *domain.Item) error
	DeleteItem(ctx context.Context, id uint) error
}

// Service agrupa todos los servicios de la aplicación
type Service struct {
	Items ItemService
}
