package admin

import (
	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
	"github.com/changwei4869/wedding/utils/response"
)

type IAdminsService interface {
	All() (res []model.Admins, e error)
	Add(admin model.AdminAddReq) (e error)
	Del(id int) (e error)
	Edit(admin model.AdminEditReq) (e error)
	GetAdminByPhone(phone string) (model.Admins, error)
}

type AdminsService struct {
}

var AdminIns IAdminsService = AdminsService{}

// List admins列表
func (this AdminsService) All() (res []model.Admins, e error) {
	// 数据
	query := db.GetDb().Model(&model.Admins{})
	var rows []model.Admins
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "qrCodes qrCodesService  All Find err"); e != nil {
		return
	}
	return rows, nil
}

func (this AdminsService) Add(admin model.AdminAddReq) (e error) {
	var row model.Admins
	response.CopyStruct(&row, admin)
	err := db.GetDb().Create(&row).Error
	e = response.CheckErr(err, "admin AdminsService Create err")
	return
}

func (this AdminsService) Del(id int) error {
	return db.GetDb().Delete(&model.Admins{}, id).Error
}

func (this AdminsService) Edit(admin model.AdminEditReq) error {
	return db.GetDb().Model(&model.Admins{}).Where("id = ?", admin.ID).Updates(admin).Error
}

func (this AdminsService) GetAdminByPhone(phone string) (model.Admins, error) {
	var admin model.Admins
	err := db.GetDb().Where("phone = ?", phone).First(&admin).Error
	response.CheckErr(err, "admin AdminsService GetAdminByPhone err")
	return admin, err
}
