package model

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	Id                  int            `gorm:"primarykey;comment:''" json:"id"`
	Uid                 string         `AssetStatusjson:"uid"`
	Avatar              string         `AssetStatusjson:"avatar"`
	Vitality            int            `AssetStatusjson:"vitality"`
	Diamond             int            `AssetStatusjson:"diamond"`
	Vip                 string         `AssetStatusjson:"vip"`
	Name                string         `AssetStatusjson:"name"`
	Password            string         `AssetStatusjson:"password"`
	RoleId              int            `AssetStatusjson:"role_id"`
	Phone               string         `AssetStatusjson:"phone"`
	Wechat              string         `AssetStatusjson:"wechat"`
	Gender              int            `AssetStatusjson:"gender"`
	Blacklist           int            `AssetStatusjson:"blacklist"`
	Age                 int            `AssetStatusjson:"age"`
	Height              int            `AssetStatusjson:"height"`
	Weight              int            `AssetStatusjson:"weight"`
	Location            string         `AssetStatusjson:"location"`
	Residence           string         `AssetStatusjson:"residence"`
	Qualifications      string         `AssetStatusjson:"qualifications"`
	SexualOrientation   string         `AssetStatusjson:"sexual_orientation"`
	MarriageHistory     string         `AssetStatusjson:"marriage_history"`
	Profession          string         `AssetStatusjson:"profession"`
	AnnualIncome        string         `AssetStatusjson:"annual_income"`
	AssetStatus         string         `AssetStatusjson:"asset_status"`
	SelfDescription     string         `AssetStatusjson:"self_description"`
	MarriageCertificate string         `AssetStatusjson:"marriage_certificate"`
	LiveTogether        string         `AssetStatusjson:"live_together"`
	NeedChild           string         `AssetStatusjson:"need_child"`
	BridePrice          string         `AssetStatusjson:"bride_price"`
	Distance            string         `AssetStatusjson:"distance"`
	WeddingMode         string         `AssetStatusjson:"wedding_mode"`
	MarriedLife         string         `AssetStatusjson:"married_life"`
	LookingFor          string         `AssetStatusjson:"looking_for"`
	ExpectHelp          string         `AssetStatusjson:"expect_help"`
	Channel             string         `AssetStatusjson:"channel"`
	LogOut              string         `AssetStatusjson:"log_out"`
	CreatedBy           string         `AssetStatusjson:"created_by"`
	CreatedAt           time.Time      `AssetStatusjson:"created_at"`
	UpdatedAt           time.Time      `AssetStatusjson:"updated_at"`
	DeletedAt           gorm.DeletedAt `AssetStatusjson:"deleted_at"`
}

// UsersListReq users列表参数
type UsersListReq struct {
	Id                  int            `json:"id" form:"id"`
	Uid                 string         `json:"uid" form:"uid"`
	Avatar              string         `json:"avatar" form:"avatar"`
	Vitality            int            `json:"vitality" form:"vitality"`
	Diamond             int            `json:"diamond" form:"diamond"`
	Vip                 string         `json:"vip" form:"vip"`
	Name                string         `json:"name" form:"name"`
	Password            string         `json:"password" form:"password"`
	RoleId              int            `json:"role_id" form:"role_id"`
	Phone               string         `json:"phone" form:"phone"`
	Wechat              string         `json:"wechat" form:"wechat"`
	Gender              int            `json:"gender" form:"gender"`
	Blacklist           int            `json:"blacklist" form:"blacklist"`
	Age                 int            `json:"age" form:"age"`
	Height              int            `json:"height" form:"height"`
	Weight              int            `json:"weight" form:"weight"`
	Location            string         `json:"location" form:"location"`
	Residence           string         `json:"residence" form:"residence"`
	Qualifications      string         `json:"qualifications" form:"qualifications"`
	SexualOrientation   string         `json:"sexual_orientation" form:"sexual_orientation"`
	MarriageHistory     string         `json:"marriage_history" form:"marriage_history"`
	Profession          string         `json:"profession" form:"profession"`
	AnnualIncome        string         `json:"annual_income" form:"annual_income"`
	AssetStatus         string         `json:"asset_status" form:"asset_status"`
	SelfDescription     string         `json:"self_description" form:"self_description"`
	MarriageCertificate string         `json:"marriage_certificate" form:"marriage_certificate"`
	LiveTogether        string         `json:"live_together" form:"live_together"`
	NeedChild           string         `json:"need_child" form:"need_child"`
	BridePrice          string         `json:"bride_price" form:"bride_price"`
	Distance            string         `json:"distance" form:"distance"`
	WeddingMode         string         `json:"wedding_mode" form:"wedding_mode"`
	MarriedLife         string         `json:"married_life" form:"married_life"`
	LookingFor          string         `json:"looking_for" form:"looking_for"`
	ExpectHelp          string         `json:"expect_help" form:"expect_help"`
	Channel             string         `json:"channel" form:"channel"`
	LogOut              string         `json:"log_out" form:"log_out"`
	CreatedBy           string         `json:"created_by" form:"created_by"`
	CreatedAt           time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}

// UsersDetailReq users详情参数
type UsersDetailReq struct {
	Id int `json:"id" form:"id"`
}

// UsersAddReq users新增参数
type UsersAddReq struct {
}

