package tag

import (
	"errors"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/utils/response"
	"gorm.io/gorm"
)

type ITagsService interface {
	GetAll() (res response.PageResp, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.TagsListReq) (res response.PageResp, e error)
	Detail(id int) (res model.TagsResp, e error)
	Add(addReq model.TagsAddReq) (e error)
	Edit(editReq model.TagsEditReq) (e error)
	Change(changeReq model.TagsDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.TagsDelBatchReq) (e error)
}

// tagsService tags
type tagsService struct {
	db *gorm.DB
}

// NewTagsService 初始化
func NewTagsService(db *gorm.DB) *tagsService {
	return &tagsService{db: db}
}

// All tags列表
func (this tagsService) GetAll() ([]model.TagsResp, error) {
	var tags []model.Tags
	if err := this.db.Find(&tags).Error; err != nil {
		return nil, err
	}

	var tagsResp []model.TagsResp
	for _, tag := range tags {
		tagsResp = append(tagsResp, model.TagsResp{
			Id:     tag.Id,
			Name:   tag.Name,
			Gender: tag.Gender,
		})
	}
	return tagsResp, nil
}

// Count tags
func (this tagsService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := this.db.Model(&model.Tags{})
	var rows []model.Tags
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "tags tagsService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List tags列表
func (this tagsService) List(page response.PageReq, listReq model.TagsListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	query := this.db.Model(&model.Tags{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Name != "" {
		query = query.Where("name = ?", listReq.Name)
	}
	if listReq.Gender != "" {
		query = query.Where("gender = ?", listReq.Gender)
	}
	if len(listReq.CreatedAt) > 0 {
		query = query.Where("created_at = ?", listReq.CreatedAt)
	}
	if len(listReq.UpdatedAt) > 0 {
		query = query.Where("updated_at = ?", listReq.UpdatedAt)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "tags tagsService  List Count err"); e != nil {
		return
	}
	// 数据
	var rows []model.Tags
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "tags tagsService  List Find err"); e != nil {
		return
	}
	resps := []model.TagsResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail tags详情
func (this tagsService) Detail(id int) (res model.TagsResp, e error) {
	var row model.Tags
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "tags tags  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "tags tagsService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	return
}

// Add tags新增
func (this tagsService) Add(addReq model.TagsAddReq) (e error) {
	var row model.Tags
	response.CopyStruct(&row, addReq)
	err := this.db.Create(&row).Error
	e = response.CheckErr(err, "tags tagsService  Add Create err")
	return
}

// Edit tags编辑
func (this tagsService) Edit(editReq model.TagsEditReq) (e error) {
	var row model.Tags
	err := this.db.Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "tags tagsService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "tags tagsService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = this.db.Model(&row).Updates(row).Error
	e = response.CheckErr(err, "tags tagsService Edit Updates err")

	//强制更新 当IsShow=0
	//err = this.db.Model(&row).Select("IsShow").Updates(row).Error
	//e = response.CheckErr(err, "tags tagsService  Edit Updates err")
	//强制更新 isDisable=0
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": editReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	//e = response.CheckErr(err, "tags tagsService  Edit Updates err")
	return
}

// Change tags 状态切换
func (this tagsService) Change(changeReq model.TagsDetailReq) (e error) {
	var row model.Tags
	err := this.db.Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "tags tagsService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "tags tagsService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = this.db.Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "tags tagsService  Change Updates err")
	return
}

// Del tags删除
func (this tagsService) Del(id int) (e error) {
	var row model.Tags
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "tags tagsService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "tags tagsService Del First err"); e != nil {
		return
	}
	// 删除
	err = this.db.Delete(&row).Error
	e = response.CheckErr(err, "tags tagsService Del Delete err")
	return
}

// DelBatch tags 批量删除
func (this tagsService) DelBatch(delReq model.TagsDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := this.db.Where("id IN (?)", delReq.Ids).Delete(model.Tags{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 TagsService DelBatch Delete err")
	return
}
