package models

import "gorm.io/gorm"

type User struct {
	ID         uint   `json:"id"`
	UserName   string `json:"username"`
	Age        int    `json:"age"`
	gorm.Model        // Build in functionality for created-, updated- and deleted-at
}
