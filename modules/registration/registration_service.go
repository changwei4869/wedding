package registration

import (
	"errors"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/utils/response"
	"gorm.io/gorm"
)

type IRegistrationsService interface {
	All() (res response.PageResp, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.RegistrationsListReq) (res response.PageResp, e error)
	Detail(id int) (res model.RegistrationsResp, e error)
	Add(addReq model.RegistrationsAddReq) (e error)
	Edit(editReq model.RegistrationsEditReq) (e error)
	Change(changeReq model.RegistrationsDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.RegistrationsDelBatchReq) (e error)
}

// NewRegistrationsService 初始化
func NewRegistrationsService(db *gorm.DB) IRegistrationsService {
	return &registrationsService{db: db}
}

// registrationsService registrations
type registrationsService struct {
	db *gorm.DB
}

// All registrations列表
func (this registrationsService) All() (res response.PageResp, e error) {
	// 数据
	query := this.db.Model(&model.Registrations{})
	var rows []model.Registrations
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "registrations registrationsService  All Find err"); e != nil {
		return
	}
	resps := []model.RegistrationsResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   0,
		PageSize: 0,
		Count:    0,
		Data:     resps,
	}, nil
}

// Count registrations
func (this registrationsService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := this.db.Model(&model.Registrations{})
	var rows []model.Registrations
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "registrations registrationsService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List registrations列表
func (this registrationsService) List(page response.PageReq, listReq model.RegistrationsListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	query := this.db.Model(&model.Registrations{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Number != "" {
		query = query.Where("number = ?", listReq.Number)
	}
	if listReq.Registrationscol != "" {
		query = query.Where("registrationscol = ?", listReq.Registrationscol)
	}
	if listReq.Name != "" {
		query = query.Where("name = ?", listReq.Name)
	}
	if listReq.Gender > 0 {
		query = query.Where("gender = ?", listReq.Gender)
	}
	if listReq.Phone != "" {
		query = query.Where("phone = ?", listReq.Phone)
	}
	if listReq.Wechat != "" {
		query = query.Where("wechat = ?", listReq.Wechat)
	}
	if listReq.Age > 0 {
		query = query.Where("age = ?", listReq.Age)
	}
	if listReq.Height > 0 {
		query = query.Where("height = ?", listReq.Height)
	}
	if listReq.Weight > 0 {
		query = query.Where("weight = ?", listReq.Weight)
	}
	if listReq.Location != "" {
		query = query.Where("location = ?", listReq.Location)
	}
	if listReq.Residence != "" {
		query = query.Where("residence = ?", listReq.Residence)
	}
	if listReq.Education != "" {
		query = query.Where("education = ?", listReq.Education)
	}
	if listReq.Qualifications != "" {
		query = query.Where("qualifications = ?", listReq.Qualifications)
	}
	if listReq.SexualOrientation != "" {
		query = query.Where("sexual_orientation = ?", listReq.SexualOrientation)
	}
	if listReq.MarriageHistory != "" {
		query = query.Where("marriage_history = ?", listReq.MarriageHistory)
	}
	if listReq.Profession != "" {
		query = query.Where("profession = ?", listReq.Profession)
	}
	if listReq.AnnualIncome != "" {
		query = query.Where("annual_income = ?", listReq.AnnualIncome)
	}
	if listReq.AssetStatus != "" {
		query = query.Where("asset_status = ?", listReq.AssetStatus)
	}
	if listReq.SelfDescription != "" {
		query = query.Where("self_description = ?", listReq.SelfDescription)
	}
	if listReq.MarriageCertificate != "" {
		query = query.Where("marriage_certificate = ?", listReq.MarriageCertificate)
	}
	if listReq.LiveTogether != "" {
		query = query.Where("live_together = ?", listReq.LiveTogether)
	}
	if listReq.NeedChild != "" {
		query = query.Where("need_child = ?", listReq.NeedChild)
	}
	if listReq.BridePrice != "" {
		query = query.Where("bride_price = ?", listReq.BridePrice)
	}
	if listReq.Distance != "" {
		query = query.Where("distance = ?", listReq.Distance)
	}
	if listReq.WeddingMode != "" {
		query = query.Where("wedding_mode = ?", listReq.WeddingMode)
	}
	if listReq.MarriedLife != "" {
		query = query.Where("married_life = ?", listReq.MarriedLife)
	}
	if listReq.LookingFor != "" {
		query = query.Where("looking_for = ?", listReq.LookingFor)
	}
	if listReq.ExpectHelp != "" {
		query = query.Where("expect_help = ?", listReq.ExpectHelp)
	}
	if listReq.Channel != "" {
		query = query.Where("channel = ?", listReq.Channel)
	}
	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "registrations registrationsService  List Count err"); e != nil {
		return
	}
	// 数据
	var rows []model.Registrations
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "registrations registrationsService  List Find err"); e != nil {
		return
	}
	resps := []model.RegistrationsResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail registrations详情
func (this registrationsService) Detail(id int) (res model.RegistrationsResp, e error) {
	var row model.Registrations
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "registrations registrations  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "registrations registrationsService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	return
}

// Add registrations新增
func (this registrationsService) Add(addReq model.RegistrationsAddReq) (e error) {
	var row model.Registrations
	response.CopyStruct(&row, addReq)
	err := this.db.Create(&row).Error
	e = response.CheckErr(err, "registrations registrationsService  Add Create err")
	return
}

// Edit registrations编辑
func (this registrationsService) Edit(editReq model.RegistrationsEditReq) (e error) {
	var row model.Registrations
	err := this.db.Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "registrations registrationsService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "registrations registrationsService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = this.db.Model(&row).Updates(row).Error
	e = response.CheckErr(err, "registrations registrationsService Edit Updates err")

	//强制更新 当IsShow=0
	//err = this.db.Model(&row).Select("IsShow").Updates(row).Error
	//e = response.CheckErr(err, "registrations registrationsService  Edit Updates err")
	//强制更新 isDisable=0
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": editReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	//e = response.CheckErr(err, "registrations registrationsService  Edit Updates err")
	return
}

// Change registrations 状态切换
func (this registrationsService) Change(changeReq model.RegistrationsDetailReq) (e error) {
	var row model.Registrations
	err := this.db.Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "registrations registrationsService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "registrations registrationsService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = this.db.Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "registrations registrationsService  Change Updates err")
	return
}

// Del registrations删除
func (this registrationsService) Del(id int) (e error) {
	var row model.Registrations
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "registrations registrationsService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "registrations registrationsService Del First err"); e != nil {
		return
	}
	// 删除
	err = this.db.Delete(&row).Error
	e = response.CheckErr(err, "registrations registrationsService Del Delete err")
	return
}

// DelBatch registrations 批量删除
func (this registrationsService) DelBatch(delReq model.RegistrationsDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := this.db.Where("id IN (?)", delReq.Ids).Delete(model.Registrations{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 RegistrationsService DelBatch Delete err")
	return
}
