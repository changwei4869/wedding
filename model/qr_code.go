package model

import (
	"time"

	"gorm.io/gorm"
)

type QrCodes struct {
	Id        int            `gorm:"primarykey;comment:''" json:"id"`
	QrCode    string         `gorm:"comment:''" json:"qr_code"`
	Location  string         `gorm:"comment:''" json:"location"`
	CreatedAt time.Time      `gorm:"comment:''" json:"created_at"`
	UpdatedAt time.Time      `gorm:"comment:''" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"comment:''" json:"deleted_at" swaggerignore:"true"`
}

// QrCodesListReq qrCodes列表参数
type QrCodesListReq struct {
	Id        int       `json:"id" form:"id"`
	QrCode    string    `json:"qr_code" form:"qr_code"`
	Location  string    `json:"location" form:"location"`
	CreatedAt time.Time `json:"created_at" form:"created_at"`
	UpdatedAt time.Time `json:"updated_at" form:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" form:"deleted_at"`
}

// QrCodesDetailReq qrCodes详情参数
type QrCodesDetailReq struct {
	Id int `json:"id" form:"id"`
}

// QrCodesAddReq qrCodes新增参数
type QrCodesAddReq struct {
	QrCode   string `json:"qr_code" form:"qr_code"`
	Location string `json:"location" form:"location"`
}

// QrCodesEditReq qrCodes新增参数
type QrCodesEditReq struct {
	Id       int    `json:"id" form:"id"`
	QrCode   string `json:"qr_code" form:"qr_code"`
	Location string `json:"location" form:"location"`
}

// QrCodesDelReq qrCodes删除参数
type QrCodesDelReq struct {
	Id int `json:"id" form:"id"`
}

// QrCodesDelBatchReq qrCodes批量删除参数
type QrCodesDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"`
}

// QrCodesResp qrCodes返回信息
type QrCodesResp struct {
	Id        int       `json:"id" structs:"Id"`
	QrCode    string    `json:"qr_code" structs:"QrCode"`
	Location  string    `json:"location" structs:"Location"`
	CreatedAt time.Time `json:"created_at" structs:"CreatedAt"`
	UpdatedAt time.Time `json:"updated_at" structs:"UpdatedAt"`
	DeletedAt time.Time `json:"deleted_at" structs:"DeletedAt"`
}
