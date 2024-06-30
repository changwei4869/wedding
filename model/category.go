package model

import (
	"time"

	"gorm.io/gorm"
)

type Categories struct {
	Id        int            `gorm:"primarykey;comment:''" json:"id"`
	Name      string         `gorm:"comment:''" json:"name"`
	ParentId  int            `gorm:"comment:''" json:"parent_id"`
	CreatedAt time.Time      `gorm:"comment:''" json:"created_at"`
	UpdatedAt time.Time      `gorm:"comment:''" json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

// CategoriesListReq categories列表参数
type CategoriesListReq struct {
	Id        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	ParentId  int       `json:"parent_id" form:"parent_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

// CategoriesDetailReq categories详情参数
type CategoriesDetailReq struct {
	Id int `json:"id" form:"id"`
}

// CategoriesAddReq categories新增参数
type CategoriesAddReq struct {
	Name     string `json:"name" form:"name"`
	ParentId int    `json:"parent_id" form:"parent_id"`
}

// CategoriesEditReq categories新增参数
type CategoriesEditReq struct {
	Id        int       `json:"id" form:"id"`
	Name      string    `json:"name" form:"name"`
	ParentId  int       `json:"parent_id" form:"parent_id"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

// CategoriesDelReq categories删除参数
type CategoriesDelReq struct {
	Id int `json:"id" form:"id"`
}

// CategoriesDelBatchReq categories批量删除参数
type CategoriesDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// CategoriesResp categories返回信息
type CategoriesResp struct {
	Id        int       `json:"id" structs:"Id"`
	Name      string    `json:"name" structs:"Name"`
	ParentId  int       `json:"parent_id" structs:"ParentId"`
	CreatedAt time.Time `json:"created_at" structs:"CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" structs:"UpdatedAt"`
	DeletedAt time.Time `json:"deleted_at" structs:"DeletedAt"`
}
