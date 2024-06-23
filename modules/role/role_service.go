package role

import (
	"errors"
	"strconv"
	"strings"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/utils/response"
	"gorm.io/gorm"
)

type IRolesService interface {
	All() (res []model.Roles, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.RolesListReq) (res response.PageResp, e error)
	Detail(id int) (res model.RolesResp, e error)
	Add(addReq model.RolesAddReq) (e error)
	Edit(editReq model.RolesEditReq) (e error)
	Change(changeReq model.RolesDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.RolesDelBatchReq) (e error)
}

// NewRolesService 初始化
func NewRolesService(db *gorm.DB) IRolesService {
	return &rolesService{db: db}
}

// rolesService roles
type rolesService struct {
	db *gorm.DB
}

// All roles列表
func (this rolesService) All() (res []model.Roles, e error) {
	// 数据
	query := this.db.Model(&model.Roles{})
	var rows []model.Roles
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "roles rolesService  All Find err"); e != nil {
		return
	}
	return rows, nil
}

// Count roles
func (this rolesService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := this.db.Model(&model.Roles{})
	var rows []model.Roles
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "roles rolesService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List roles列表
func (this rolesService) List(page response.PageReq, listReq model.RolesListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)

	// 查询
	query := this.db.Model(&model.Roles{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name = ?", listReq.Name)
	}
	if listReq.Description != "" {
		query = query.Where("description = ?", listReq.Description)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "roles rolesService List Count err"); e != nil {
		return
	}

	// 数据
	var rows []model.Roles
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "roles rolesService List Find err"); e != nil {
		return
	}

	resps := make([]model.RolesResp, len(rows))
	for i, role := range rows {
		resps[i] = model.RolesResp{
			Id:          role.Id,
			Name:        role.Name,
			Description: role.Description,
			CreatedAt:   role.CreatedAt,
			UpdatedAt:   role.UpdatedAt,
		}
	}

	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail roles详情
func (this rolesService) Detail(id int) (res model.RolesResp, e error) {
	var row model.Roles
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "roles roles  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "roles rolesService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	return
}

// Add roles新增
func (this rolesService) Add(addReq model.RolesAddReq) (e error) {
	var row model.Roles
	var permissionIds []string
	row.Name = addReq.Name
	row.Description = addReq.Description
	for _, id := range addReq.PermissionIds {
		permissionIds = append(permissionIds, strconv.Itoa(id))
	}
	row.PermissionIds = strings.Join(permissionIds, ",")
	err := this.db.Create(&row).Error
	e = response.CheckErr(err, "roles rolesService  Add Create err")
	return
}

// Edit roles编辑
func (this rolesService) Edit(editReq model.RolesEditReq) (e error) {
	var row model.Roles
	err := this.db.Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "roles rolesService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "roles rolesService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = this.db.Model(&row).Updates(row).Error
	e = response.CheckErr(err, "roles rolesService Edit Updates err")

	//强制更新 当IsShow=0
	//err = this.db.Model(&row).Select("IsShow").Updates(row).Error
	//e = response.CheckErr(err, "roles rolesService  Edit Updates err")
	//强制更新 isDisable=0
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": editReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	//e = response.CheckErr(err, "roles rolesService  Edit Updates err")
	return
}

// Change roles 状态切换
func (this rolesService) Change(changeReq model.RolesDetailReq) (e error) {
	var row model.Roles
	err := this.db.Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "roles rolesService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "roles rolesService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = this.db.Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "roles rolesService  Change Updates err")
	return
}

// Del roles删除
func (this rolesService) Del(id int) (e error) {
	var row model.Roles
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "roles rolesService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "roles rolesService Del First err"); e != nil {
		return
	}
	// 删除
	err = this.db.Delete(&row).Error
	e = response.CheckErr(err, "roles rolesService Del Delete err")
	return
}

// DelBatch roles 批量删除
func (this rolesService) DelBatch(delReq model.RolesDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := this.db.Where("id IN (?)", delReq.Ids).Delete(model.Roles{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 RolesService DelBatch Delete err")
	return
}