// UsersEditReq users新增参数
type UsersEditReq struct {
	Id                  int            `json:"id" form:"id"`
	Uid                 string         `json:"uid" form:"uid"`
	Avatar              string         `json:"avatar" form:"avatar"`
	Vitality            int            `json:"vitality" form:"vitality"`
	Diamond             int            `json:"diamond" form:"diamond"`
	Vip                 string         `json:"vip" form:"vip"`
	Name                string         `json:"name" form:"name"`
	Password            string         `json:"password" form:"password"`
	RoleId              int            `json:"role_id" form:"role_id"`
	Phone               string         `json:"phone" form:"phone"`
	Wechat              string         `json:"wechat" form:"wechat"`
	Gender              int            `json:"gender" form:"gender"`
	Blacklist           int            `json:"blacklist" form:"blacklist"`
	Age                 int            `json:"age" form:"age"`
	Height              int            `json:"height" form:"height"`
	Weight              int            `json:"weight" form:"weight"`
	Location            string         `json:"location" form:"location"`
	Residence           string         `json:"residence" form:"residence"`
	Qualifications      string         `json:"qualifications" form:"qualifications"`
	SexualOrientation   string         `json:"sexual_orientation" form:"sexual_orientation"`
	MarriageHistory     string         `json:"marriage_history" form:"marriage_history"`
	Profession          string         `json:"profession" form:"profession"`
	AnnualIncome        string         `json:"annual_income" form:"annual_income"`
	AssetStatus         string         `json:"asset_status" form:"asset_status"`
	SelfDescription     string         `json:"self_description" form:"self_description"`
	MarriageCertificate string         `json:"marriage_certificate" form:"marriage_certificate"`
	LiveTogether        string         `json:"live_together" form:"live_together"`
	NeedChild           string         `json:"need_child" form:"need_child"`
	BridePrice          string         `json:"bride_price" form:"bride_price"`
	Distance            string         `json:"distance" form:"distance"`
	WeddingMode         string         `json:"wedding_mode" form:"wedding_mode"`
	MarriedLife         string         `json:"married_life" form:"married_life"`
	LookingFor          string         `json:"looking_for" form:"looking_for"`
	ExpectHelp          string         `json:"expect_help" form:"expect_help"`
	Channel             string         `json:"channel" form:"channel"`
	LogOut              string         `json:"log_out" form:"log_out"`
	CreatedBy           string         `json:"created_by" form:"created_by"`
	CreatedAt           time.Time      `json:"created_at" form:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at" form:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at" form:"deleted_at"`
}

// UsersDelBatchReq users批量删除参数
type UsersDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` // 主键列表
}

// UsersResp users返回信息
type UsersResp struct {
	Id                  int       `json:"id" structs:"Id"`
	Uid                 string    `json:"uid" structs:"Uid"`
	Avatar              string    `json:"avatar" structs:"Avatar"`
	Vitality            int       `json:"vitality" structs:"Vitality"`
	Diamond             int       `json:"diamond" structs:"Diamond"`
	Vip                 string    `json:"vip" structs:"Vip"`
	Name                string    `json:"name" structs:"Name"`
	Password            string    `json:"password" structs:"Password"`
	RoleId              int       `json:"role_id" structs:"RoleId"`
	Phone               string    `json:"phone" structs:"Phone"`
	Wechat              string    `json:"wechat" structs:"Wechat"`
	Gender              int       `json:"gender" structs:"Gender"`
	Blacklist           int       `json:"blacklist" structs:"Blacklist"`
	Age                 int       `json:"age" structs:"Age"`
	Height              int       `json:"height" structs:"Height"`
	Weight              int       `json:"weight" structs:"Weight"`
	Location            string    `json:"location" structs:"Location"`
	Residence           string    `json:"residence" structs:"Residence"`
	Qualifications      string    `json:"qualifications" structs:"Qualifications"`
	SexualOrientation   string    `json:"sexual_orientation" structs:"SexualOrientation"`
	MarriageHistory     string    `json:"marriage_history" structs:"MarriageHistory"`
	Profession          string    `json:"profession" structs:"Profession"`
	AnnualIncome        string    `json:"annual_income" structs:"AnnualIncome"`
	AssetStatus         string    `json:"asset_status" structs:"AssetStatus"`
	SelfDescription     string    `json:"self_description" structs:"SelfDescription"`
	MarriageCertificate string    `json:"marriage_certificate" structs:"MarriageCertificate"`
	LiveTogether        string    `json:"live_together" structs:"LiveTogether"`
	NeedChild           string    `json:"need_child" structs:"NeedChild"`
	BridePrice          string    `json:"bride_price" structs:"BridePrice"`
	Distance            string    `json:"distance" structs:"Distance"`
	WeddingMode         string    `json:"wedding_mode" structs:"WeddingMode"`
	MarriedLife         string    `json:"married_life" structs:"MarriedLife"`
	LookingFor          string    `json:"looking_for" structs:"LookingFor"`
	ExpectHelp          string    `json:"expect_help" structs:"ExpectHelp"`
	Channel             string    `json:"channel" structs:"Channel"`
	LogOut              string    `json:"log_out" structs:"LogOut"`
	CreatedBy           string    `json:"created_by" structs:"CreatedBy"`
	CreatedAt           time.Time `json:"created_at" structs:"CreatedAt"`
	UpdatedAt           time.Time `json:"updated_at" structs:"UpdatedAt"`
	DeletedAt           time.Time `json:"deleted_at" structs:"DeletedAt"`
}
