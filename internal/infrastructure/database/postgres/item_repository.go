package postgres

import (
	"context"
	"errors"

	"api-rest-with-go/internal/core/domain"

	"gorm.io/gorm"
)

type itemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *itemRepository {
	return &itemRepository{
		db: db,
	}
}

func (r *itemRepository) Create(ctx context.Context, item *domain.Item) error {
	return r.db.WithContext(ctx).Create(item).Error
}

func (r *itemRepository) GetAll(ctx context.Context) ([]domain.Item, error) {
	var items []domain.Item
	err := r.db.WithContext(ctx).Find(&items).Error
	return items, err
}

func (r *itemRepository) GetByID(ctx context.Context, id uint) (*domain.Item, error) {
	var item domain.Item
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var count int64
			r.db.WithContext(ctx).Unscoped().Model(&domain.Item{}).Where("id = ? AND deleted_at IS NOT NULL", id).Count(&count)
			if count > 0 {
				return nil, errors.New("item has been deleted")
			}
			return nil, errors.New("item not found")
		}
		return nil, err
	}
	return &item, nil
}

func (r *itemRepository) Update(ctx context.Context, item *domain.Item) error {
	result := r.db.WithContext(ctx).Model(&domain.Item{}).Where("id = ? AND deleted_at IS NULL", item.ID).Updates(map[string]interface{}{
		"name":        item.Name,
		"description": item.Description,
		"price":       item.Price,
		"updated_at":  item.UpdatedAt,
	})

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		var count int64
		r.db.WithContext(ctx).Unscoped().Model(&domain.Item{}).Where("id = ? AND deleted_at IS NOT NULL", item.ID).Count(&count)
		if count > 0 {
			return errors.New("item has been deleted")
		}
		return errors.New("item not found")
	}

	return nil
}

func (r *itemRepository) Delete(ctx context.Context, id uint) error {
	var item domain.Item
	err := r.db.WithContext(ctx).First(&item, id).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			var count int64
			r.db.WithContext(ctx).Unscoped().Model(&domain.Item{}).Where("id = ? AND deleted_at IS NOT NULL", id).Count(&count)
			if count > 0 {
				return errors.New("item has already been deleted")
			}
			return errors.New("item not found")
		}
		return err
	}

	return r.db.WithContext(ctx).Delete(&item).Error
}
