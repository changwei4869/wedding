package admin

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/changwei4869/wedding/middleware"
	"github.com/changwei4869/wedding/utils"
	"github.com/dgrijalva/jwt-go"

	"github.com/changwei4869/wedding/model"
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

	admin, err := AdminIns.GetAdminByPhone(login.Phone)
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
// @success 200 {array} model.Admins "成功获取所有管理员"
// @router /admin [get]
func ListAdmin(c *gin.Context) {
	res, err := AdminIns.All()
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
	if err := AdminIns.Add(admin); err != nil {
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
	err = AdminIns.Del(adminID)
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
	err := AdminIns.Edit(updatedAdmin)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating admin in db: %s", err))
		return
	}

	c.JSON(http.StatusOK, updatedAdmin)
}
