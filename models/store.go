// models/task.go

package models

import (
	"time"
)

type Product struct {
	ID          uint      `json:"id" gorm:"primary_key"`
	CodeProduct string    `json:"codeProduct"`
	Name        string    `json:"name"`
	Price       string    `json:"price"`
	Stock       string    `json:"stock"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

type Category struct {
	ID           uint      `json:"id" gorm:"primary_key"`
	CodeCategory string    `json:"codeCategory"`
	Name         string    `json:"name"`
	CreatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt    time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}
