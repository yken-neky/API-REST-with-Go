package ports

import (
	"context"

	"api-rest-with-go/internal/core/domain"
)

// ItemRepository define el contrato para el repositorio de items
type ItemRepository interface {
	Create(ctx context.Context, item *domain.Item) error
	GetAll(ctx context.Context) ([]domain.Item, error)
	GetByID(ctx context.Context, id uint) (*domain.Item, error)
	Update(ctx context.Context, item *domain.Item) error
	Delete(ctx context.Context, id uint) error
}

// Repository agrupa todos los repositorios de la aplicación
type Repository struct {
	Items ItemRepository
}
