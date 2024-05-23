package model

type Sites struct {
	ID     int    `gorm:"primarykey" json:"id"`
	City   string `json:"city"`
	Status int    `json:"status"`
}

type SitesAddReq struct {
	City   string `json:"city" form:"city"`
	Status int    `json:"status" form:"stago tus"`
}

type SitesEditReq struct {
	ID     int    `json:"id" form:"id"`
	City   string `json:"city" form:"city"`
	Status int    `json:"status" form:"status"`
}

type SitesResp struct {
	ID     int    `json:"id" form:"id"`
	City   string `json:"city" form:"city"`
	Status int    `json:"status" form:"status"`
}
