package model

import (
	"gorm.io/gorm"
	"time"
)

type Tags struct {
	Id        int            `gorm:"primarykey;comment:''" json:"id"` //
	Name      string         `gorm:"comment:''" json:"name"`          //
	Gender    string         `gorm:"comment:''" json:"gender"`        //
	CreatedAt time.Time      `gorm:"comment:''" json:"created_at"`    //
	UpdatedAt time.Time      `gorm:"comment:''" json:"updated_at"`    //
	DeletedAt gorm.DeletedAt `gorm:"comment:''" json:"deleted_at"`    //
}

// TagsListReq tags列表参数
type TagsListReq struct {
	Id        int    `json:"id" form:"id"`                 //
	Name      string `json:"name" form:"name"`             //
	Gender    string `json:"gender" form:"gender"`         //
	CreatedAt string `json:"created_at" form:"created_at"` //
	UpdatedAt string `json:"updated_at" form:"updated_at"` //
}

// TagsDetailReq tags详情参数
type TagsDetailReq struct {
	Id int `json:"id" form:"id"` //
}

// TagsAddReq tags新增参数
type TagsAddReq struct {
	Name   string `gorm:"comment:''" json:"name"`   //
	Gender string `gorm:"comment:''" json:"gender"` //
}

// TagsEditReq tags新增参数
type TagsEditReq struct {
	Id        int       `json:"id" form:"id"`                 //
	Name      string    `json:"name" form:"name"`             //
	Gender    string    `json:"gender" form:"gender"`         //
	CreatedAt time.Time `json:"created_at" form:"created_at"` //
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"` //
}

// TagsDelReq tags删除参数
type TagsDelReq struct {
	Id int `json:"id" form:"id"` //
}

// TagsDelBatchReq tags批量删除参数
type TagsDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` // 主键列表
}

// TagsResp tags返回信息
type TagsResp struct {
	Id        int            `json:"id" structs:"Id"`                //
	Name      string         `json:"name" structs:"Name"`            //
	Gender    string         `json:"gender" structs:"Gender"`        //
	CreatedAt time.Time      `json:"created_at" structs:"CreatedAt"` //
	UpdatedAt time.Time      `json:"updated_at" structs:"UpdatedAt"` //
	DeletedAt gorm.DeletedAt `json:"deleted_at" structs:"DeletedAt"` //
}
