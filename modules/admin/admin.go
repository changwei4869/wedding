package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/utils"
	"github.com/changwei4869/wedding/utils/response"
	"github.com/dgrijalva/jwt-go"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// LoginAdmin 管理员登陆
// @summary 管理员登陆
// @description 管理员登陆
// @tags admin
// @accept application/json
// @produce application/json
// @param admin body model.AdminLoginReq true "Admin 信息"
// @success 200 {string} string "管理员成功登陆"
// @router /admin/login [post]
func AdminLogin(c *gin.Context) {
	login := model.AdminLoginReq{}
	if err := c.BindJSON(&login); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	admin, err := NewAdminsService(db.GetDb()).GetAdminByPhone(login.Phone)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error getting admin from db: %s", err))
		return
	}
	if admin.Password != utils.ComputeMD5(login.Password) {
		c.String(http.StatusUnauthorized, "password incorrect")
		return
	}
	token, err := middleware.NewJWT().CreateToken(middleware.MyClaims{
		Phone: login.Phone,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
		},
	})
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error creating token: %s", err))
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"token":  token,
	})
}

// ListAdmin 列出所有管理员
// @summary 列出所有管理员
// @description 列出所有管理员
// @tags admin
// @produce application/json
// @param pageNo query int false "页码"
// @param pageSize query int false "每页数量"
// @param id query int false "管理员ID"
// @param name query string false "管理员名称"
// @param email query string false "管理员邮箱"
// @param createdAt query string false "创建时间"
// @param updatedAt query string false "更新时间"
// @success 200 {object} response.PageResp "成功获取所有管理员"
// @router /admin [get]
func ListAdmin(c *gin.Context) {
	var pageReq response.PageReq
	var listReq model.AdminListReq
	var err error

	pageNo := c.DefaultQuery("pageNo", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageReq.PageNo, err = strconv.Atoi(pageNo)
	if err != nil {
		c.String(http.StatusBadRequest, "pageNo is not a number")
		return
	}
	pageReq.PageSize, err = strconv.Atoi(pageSize)
	if err != nil {
		c.String(http.StatusBadRequest, "pageSize is not a number")
		return
	}

	if id := c.Query("id"); id != "" {
		listReq.Id, err = strconv.Atoi(id)
		if err != nil {
			c.String(http.StatusBadRequest, "id is not a number")
			return
		}
	}
	listReq.Name = c.Query("name")
	listReq.Phone = c.Query("phone")
	//listReq.Role_id = c.Query("role_id")
	roleIDStr := c.Query("role_id")
	if roleIDStr != "" {
		roleID, err := strconv.Atoi(roleIDStr)
		if err != nil {
			// 处理字符串转换为整数错误
			c.String(http.StatusInternalServerError, "failed string to int")
			return
		}
		listReq.Role_id = roleID
	}
	//listReq.Status = c.Query("status")
	statusStr := c.Query("status")
	if statusStr != "" {
		status, err := strconv.Atoi(statusStr)
		if err != nil {
			// 处理字符串转换为整数错误
			c.String(http.StatusInternalServerError, "failed string to int")
			return
		}
		listReq.Status = status
	}

	res, err := NewAdminsService(db.GetDb()).List(pageReq, listReq)
	if err != nil {
		c.String(http.StatusInternalServerError, "error fetching admins from db")
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddAdmin 添加新管理员
// @summary 添加新管理员
// @description 添加新管理员
// @tags admin
// @accept application/json
// @produce application/json
// @param admin body model.AdminAddReq true "Admin 信息"
// @success 201 {object} model.AdminAddReq "成功添加管理员"
// @router /admin [post]
func AddAdmin(c *gin.Context) {
	admin := model.AdminAddReq{}
	if err := c.BindJSON(&admin); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	admin.Password = utils.ComputeMD5(admin.Password)
	if err := NewAdminsService(db.GetDb()).Add(admin); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding admin to db: %s", err))
		return
	}

	c.String(http.StatusCreated, "created admin successfully")
}

// DeleteAdmin 删除指定id的管理员
// @summary 删除指定id的管理员
// @description 删除指定id的管理员
// @tags admin
// @param id path string true "id"
// @produce text/plain
// @success 200 {string} string "成功删除管理员"
// @router /admin/:id [delete]
func DeleteAdmin(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "admin id empty")
		return
	}

	adminID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	err = NewAdminsService(db.GetDb()).Del(adminID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete admin from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "admin deleted successfully")
}

// EditAdmin 编辑管理员
// @summary 编辑管理员
// @description 编辑管理员
// @tags admin
// @accept application/json
// @produce application/json
// @param admin body model.AdminEditReq true "Admin 信息"
// @success 200 {object} model.AdminEditReq "成功编辑管理员"
// @router /admin [put]
func EditAdmin(c *gin.Context) {
	var updatedAdmin model.AdminEditReq
	if err := c.BindJSON(&updatedAdmin); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	err := NewAdminsService(db.GetDb()).Edit(updatedAdmin)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating admin in db: %s", err))
		return
	}

	c.JSON(http.StatusOK, updatedAdmin)
}
