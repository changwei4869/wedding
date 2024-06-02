package service //users  users //users

import (
	"errors"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/utils/response"
	"gorm.io/gorm"
)

type IUsersService interface {
	All() (res response.PageResp, e error)
	Count() (res map[string]interface{}, e error)
	List(page response.PageReq, listReq model.UsersListReq) (res response.PageResp, e error)
	Detail(id int) (res model.UsersResp, e error)
	Add(addReq model.UsersAddReq) (e error)
	Edit(editReq model.UsersEditReq) (e error)
	Change(changeReq model.UsersDetailReq) (e error)
	Del(id int) (e error)
	DelBatch(delReq model.UsersDelBatchReq) (e error)
}

// NewUsersService 初始化
func NewUsersService(db *gorm.DB) IUsersService {
	return &usersService{db: db}
}

// usersService users
type usersService struct {
	db *gorm.DB
}

// All users列表
func (this usersService) All() (res response.PageResp, e error) {
	// 数据
	query := this.db.Model(&model.Users{})
	var rows []model.Users
	err := query.Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "users usersService  All Find err"); e != nil {
		return
	}
	resps := []model.UsersResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   0,
		PageSize: 0,
		Count:    0,
		Data:     resps,
	}, nil
}

// Count users
func (this usersService) Count() (res map[string]interface{}, e error) {
	var Count int64
	query := this.db.Model(&model.Users{})
	var rows []model.Users
	err := query.Find(&rows).Count(&Count).Error
	if e = response.CheckErr(err, "users usersService  All Find err"); e != nil {
		return
	}
	return map[string]interface{}{
		"Count": Count,
	}, nil
}

// List users列表
func (this usersService) List(page response.PageReq, listReq model.UsersListReq) (res response.PageResp, e error) {
	// 分页信息
	limit := page.PageSize
	offset := page.PageSize * (page.PageNo - 1)
	// 查询
	query := this.db.Model(&model.Users{})
	if listReq.Id > 0 {
		query = query.Where("id = ?", listReq.Id)
	}
	if listReq.Uid != "" {
		query = query.Where("uid = ?", listReq.Uid)
	}
	if listReq.Avatar != "" {
		query = query.Where("avatar = ?", listReq.Avatar)
	}
	if listReq.Vitality > 0 {
		query = query.Where("vitality = ?", listReq.Vitality)
	}
	if listReq.Diamond > 0 {
		query = query.Where("diamond = ?", listReq.Diamond)
	}
	if listReq.Vip != "" {
		query = query.Where("vip = ?", listReq.Vip)
	}
	if listReq.Name != "" {
		query = query.Where("name = ?", listReq.Name)
	}
	if listReq.Password != "" {
		query = query.Where("password = ?", listReq.Password)
	}
	if listReq.RoleId > 0 {
		query = query.Where("role_id = ?", listReq.RoleId)
	}
	if listReq.Phone != "" {
		query = query.Where("phone = ?", listReq.Phone)
	}
	if listReq.Wechat != "" {
		query = query.Where("wechat = ?", listReq.Wechat)
	}
	if listReq.Gender > 0 {
		query = query.Where("gender = ?", listReq.Gender)
	}
	if listReq.Blacklist > 0 {
		query = query.Where("blacklist = ?", listReq.Blacklist)
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
	if listReq.LogOut != "" {
		query = query.Where("log_out = ?", listReq.LogOut)
	}
	if listReq.CreatedBy != "" {
		query = query.Where("created_by = ?", listReq.CreatedBy)
	}

	// 总数
	var count int64
	err := query.Count(&count).Error
	if e = response.CheckErr(err, "users usersService  List Count err"); e != nil {
		return
	}
	// 数据
	var rows []model.Users
	err = query.Limit(limit).Offset(offset).Order("id desc").Find(&rows).Error
	if e = response.CheckErr(err, "users usersService  List Find err"); e != nil {
		return
	}
	resps := []model.UsersResp{}
	response.CopyStruct(&resps, rows)
	return response.PageResp{
		PageNo:   page.PageNo,
		PageSize: page.PageSize,
		Count:    count,
		Data:     resps,
	}, nil
}

// Detail users详情
func (this usersService) Detail(id int) (res model.UsersResp, e error) {
	var row model.Users
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	if e = response.ErrRecordNotFound(err, "users users  Detail ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "users usersService  Detail First err"); e != nil {
		return
	}
	response.CopyStruct(&res, row)
	//URL转换
	// res.avatar = utils.Url.ToAbsoluteUrl(res.avatar)
	return
}

// Add users新增
func (this usersService) Add(addReq model.UsersAddReq) (e error) {
	var row model.Users
	response.CopyStruct(&row, addReq)
	err := this.db.Create(&row).Error
	e = response.CheckErr(err, "users usersService  Add Create err")
	return
}

// Edit users编辑
func (this usersService) Edit(editReq model.UsersEditReq) (e error) {
	var row model.Users
	err := this.db.Where("id = ?", editReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "users usersService Edit ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "users usersService  Edit First Err"); e != nil {
		return
	}
	// 更新
	response.CopyStruct(&row, editReq)
	err = this.db.Model(&row).Updates(row).Error
	e = response.CheckErr(err, "users usersService Edit Updates err")

	//强制更新 当IsShow=0
	//err = this.db.Model(&row).Select("IsShow").Updates(row).Error
	//e = response.CheckErr(err, "users usersService  Edit Updates err")
	//强制更新 isDisable=0
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": editReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	//e = response.CheckErr(err, "users usersService  Edit Updates err")
	return
}

// Change users 状态切换
func (this usersService) Change(changeReq model.UsersDetailReq) (e error) {
	var row model.Users
	err := this.db.Where("id = ?", changeReq.Id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "users usersService Change ErrRecordNotFound!(id="+strconv.Itoa(int(changeReq.Id))+")"); e != nil {
		return
	}
	if e = response.CheckErr(err, "users usersService  Change Err"); e != nil {
		return
	}
	// 更新
	//err = this.db.Model(&row).Select("Enabled").Updates(row).Error
	//e = response.CheckErr(err, "广告 adService  Edit Updates err")
	//err = this.db.Model(&row).Updates(map[string]interface{}{"IsDisable": changeReq.isDisable, "UpdateTime": time.Now().Unix()}).Error
	// e = response.CheckErr(err, "users usersService  Change Updates err")
	return
}

// Del users删除
func (this usersService) Del(id int) (e error) {
	var row model.Users
	err := this.db.Where("id = ?", id).Limit(1).First(&row).Error
	// 校验
	if e = response.ErrRecordNotFound(err, "users usersService Del ErrRecordNotFound!"); e != nil {
		return
	}
	if e = response.CheckErr(err, "users usersService Del First err"); e != nil {
		return
	}
	// 删除
	err = this.db.Delete(&row).Error
	e = response.CheckErr(err, "users usersService Del Delete err")
	return
}

// DelBatch users 批量删除
func (this usersService) DelBatch(delReq model.UsersDelBatchReq) (e error) {
	// 校验ID列表是否为空
	if len(delReq.Ids) == 0 {
		err := errors.New("没有提供任何ID进行删除")
		response.CheckErr(err, "时间段 timeslotService DelBatch err")
		return
	}
	// 执行批量删除
	err := this.db.Where("id IN (?)", delReq.Ids).Delete(model.Users{}).Error
	// 检查并处理错误
	e = response.CheckErr(err, "时间段 UsersService DelBatch Delete err")
	return
}
