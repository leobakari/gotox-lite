package models

import "gorm.io/gorm"

type Todo struct {
	ID          uint   `json:"id"`
	Title       string `json:"name"`
	Description string `json:"description"`
	IsDone      bool   `json:"is_done"`
	UserID      uint   `json:"user_id"`
	gorm.Model         // Build in functionality for created-, updated- and deleted-at
}
