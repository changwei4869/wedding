package model

import (
	"gorm.io/gorm"
	"time"
)

type Roles struct {
	Id          int            `gorm:"primarykey;comment:''" json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}

type RolesListReq struct {
	Id          int    `json:"id" form:"id"`                   //
	Name        string `json:"name" form:"name"`               //
	Description string `json:"description" form:"description"` //
	CreatedAt   string `json:"created_at" form:"created_at"`   //
	UpdatedAt   string `json:"updated_at" form:"updated_at"`   //
}

// RolesDetailReq roles详情参数
type RolesDetailReq struct {
	Id int `json:"id" form:"id"` //
}

// RolesAddReq roles新增参数
type RolesAddReq struct {
}

// RolesEditReq roles新增参数
type RolesEditReq struct {
	Id          int            `json:"id" form:"id"`                   //
	Name        string         `json:"name" form:"name"`               //
	Description string         `json:"description" form:"description"` //
	CreatedAt   time.Time      `json:"created_at" form:"created_at"`   //
	UpdatedAt   time.Time      `json:"updated_at" form:"updated_at"`   //
	DeletedAt   gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`   //
}

// RolesDelReq roles删除参数
type RolesDelReq struct {
	Id int `json:"id" form:"id"` //
}

// RolesDelBatchReq roles批量删除参数
type RolesDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` // 主键列表
}

// RolesResp roles返回信息
type RolesResp struct {
	Id          int            `json:"id" structs:"Id"`                   //
	Name        string         `json:"name" structs:"Name"`               //
	Description string         `json:"description" structs:"Description"` //
	CreatedAt   time.Time      `json:"created_at" structs:"CreatedAt"`    //
	UpdatedAt   time.Time      `json:"updated_at" structs:"UpdatedAt"`    //
	DeletedAt   gorm.DeletedAt `json:"deleted_at" structs:"DeletedAt"`    //
}
