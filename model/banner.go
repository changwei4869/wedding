package model

import (
	"time"

	"gorm.io/gorm"
)

type Banners struct {
	Id           int            `gorm:"primarykey;comment:''" json:"id"`
	ImageUrl     string         `gorm:"comment:''" json:"image_url"`
	LinkArticle  string         `gorm:"comment:''" json:"link_article"`
	LinkUrl      string         `gorm:"comment:''" json:"link_url"`
	DisplayOrder int            `gorm:"comment:''" json:"display_order"`
	CreatedAt    time.Time      `gorm:"comment:''" json:"created_at"`
	UpdatedAt    time.Time      `gorm:"comment:''" json:"updated_at"`
	DeletedAt    gorm.DeletedAt `json:"deleted_at" swaggerignore:"true"`
}

// BannersListReq banners列表参数
type BannersListReq struct {
	Id           int       `json:"id" form:"id"`
	ImageUrl     string    `json:"image_url" form:"image_url"`
	LinkArticle  string    `json:"link_article" form:"link_article"`
	LinkUrl      string    `json:"link_url" form:"link_url"`
	DisplayOrder int       `json:"display_order" form:"display_order"`
	CreatedAt    time.Time `json:"created_at" form:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt    time.Time `json:"deleted_at" form:"deleted_at"`
}

// BannersDetailReq banners详情参数
type BannersDetailReq struct {
	Id int `json:"id" form:"id"`
}

// BannersAddReq banners新增参数
type BannersAddReq struct {
	ImageUrl     string `json:"image_url" form:"image_url"`
	LinkArticle  string `json:"link_article" form:"link_article"`
	LinkUrl      string `json:"link_url" form:"link_url"`
	DisplayOrder int    `json:"display_order" form:"display_order"`
}

// BannersEditReq banners新增参数
type BannersEditReq struct {
	Id           int       `json:"id" form:"id"`
	ImageUrl     string    `json:"image_url" form:"image_url"`
	LinkArticle  string    `json:"link_article" form:"link_article"`
	LinkUrl      string    `json:"link_url" form:"link_url"`
	DisplayOrder int       `json:"display_order" form:"display_order"`
}

// BannersDelReq banners删除参数
type BannersDelReq struct {
	Id int `json:"id" form:"id"`
}

// BannersDelBatchReq banners批量删除参数
type BannersDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` // 主键列表
}

// BannersResp banners返回信息
type BannersResp struct {
	Id           int       `json:"id" structs:"Id"`
	ImageUrl     string    `json:"image_url" structs:"ImageUrl"`
	LinkArticle  string    `json:"link_article" structs:"LinkArticle"`
	LinkUrl      string    `json:"link_url" structs:"LinkUrl"`
	DisplayOrder int       `json:"display_order" structs:"DisplayOrder"`
	CreatedAt    time.Time `json:"created_at" structs:"CreatedAt"`
	UpdatedAt    time.Time `json:"updated_at" structs:"UpdatedAt"`
	DeletedAt    time.Time `json:"deleted_at" structs:"DeletedAt"`
}
