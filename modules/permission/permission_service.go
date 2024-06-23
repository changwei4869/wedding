package permission

import (
	"github.com/changwei4869/wedding/model"
	"gorm.io/gorm"
)

type ITagsService interface {
}

// NewPermissionService 创建新的 PermissionService
func NewPermissionService(db *gorm.DB) *PermissionService {
	return &PermissionService{db: db}
}

// ListAll 获取所有权限
func (service *PermissionService) ListAll() ([]model.Permissions, error) {
	var res []model.Permissions
	err := service.db.Find(&res).Error
	return res, err
}

// PermissionService 用于处理权限相关的数据库操作
type PermissionService struct {
	db *gorm.DB
}
