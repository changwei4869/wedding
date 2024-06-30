package qrcode

import (
	"errors"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
	"github.com/changwei4869/wedding/utils/response"
)

type IQrCodesService interface {
	All() (res []model.QrCodes, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.QrCodesListReq) (res response.PageResp, e error)
	Detail(id int) (res model.QrCodesResp, e error)
	Add(addReq model.QrCodesAddReq) (e error)
	Edit(editReq model.QrCodesEditReq) (e error)
	Change(changeReq model.QrCodesDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.QrCodesDelBatchReq) (e error)
}

// qrCodesService qrCodes
type qrCodesService struct{}

var QrCodeIns IQrCodesService = qrCodesService{}

// All qrCodes列表
func (this qrCodesService) All() (res []model.QrCodes, e error) {
	// 数据
	query := db.GetDb().Model(&model.QrCodes{})
	var rows []model.QrCodes
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "qrCodes qrCodesService  All Find err"); e != nil {
		return
	}

	return rows, nil
}

// Count qrCodes
func (this qrCodesService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := db.GetDb().Model(&model.QrCodes{})
	var rows []model.QrCodes
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "qrCodes qrCodesService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List qrCodes列表
func (this qrCodesService) List(page response.PageReq, listReq model.QrCodesListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	query := db.GetDb().Model(&model.QrCodes{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.QrCode != "" {
		query = query.Where("qr_code = ?", listReq.QrCode)
	}
	if listReq.Location != "" {
		query = query.Where("location = ?", listReq.Location)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "qrCodes qrCodesService  List Count err"); e != nil {
		return
	}
	// 数据
	var rows []model.QrCodes
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "qrCodes qrCodesService  List Find err"); e != nil {
		return
	}
	resps := []model.QrCodesResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail qrCodes详情
func (this qrCodesService) Detail(id int) (res model.QrCodesResp, e error) {
	var row model.QrCodes
	err := db.GetDb().Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "qrCodes qrCodes  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "qrCodes qrCodesService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	return
}

// Add qrCodes新增
func (this qrCodesService) Add(addReq model.QrCodesAddReq) (e error) {
	var row model.QrCodes
	response.CopyStruct(&row, addReq)
	err := db.GetDb().Create(&row).Error
	e = response.CheckErr(err, "qrCodes qrCodesService  Add Create err")
	return
}

// Edit qrCodes编辑
func (this qrCodesService) Edit(editReq model.QrCodesEditReq) (e error) {
	var row model.QrCodes
	err := db.GetDb().Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "qrCodes qrCodesService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "qrCodes qrCodesService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = db.GetDb().Model(&row).Updates(row).Error
	e = response.CheckErr(err, "qrCodes qrCodesService Edit Updates err")

	//强制更新 当IsShow=0
	//err = db.GetDb().Model(&row).Select("IsShow").Updates(row).Error
	//e = response.CheckErr(err, "qrCodes qrCodesService  Edit Updates err")
	//强制更新 isDisable=0
	//err = db.GetDb().Model(&row).Updates(map[string]interface{}{"IsDisable": editReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	//e = response.CheckErr(err, "qrCodes qrCodesService  Edit Updates err")
	return
}

// Change qrCodes 状态切换
func (this qrCodesService) Change(changeReq model.QrCodesDetailReq) (e error) {
	var row model.QrCodes
	err := db.GetDb().Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "qrCodes qrCodesService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "qrCodes qrCodesService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = db.GetDb().Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = db.GetDb().Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "qrCodes qrCodesService  Change Updates err")
	return
}

// Del qrCodes删除
func (this qrCodesService) Del(id int) (e error) {
	var row model.QrCodes
	err := db.GetDb().Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "qrCodes qrCodesService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "qrCodes qrCodesService Del First err"); e != nil {
		return
	}
	// 删除
	err = db.GetDb().Delete(&row).Error
	e = response.CheckErr(err, "qrCodes qrCodesService Del Delete err")
	return
}

// DelBatch qrCodes 批量删除
func (this qrCodesService) DelBatch(delReq model.QrCodesDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := db.GetDb().Where("id IN (?)", delReq.Ids).Delete(model.QrCodes{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 QrCodesService DelBatch Delete err")
	return
}
