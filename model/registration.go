package model

import "time"

// Registrations registrations 结构体
type Registrations struct {
	Id                  int       `gorm:"primarykey;comment:''" json:"id"`
	Number              string    `json:"number"`
	Registrationscol    string    `json:"registrationscol"`
	Name                string    `json:"name"`
	Gender              int       `json:"gender"`
	Phone               string    `json:"phone"`
	Wechat              string    `json:"wechat"`
	Age                 int       `json:"age"`
	Height              int       `json:"height"`
	Weight              int       `json:"weight"`
	Location            string    `json:"location"`
	Residence           string    `json:"residence"`
	Education           string    `json:"education"`
	Qualifications      string    `json:"qualifications"`
	SexualOrientation   string    `json:"sexual_orientation"`
	MarriageHistory     string    `json:"marriage_history"`
	Profession          string    `json:"profession"`
	AnnualIncome        string    `json:"annual_income"`
	AssetStatus         string    `json:"asset_status"`
	SelfDescription     string    `json:"self_description"`
	MarriageCertificate string    `json:"marriage_certificate"`
	LiveTogether        string    `json:"live_together"`
	NeedChild           string    `json:"need_child"`
	BridePrice          string    `json:"bride_price"`
	Distance            string    `json:"distance"`
	WeddingMode         string    `json:"wedding_mode"`
	MarriedLife         string    `json:"married_life"`
	LookingFor          string    `json:"looking_for"`
	ExpectHelp          string    `json:"expect_help"`
	Channel             string    `json:"channel"`
	CreatedAt           time.Time `json:"created_at"`
	UpdatedAt           time.Time `json:"updated_at"`
	DeletedAt           time.Time `json:"deleted_at"`
}

// RegistrationsListReq registrations列表参数
type RegistrationsListReq struct {
	Id                  int       `json:"id" form:"id"`
	Number              string    `json:"number" form:"number"`
	Registrationscol    string    `json:"registrationscol" form:"registrationscol"`
	Name                string    `json:"name" form:"name"`
	Gender              int       `json:"gender" form:"gender"`
	Phone               string    `json:"phone" form:"phone"`
	Wechat              string    `json:"wechat" form:"wechat"`
	Age                 int       `json:"age" form:"age"`
	Height              int       `json:"height" form:"height"`
	Weight              int       `json:"weight" form:"weight"`
	Location            string    `json:"location" form:"location"`
	Residence           string    `json:"residence" form:"residence"`
	Education           string    `json:"education" form:"education"`
	Qualifications      string    `json:"qualifications" form:"qualifications"`
	SexualOrientation   string    `json:"sexual_orientation" form:"sexual_orientation"`
	MarriageHistory     string    `json:"marriage_history" form:"marriage_history"`
	Profession          string    `json:"profession" form:"profession"`
	AnnualIncome        string    `json:"annual_income" form:"annual_income"`
	AssetStatus         string    `json:"asset_status" form:"asset_status"`
	SelfDescription     string    `json:"self_description" form:"self_description"`
	MarriageCertificate string    `json:"marriage_certificate" form:"marriage_certificate"` //
	LiveTogether        string    `json:"live_together" form:"live_together"`
	NeedChild           string    `json:"need_child" form:"need_child"`
	BridePrice          string    `json:"bride_price" form:"bride_price"`
	Distance            string    `json:"distance" form:"distance"`
	WeddingMode         string    `json:"wedding_mode" form:"wedding_mode"`
	MarriedLife         string    `json:"married_life" form:"married_life"`
	LookingFor          string    `json:"looking_for" form:"looking_for"`
	ExpectHelp          string    `json:"expect_help" form:"expect_help"`
	Channel             string    `json:"channel" form:"channel"`
	CreatedAt           time.Time `json:"created_at" form:"created_at"` 
	UpdatedAt           time.Time `json:"updated_at" form:"updated_at"` 
	DeletedAt           time.Time `json:"deleted_at" form:"deleted_at"` 
}

// RegistrationsDetailReq registrations详情参数
type RegistrationsDetailReq struct {
	Id int `json:"id" form:"id"` //
}

// RegistrationsAddReq registrations新增参数
type RegistrationsAddReq struct {
}

