package model

import (
	"time"

	"gorm.io/gorm"
)

type Roles struct {
	Id            int            `gorm:"primarykey;comment:''" json:"id"`
	Name          string         `json:"name"`
	Description   string         `json:"description"`
	PermissionIds string         `json:"permission_ids"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}
type RolesListReq struct {
	Id            int       `json:"id" form:"id"`
	Name          string    `json:"name" form:"name"`
	Description   string    `json:"description" form:"description"`
	PermissionIds string    `json:"permission_ids" form:"permission_ids"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" form:"deleted_at"`
}

type RolesDetailReq struct {
	Id int `json:"id" form:"id"`
}

type RolesAddReq struct {
	Name          string `json:"name" form:"name"`
	Description   string `json:"description" form:"description"`
	PermissionIds []int  `json:"permission_ids" form:"permission_ids"`
}

type RolesEditReq struct {
	Id            int       `json:"id" form:"id"`
	Name          string    `json:"name" form:"name"`
	Description   string    `json:"description" form:"description"`
	PermissionIds string    `json:"permission_ids" form:"permission_ids"`
	CreatedAt     time.Time `json:"created_at" form:"created_at"`
	UpdatedAt     time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt     time.Time `json:"deleted_at" form:"deleted_at"`
}

type RolesDelReq struct {
	Id int `json:"id" form:"id"`
}

type RolesDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

type RolesResp struct {
	Id            int       `json:"id" structs:"Id"`
	Name          string    `json:"name" structs:"Name"`
	Description   string    `json:"description" structs:"Description"`
	PermissionIds string    `json:"permission_ids" structs:"PermissionIds"`
	CreatedAt     time.Time `json:"created_at" structs:"CreatedAt"`
	UpdatedAt     time.Time `json:"updated_at" structs:"UpdatedAt"`
	DeletedAt     time.Time `json:"deleted_at" structs:"DeletedAt"`
}
