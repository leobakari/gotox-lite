package models

import "gorm.io/gorm"

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsActive    bool   `json:"is_active"`
	UserID      uint   `json:"user_id"`
	gorm.Model         // Build in functionality for created-, updated- and deleted-at
}