// RegistrationsEditReq registrations新增参数
type RegistrationsEditReq struct {
	Id                  int    `json:"id" form:"id"`
	Number              string `json:"number" form:"number"`
	Registrationscol    string `json:"registrationscol" form:"registrationscol"`
	Name                string `json:"name" form:"name"`
	Gender              int    `json:"gender" form:"gender"`
	Phone               string `json:"phone" form:"phone"`
	Wechat              string `json:"wechat" form:"wechat"`
	Age                 int    `json:"age" form:"age"`
	Height              int    `json:"height" form:"height"`
	Weight              int    `json:"weight" form:"weight"`
	Location            string `json:"location" form:"location"`
	Residence           string `json:"residence" form:"residence"`
	Education           string `json:"education" form:"education"`
	Qualifications      string `json:"qualifications" form:"qualifications"`
	SexualOrientation   string `json:"sexual_orientation" form:"sexual_orientation"`
	MarriageHistory     string `json:"marriage_history" form:"marriage_history"`
	Profession          string `json:"profession" form:"profession"`
	AnnualIncome        string `json:"annual_income" form:"annual_income"`
	AssetStatus         string `json:"asset_status" form:"asset_status"`
	SelfDescription     string `json:"self_description" form:"self_description"`
	MarriageCertificate string `json:"marriage_certificate" form:"marriage_certificate"` //
	LiveTogether        string `json:"live_together" form:"live_together"`
	NeedChild           string `json:"need_child" form:"need_child"`
	BridePrice          string `json:"bride_price" form:"bride_price"`
	Distance            string `json:"distance" form:"distance"`
	WeddingMode         string `json:"wedding_mode" form:"wedding_mode"`
	MarriedLife         string `json:"married_life" form:"married_life"`
	LookingFor          string `json:"looking_for" form:"looking_for"`
	ExpectHelp          string `json:"expect_help" form:"expect_help"`
	Channel             string `json:"channel" form:"channel"`
	CreatedAt           time.Time `json:"created_at" form:"created_at"` 
	UpdatedAt           time.Time `json:"updated_at" form:"updated_at"` 
	DeletedAt           time.Time `json:"deleted_at" form:"deleted_at"` 
}

// RegistrationsDelReq registrations删除参数
type RegistrationsDelReq struct {
	Id int `json:"id" form:"id"` //
}

// RegistrationsDelBatchReq registrations批量删除参数
type RegistrationsDelBatchReq struct {
	Ids []int `json:"ids" form:"ids" binding:"required"` // 主键列表
}

// RegistrationsResp registrations返回信息
type RegistrationsResp struct {
	Id                  int    `json:"id" structs:"Id"`
	Number              string `json:"number" structs:"Number"`
	Registrationscol    string `json:"registrationscol" structs:"Registrationscol"`
	Name                string `json:"name" structs:"Name"`
	Gender              int    `json:"gender" structs:"Gender"`
	Phone               string `json:"phone" structs:"Phone"`
	Wechat              string `json:"wechat" structs:"Wechat"`
	Age                 int    `json:"age" structs:"Age"`
	Height              int    `json:"height" structs:"Height"`
	Weight              int    `json:"weight" structs:"Weight"`
	Location            string `json:"location" structs:"Location"`
	Residence           string `json:"residence" structs:"Residence"`
	Education           string `json:"education" structs:"Education"`
	Qualifications      string `json:"qualifications" structs:"Qualifications"`
	SexualOrientation   string `json:"sexual_orientation" structs:"SexualOrientation"`
	MarriageHistory     string `json:"marriage_history" structs:"MarriageHistory"`
	Profession          string `json:"profession" structs:"Profession"`
	AnnualIncome        string `json:"annual_income" structs:"AnnualIncome"`
	AssetStatus         string `json:"asset_status" structs:"AssetStatus"`
	SelfDescription     string `json:"self_description" structs:"SelfDescription"`
	MarriageCertificate string `json:"marriage_certificate" structs:"MarriageCertificate"` //
	LiveTogether        string `json:"live_together" structs:"LiveTogether"`
	NeedChild           string `json:"need_child" structs:"NeedChild"`
	BridePrice          string `json:"bride_price" structs:"BridePrice"`
	Distance            string `json:"distance" structs:"Distance"`
	WeddingMode         string `json:"wedding_mode" structs:"WeddingMode"`
	MarriedLife         string `json:"married_life" structs:"MarriedLife"`
	LookingFor          string `json:"looking_for" structs:"LookingFor"`
	ExpectHelp          string `json:"expect_help" structs:"ExpectHelp"`
	Channel             string `json:"channel" structs:"Channel"`
}
