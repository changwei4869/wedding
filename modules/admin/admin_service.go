package admin

import (
	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/utils/response"
	"gorm.io/gorm"
)

type IAdminsService interface {
	List() ([]model.Admins, error)
	Add(admin model.AdminAddReq) (e error)
	Del(id int) (e error)
	Edit(admin model.AdminEditReq) (e error)
}

type AdminsService struct {
	db *gorm.DB
}

func NewAdminsService(db *gorm.DB) *AdminsService {
	return &AdminsService{db: db}
}

// List admins列表
func (this AdminsService) List(page response.PageReq, listReq model.AdminListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	// 查询
	query := this.db.Model(&model.Admins{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name = ?", listReq.Name)
	}
	if listReq.Phone != "" {
		query = query.Where("phone = ?", listReq.Phone)
	}
	if listReq.Role_id != 0 {
		query = query.Where("role_id = ?", listReq.Role_id)
	}
	if listReq.Status != 0 {
		query = query.Where("status = ?", listReq.Status)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "admin AdminService List Count err"); e != nil {
		return
	}

	// 数据
	var rows []model.Admins
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "admin AdminService List Find err"); e != nil {
		return
	}

	// 创建响应切片并复制数据
	resps := make([]model.AdminResp, len(rows))
	for i, admin := range rows {
		resps[i] = model.AdminResp{
			Id:      admin.Id,
			Name:    admin.Name,
			Phone:   admin.Phone,
			Role_id: admin.Role_id,
			Status:  admin.Status,
		}
	}

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

func (this AdminsService) Add(admin model.AdminAddReq) error {
	return this.db.Create(&admin).Error
}

func (this AdminsService) Del(id int) error {
	return this.db.Delete(&model.Admins{}, id).Error
}

func (this *AdminsService) Edit(admin model.AdminEditReq) error {
	return this.db.Model(&model.Admins{}).Where("id = ?", admin.ID).Updates(admin).Error
}
