package category

import (
	"errors"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
	"github.com/changwei4869/wedding/utils/response"
)

type ICategoriesService interface {
	All() (res response.PageResp, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.CategoriesListReq) (res response.PageResp, e error)
	Detail(id int) (res model.CategoriesResp, e error)
	Add(addReq model.CategoriesAddReq) (e error)
	Edit(editReq model.CategoriesEditReq) (e error)
	Change(changeReq model.CategoriesDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.CategoriesDelBatchReq) (e error)
}

// categoriesService categories
type categoriesService struct{}

var CategoriesIns ICategoriesService = categoriesService{}

// All categories列表
func (this categoriesService) All() (res response.PageResp, e error) {
	// 数据
	query := db.GetDb().Model(&model.Categories{})
	var rows []model.Categories
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "categories categoriesService  All Find err"); e != nil {
		return
	}
	resps := []model.CategoriesResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   0,
		PageSize: 0,
		Count:    0,
		Data:     resps,
	}, nil
}

// Count categories
func (this categoriesService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := db.GetDb().Model(&model.Categories{})
	var rows []model.Categories
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "categories categoriesService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List categories列表
func (this categoriesService) List(page response.PageReq, listReq model.CategoriesListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	query := db.GetDb().Model(&model.Categories{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name = ?", listReq.Name)
	}
	if listReq.ParentId > 0 {
		query = query.Where("parent_id = ?", listReq.ParentId)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "categories categoriesService  List Count err"); e != nil {
		return
	}
	// 数据
	var rows []model.Categories
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "categories categoriesService  List Find err"); e != nil {
		return
	}
	resps := []model.CategoriesResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail categories详情
func (this categoriesService) Detail(id int) (res model.CategoriesResp, e error) {
	var row model.Categories
	err := db.GetDb().Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "categories categories  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "categories categoriesService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	return
}

// Add categories新增
func (this categoriesService) Add(addReq model.CategoriesAddReq) (e error) {
	var row model.Categories
	response.CopyStruct(&row, addReq)
	err := db.GetDb().Create(&row).Error
	e = response.CheckErr(err, "categories categoriesService  Add Create err")
	return
}

// Edit categories编辑
func (this categoriesService) Edit(editReq model.CategoriesEditReq) (e error) {
	var row model.Categories
	err := db.GetDb().Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "categories categoriesService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "categories categoriesService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = db.GetDb().Model(&row).Updates(row).Error
	e = response.CheckErr(err, "categories categoriesService Edit Updates err")

	//强制更新 当IsShow=0
	//err = db.GetDb().Model(&row).Select("IsShow").Updates(row).Error
	//e = response.CheckErr(err, "categories categoriesService  Edit Updates err")
	//强制更新 isDisable=0
	//err = db.GetDb().Model(&row).Updates(map[string]interface{}{"IsDisable": editReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	//e = response.CheckErr(err, "categories categoriesService  Edit Updates err")
	return
}

// Change categories 状态切换
func (this categoriesService) Change(changeReq model.CategoriesDetailReq) (e error) {
	var row model.Categories
	err := db.GetDb().Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "categories categoriesService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "categories categoriesService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = db.GetDb().Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = db.GetDb().Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "categories categoriesService  Change Updates err")
	return
}

// Del categories删除
func (this categoriesService) Del(id int) (e error) {
	var row model.Categories
	err := db.GetDb().Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "categories categoriesService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "categories categoriesService Del First err"); e != nil {
		return
	}
	// 删除
	err = db.GetDb().Delete(&row).Error
	e = response.CheckErr(err, "categories categoriesService Del Delete err")
	return
}

// DelBatch categories 批量删除
func (this categoriesService) DelBatch(delReq model.CategoriesDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := db.GetDb().Where("id IN (?)", delReq.Ids).Delete(model.Categories{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 CategoriesService DelBatch Delete err")
	return
}
