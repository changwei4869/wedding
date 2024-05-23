package model

import (
	"time"

	"gorm.io/gorm"
)

type Permissions struct {
	Id          int            `gorm:"primarykey;comment:''" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type PermissionsListReq struct {
	Id          int       `gorm:"primarykey;comment:''" json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
