package permission

import (
	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
)

type IPermissionsService interface {
	ListAll() ([]model.Permissions, error)
}

type permissionService struct{}

var PermissionIns IPermissionsService = permissionService{}

// ListAll 获取所有权限
func (this permissionService) ListAll() ([]model.Permissions, error) {
	var res []model.Permissions
	err := db.GetDb().Find(&res).Error
	return res, err
}
