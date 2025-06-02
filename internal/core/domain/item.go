package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// Item representa el modelo de dominio para un item
type Item struct {
	ID          uint           `json:"id" gorm:"primaryKey"`
	Name        string         `json:"name" validate:"required,min=3,max=100" gorm:"not null"`
	Description string         `json:"description" validate:"max=500"`
	Price       float64        `json:"price" validate:"required,gt=0" gorm:"not null"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at,omitempty" gorm:"index"`
}

// Validate realiza la validación del modelo
func (i *Item) Validate() error {
	validate := validator.New()
	return validate.Struct(i)
}

func (i *Item) SetCreateTime() {
	now := time.Now().UTC()
	i.CreatedAt = now
	i.UpdatedAt = now
}

func (i *Item) SetUpdateTime() {
	i.UpdatedAt = time.Now().UTC()
}
