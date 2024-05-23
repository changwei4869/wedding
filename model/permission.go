package model

import (
	"time"

	"gorm.io/gorm"
)

type Permissions struct {
	Id          int            `gorm:"primarykey;comment:''" json:"id"` //
	Name        string         `gorm:"comment:''" json:"name"`          //
	Description string         `gorm:"comment:''" json:"description"`   //
	CreatedAt   time.Time      `gorm:"comment:''" json:"created_at"`    //
	UpdatedAt   time.Time      `gorm:"comment:''" json:"updated_at"`    //
	DeletedAt   gorm.DeletedAt `gorm:"comment:''" json:"deleted_at"`    //
}

type PermissionsListReq struct {
	Id          int       `gorm:"primarykey;comment:''" json:"id"` //
	Name        string    `gorm:"comment:''" json:"name"`          //
	Description string    `gorm:"comment:''" json:"description"`   //
	CreatedAt   time.Time `gorm:"comment:''" json:"created_at"`    //
	UpdatedAt   time.Time `gorm:"comment:''" json:"updated_at"`    //
}
