package banner

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
	"github.com/changwei4869/wedding/utils/response"
)

type IBannersService interface {
	All() (res response.PageResp, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.BannersListReq) (res response.PageResp, e error)
	Detail(id int) (res model.BannersResp, e error)
	Add(addReq model.BannersAddReq) (e error)
	Edit(editReq model.BannersEditReq) (e error)
	Change(changeReq model.BannersDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.BannersDelBatchReq) (e error)
	GetMaxDisplayOrder() (res int, e error)
	GetAll() (res []model.Banners, e error)
	Up(id int) (e error)
	Down(id int) (e error)
}

// bannersService banners
type bannersService struct{}

var BannerIns IBannersService = bannersService{}

// 上移
func (this bannersService) Up(id int) (e error) {
	var row model.Banners
	err := db.GetDb().Where("id = ?", id).First(&row).Error
	fmt.Println(err)
	if e = response.ErrRecordNotFound(err, "banners bannersService Up ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService  Up First err"); e != nil {
		return
	}
	frontBanner := model.Banners{}
	err = db.GetDb().Order("display_order desc").Where("display_order < ?", row.DisplayOrder).First(&frontBanner).Error
	if e = response.ErrRecordNotFound(err, "banners bannersService Up ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService Order Up First err"); e != nil {
		return
	}

	// 交换
	row.DisplayOrder, frontBanner.DisplayOrder = frontBanner.DisplayOrder, row.DisplayOrder
	err = db.GetDb().Model(&row).Update("display_order", row.DisplayOrder).Error
	if e = response.CheckErr(err, "banners bannersService Up Update err"); e != nil {
		return
	}
	err = db.GetDb().Model(&frontBanner).Update("display_order", frontBanner.DisplayOrder).Error
	if e = response.CheckErr(err, "banners bannersService Up Update err"); e != nil {
		return
	}

	return
}

// 下移
func (this bannersService) Down(id int) (e error) {
	var row model.Banners
	err := db.GetDb().Where("id = ?", id).First(&row).Error
	if e = response.ErrRecordNotFound(err, "banners bannersService Down ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService  Down First err"); e != nil {
		return
	}
	nextBanner := model.Banners{}
	err = db.GetDb().Order("display_order asc").Where("display_order > ?", row.DisplayOrder).First(&nextBanner).Error
	if e = response.ErrRecordNotFound(err, "banners bannersService Down ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService Order Down First err"); e != nil {
		return
	}

	// 交换
	row.DisplayOrder, nextBanner.DisplayOrder = nextBanner.DisplayOrder, row.DisplayOrder
	err = db.GetDb().Model(&row).Update("display_order", row.DisplayOrder).Error
	if e = response.CheckErr(err, "banners bannersService Down Update err"); e != nil {
		return
	}
	err = db.GetDb().Model(&nextBanner).Update("display_order", nextBanner.DisplayOrder).Error
	if e = response.CheckErr(err, "banners bannersService Down Update err"); e != nil {
		return
	}
	return
}

// 按display_order升序排列
func (this bannersService) GetAll() (res []model.Banners, e error) {
	err := db.GetDb().Order("display_order asc").Find(&res).Error
	e = response.CheckErr(err, "banners bannersService  GetAll Find err")
	return
}

// 获取表中最大display_order
func (this bannersService) GetMaxDisplayOrder() (res int, e error) {
	var row model.Banners
	err := db.GetDb().Order("display_order desc").Limit(1).Find(&row).Error
	if e = response.CheckErr(err, "banners bannersService  GetMaxDisplayOrder First err"); e != nil {
		return
	}
	return row.DisplayOrder, nil
}

// All banners列表
func (this bannersService) All() (res response.PageResp, e error) {
	// 数据
	query := db.GetDb().Model(&model.Banners{})
	var rows []model.Banners
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "banners bannersService  All Find err"); e != nil {
		return
	}
	resps := []model.BannersResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   0,
		PageSize: 0,
		Count:    0,
		Data:     resps,
	}, nil
}

// Count banners
func (this bannersService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := db.GetDb().Model(&model.Banners{})
	var rows []model.Banners
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "banners bannersService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List banners列表
func (this bannersService) List(page response.PageReq, listReq model.BannersListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	query := db.GetDb().Model(&model.Banners{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.ImageUrl != "" {
		query = query.Where("image_url = ?", listReq.ImageUrl)
	}
	if listReq.LinkArticle != "" {
		query = query.Where("link_article = ?", listReq.LinkArticle)
	}
	if listReq.LinkUrl != "" {
		query = query.Where("link_url = ?", listReq.LinkUrl)
	}
	if listReq.DisplayOrder > 0 {
		query = query.Where("display_order = ?", listReq.DisplayOrder)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "banners bannersService  List Count err"); e != nil {
		return
	}
	// 数据
	var rows []model.Banners
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "banners bannersService  List Find err"); e != nil {
		return
	}
	resps := []model.BannersResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail banners详情
func (this bannersService) Detail(id int) (res model.BannersResp, e error) {
	var row model.Banners
	err := db.GetDb().Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "banners banners  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	return
}

// Add banners新增
func (this bannersService) Add(addReq model.BannersAddReq) (e error) {
	var row model.Banners
	response.CopyStruct(&row, addReq)
	err := db.GetDb().Create(&row).Error
	e = response.CheckErr(err, "banners bannersService  Add Create err")
	return
}

// Edit banners编辑
func (this bannersService) Edit(editReq model.BannersEditReq) (e error) {
	var row model.Banners
	err := db.GetDb().Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "banners bannersService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = db.GetDb().Model(&row).Updates(row).Error
	e = response.CheckErr(err, "banners bannersService Edit Updates err")

	return
}

// Change banners 状态切换
func (this bannersService) Change(changeReq model.BannersDetailReq) (e error) {
	var row model.Banners
	err := db.GetDb().Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "banners bannersService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = db.GetDb().Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = db.GetDb().Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "banners bannersService  Change Updates err")
	return
}

// Del banners删除
func (this bannersService) Del(id int) (e error) {
	var row model.Banners
	err := db.GetDb().Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "banners bannersService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "banners bannersService Del First err"); e != nil {
		return
	}
	// 删除
	err = db.GetDb().Delete(&row).Error
	e = response.CheckErr(err, "banners bannersService Del Delete err")
	return
}

// DelBatch banners 批量删除
func (this bannersService) DelBatch(delReq model.BannersDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := db.GetDb().Where("id IN (?)", delReq.Ids).Delete(model.Banners{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 BannersService DelBatch Delete err")
	return
}
