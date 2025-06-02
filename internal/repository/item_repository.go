package repository

import (
	"api-rest-with-go/internal/models"
	"api-rest-with-go/pkg/database"
	"fmt"

	"gorm.io/gorm"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository() *ItemRepository {
	return &ItemRepository{
		db: database.DB,
	}
}

func (r *ItemRepository) Create(item *models.Item) error {
	return r.db.Create(item).Error
}

func (r *ItemRepository) GetAll() ([]models.Item, error) {
	var items []models.Item
	err := r.db.Find(&items).Error
	return items, err
}

func (r *ItemRepository) GetByID(id uint) (*models.Item, error) {
	var item models.Item
	err := r.db.First(&item, id).Error
	return &item, err
}

func (r *ItemRepository) Update(item *models.Item) error {
	// First check if item exists and is not soft deleted
	var existingItem models.Item
	result := r.db.Unscoped().First(&existingItem, item.ID)
	if result.Error != nil {
		return fmt.Errorf("Item with ID %d not found", item.ID)
	}

	// If item is soft deleted, return error
	if !existingItem.DeletedAt.Time.IsZero() {
		return fmt.Errorf("Item with ID %d has been deleted", item.ID)
	}

	// Update only if item exists and is not deleted
	result = r.db.Model(&existingItem).Updates(map[string]interface{}{
		"name":        item.Name,
		"description": item.Description,
		"price":       item.Price,
	})

	if result.Error != nil {
		return result.Error
	}

	*item = existingItem
	return nil
}

func (r *ItemRepository) Delete(id uint) error {
	result := r.db.Delete(&models.Item{}, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("item with ID %d not found", id)
	}
	return nil
}
