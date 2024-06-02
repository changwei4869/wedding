package registration

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// @Summary Get registrations based on status and location
// @Description status 和 location放在quaty param中
// @Tags registrations
// @Accept json
// @Produce json
// @Param status query int false "Status of the registration"
// @Param location query string false "Location of the registration"
// @Success 200 {array} model.RegistrationsResp
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 500 {string} string "get registration from db error"
// @Router /registration [get]
func GetRegistration(c *gin.Context) {
	statusstr := c.Query("status")  // 获取status查询参数
	location := c.Query("location") // 获取location查询参数

	// 将 status 从字符串转换为整数
	var status int
	var err error
	if statusstr != "" {
		status, err = strconv.Atoi(statusstr)
		if err != nil {
			c.String(http.StatusBadRequest, "Invalid request parameters: status must be an integer")
			return
		}
	}

	var req model.RegistrationsDetailReq
	req.Status = status
	req.Location = location

	// 调用服务获取注册信息
	registrations, err := NewRegistrationsService(db.GetDb()).Detail(req)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("get registration from db error: %s", err))
		return
	}
	c.JSON(http.StatusOK, registrations)
}

// InitRegistration 创建新的registration
// @Summary Create a new registration
// @Description Create a new registration entry with the provided data
// @Tags registrations
// @Accept json
// @Produce json
// @Param registration body model.Registrations true "Registration data"
// @Success 200 {object} model.Registrations
// @Failure 400 {string} string "get registration req error"
// @Failure 500 {string} string "create registration to db error"
// @Router /registration/initregistration [post]
func InitRegistration(c *gin.Context) {
	var registration model.Registrations
	if err := c.ShouldBindJSON(&registration); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Sprintf("get registration req error: %s", err))
		return
	}

	// 设置创建时间和更新时间
	registration.CreatedAt = time.Now()
	registration.UpdatedAt = time.Now()

	// 插入数据库
	if err := NewRegistrationsService(db.GetDb()).Create(&registration); err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Sprintf("create registration to db error: %s", err))
		return
	}

	c.JSON(http.StatusOK, registration)
}

// GetAllRegistrations 获取数据库中所有注册信息
// @Summary 获取所有注册信息
// @Description 获取数据库中所有注册信息记录
// @Tags Registration
// @Accept json
// @Produce json
// @Success 200 {array} model.RegistrationsResp "成功获取所有注册信息"
// @Router /registrations [get]
func GetAllRegistrations(c *gin.Context) {
	registrations, err := NewRegistrationsService(db.GetDb()).GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, registrations)
}

// DeleteRegistration 删除指定id的registration
// @summary 删除指定id的registration
// @description 删除指定id的registration
// @tags DeleteRegistration
// @param id path string true "id"
// @produce text/plain
// @success 200 {string} string "成功删除registration"
// @router /registration/:id [delete]
func DeleteRegistration(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "registration id empty")
		return
	}

	registrationId, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	err = NewRegistrationsService(db.GetDb()).Del(registrationId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete registration from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "registration deleted successfully")
}

// EditRegistration 编辑registration
// @summary 编辑registration
// @description 编辑registration
// @tags EditRegistration
// @accept application/json
// @produce application/json
// @param registration body model.RegistrationsEditReq true "Registration 信息"
// @success 200 {object} model.RegistrationsEditReq "成功编辑registration"
// @router /registration [put]
func EditRegistration(c *gin.Context) {
	var updatedRegistration model.RegistrationsEditReq
	if err := c.BindJSON(&updatedRegistration); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	err := NewRegistrationsService(db.GetDb()).Edit(updatedRegistration)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating registration in db: %s", err))
		return
	}

	c.JSON(http.StatusOK, updatedRegistration)
}

// CheckRegistration 审核
// @Summary Check registration
// @Description 通过审核status值设为1，拒绝设为2。待审核为0
// @Tags registrations
// @Accept json
// @Produce json
// @Param id body int true "Registration ID"
// @Param reason body string false "Reason for approval or rejection"
// @Success 200 {object} model.RegistrationsCheckResp
// @Failure 400 {string} string "Invalid request parameters"
// @Failure 500 {string} string "Failed to update registration status"
// @Router /checkregistration [post]
func CheckRegistration(c *gin.Context) {
	var req model.RegistrationsCheckReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, model.RegistrationsCheckResp{
			Success: false,
			Message: "Invalid request parameters",
		})
		return
	}

	db := db.GetDb()
	var registration model.Registrations
	if err := db.First(&registration, req.Id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.RegistrationsCheckResp{
			Success: false,
			Message: "Registration not found",
		})
		return
	}

	registration.Status = req.Status
	registration.Reason = req.Reason

	if err := db.Save(&registration).Error; err != nil {
		c.JSON(http.StatusInternalServerError, model.RegistrationsCheckResp{
			Success: false,
			Message: "Failed to update registration status",
		})
		return
	}
	if req.Status == 1 {
		user := model.Users{
			Uid:                 fmt.Sprintf("user_%d", registration.Id), // 示例生成UID
			Name:                registration.Name,
			Phone:               registration.Phone,
			Wechat:              registration.Wechat,
			Gender:              registration.Gender,
			Age:                 registration.Age,
			Height:              registration.Height,
			Weight:              registration.Weight,
			Location:            registration.Location,
			Residence:           registration.Residence,
			Qualifications:      registration.Qualifications,
			SexualOrientation:   registration.SexualOrientation,
			MarriageHistory:     registration.MarriageHistory,
			Profession:          registration.Profession,
			AnnualIncome:        registration.AnnualIncome,
			AssetStatus:         registration.AssetStatus,
			SelfDescription:     registration.SelfDescription,
			MarriageCertificate: registration.MarriageCertificate,
			LiveTogether:        registration.LiveTogether,
			NeedChild:           registration.NeedChild,
			BridePrice:          registration.BridePrice,
			Distance:            registration.Distance,
			WeddingMode:         registration.WeddingMode,
			MarriedLife:         registration.MarriedLife,
			LookingFor:          registration.LookingFor,
			ExpectHelp:          registration.ExpectHelp,
			Channel:             registration.Channel,
		}

		if err := db.Create(&user).Error; err != nil {
			c.JSON(http.StatusInternalServerError, model.RegistrationsCheckResp{
				Success: false,
				Message: "Failed to create user from registration",
			})
			return
		}
	}
	c.JSON(http.StatusOK, model.RegistrationsCheckResp{
		Success: true,
		Message: "Registration status updated successfully",
	})
}

// DelBatchRegistrations 批量删除注册信息
// @summary 批量删除注册信息
// @description 批量删除提供的所有注册信息
// @tags DelBatchRegistrations
// @accept application/json
// @produce application/json
// @param delReq body model.RegistrationsDelBatchReq true "需要删除的注册信息ID列表"
// @success 200 {string} string "成功删除所有选中注册信息"
// @router /registration/batch_delete [delete]
func DelBatchRegistrations(c *gin.Context) {
	var delReq model.RegistrationsDelBatchReq
	if err := c.BindJSON(&delReq); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if len(delReq.Ids) == 0 {
		c.String(http.StatusBadRequest, "没有提供任何ID进行删除")
		return
	}

	if err := NewRegistrationsService(db.GetDb()).DelBatch(delReq); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error deleting registrations: %s", err))
		return
	}

	c.String(http.StatusOK, "成功删除所有选中注册信息")
}
