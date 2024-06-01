package model

type Admins struct {
	Id       int    `gorm:"primarykey;comment:''" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role_id  int    `json:"role_id"`
	Status   int    `json:"status"`
}

type AdminListReq struct {
	Id      int    `gorm:"primarykey;comment:''" json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Role_id int    `json:"role_id"`
	Status  int    `json:"status"`
}

type AdminAddReq struct {
	Id       int    `gorm:"primarykey;comment:''" json:"id"`
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Password string `json:"password"`
	Role_id  int    `json:"role_id"`
	Status   int    `json:"status"`
}

type AdminEditReq struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Role_id  int    `json:"role_id"`
	Status   int    `json:"status"`
}

type AdminUpdatePwdReq struct {
	PassWord string `json:"password"`
}

type AdminResp struct {
	Id      int    `gorm:"primarykey;comment:''" json:"id"`
	Name    string `json:"name"`
	Phone   string `json:"phone"`
	Role_id int    `json:"role_id"`
	Status  int    `json:"status"`
}

type AdminLoginReq struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
