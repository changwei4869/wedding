package model

type Admins struct {
	Id       int    `gorm:"primarykey;comment:''" json:"id"` //
	Name     string `gorm:"comment:''" json:"name"`          //
	Phone    string `gorm:"comment:''" json:"phone"`         //
	PassWord string `gorm:"comment:''" json:"password"`
	Role_id  int    `gorm:"comment:''" json:"role_id"`
	Status   int    `gorm:"comment:''" json:"status"`
}

type AdminListReq struct {
	Id      int    `gorm:"primarykey;comment:''" json:"id"` //
	Name    string `gorm:"comment:''" json:"name"`          //
	Phone   string `gorm:"comment:''" json:"phone"`         //
	Role_id int    `gorm:"comment:''" json:"role_id"`
	Status  int    `gorm:"comment:''" json:"status"`
}

type AdminAddReq struct {
	Id       int    `gorm:"primarykey;comment:''" json:"id"` //
	Name     string `gorm:"comment:''" json:"name"`          //
	Phone    string `gorm:"comment:''" json:"phone"`         //
	PassWord string `gorm:"comment:''" json:"password"`
	Role_id  int    `gorm:"comment:''" json:"role_id"`
	Status   int    `gorm:"comment:''" json:"status"`
}

type AdminEditReq struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role_id  int    `gorm:"comment:''" json:"role_id"`
	Status   int    `gorm:"comment:''" json:"status"`
}

type AdminUpdatePwdReq struct {
	PassWord string `gorm:"comment:''" json:"password"`
}

type AdminResp struct {
	Id      int    `gorm:"primarykey;comment:''" json:"id"` //
	Name    string `gorm:"comment:''" json:"name"`          //
	Phone   string `gorm:"comment:''" json:"phone"`         //
	Role_id int    `gorm:"comment:''" json:"role_id"`
	Status  int    `gorm:"comment:''" json:"status"`
}
