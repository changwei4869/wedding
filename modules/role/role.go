package role

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/changwei4869/wedding/utils/response"

	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// ListRole 列出所有角色
// @summary 列出所有角色
// @description 列出所有角色
// @tags ListRole
// @produce application/json
// @param pageNo query int false "页码"
// @param pageSize query int false "每页数量"
// @param id query int false "角色ID"
// @param name query string false "角色名称"
// @param description query string false "角色描述"
// @param createdAt query string false "创建时间"
// @param updatedAt query string false "更新时间"
// @success 200 {object} response.PageResp "成功获取所有角色"
// @router /role [get]
func ListRole(c *gin.Context) {
	var pageReq response.PageReq
	var listReq model.RolesListReq

	pageNo := c.DefaultQuery("pageNo", "1")
	pageSize := c.DefaultQuery("pageSize", "10")
	pageReq.PageNo, _ = strconv.Atoi(pageNo)
	pageReq.PageSize, _ = strconv.Atoi(pageSize)

	if id := c.Query("id"); id != "" {
		listReq.Id, _ = strconv.Atoi(id)
	}
	listReq.Name = c.Query("name")
	listReq.Description = c.Query("description")
	listReq.CreatedAt = c.Query("createdAt")
	listReq.UpdatedAt = c.Query("updatedAt")

	res, err := NewRolesService(db.GetDb()).List(pageReq, listReq)
	if err != nil {
		c.String(http.StatusInternalServerError, "error fetching roles from db")
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddRole 添加新角色
// @summary 添加新角色
// @description 添加新角色
// @tags AddRole
// @accept application/json
// @produce application/json
// @param role body model.RolesAddReq true "Role 信息"
// @success 201 {object} model.RolesAddReq "成功添加角色"
// @router /role [post]
func AddRole(c *gin.Context) {
	role := model.RolesAddReq{}
	if err := c.BindJSON(&role); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if err := NewRolesService(db.GetDb()).Add(role); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding role to db: %s", err))
		return
	}

	c.JSON(http.StatusCreated, role)
}

// DeleteRole 删除指定id的角色
// @summary 删除指定id的角色
// @description 删除指定id的角色
// @tags DeleteRole
// @param id path string true "id"
// @produce text/plain
// @success 200 {string} string "成功删除角色"
// @router /role/:id [delete]
func DeleteRole(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "role id empty")
		return
	}

	roleID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	err = NewRolesService(db.GetDb()).Del(roleID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete role from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "role deleted successfully")
}

// EditRole 编辑角色
// @summary 编辑角色
// @description 编辑角色
// @tags EditRole
// @accept application/json
// @produce application/json
// @param role body model.RolesEditReq true "Role 信息"
// @success 200 {object} model.RolesEditReq "成功编辑角色"
// @router /role [put]
func EditRole(c *gin.Context) {
	var updatedRole model.RolesEditReq
	if err := c.BindJSON(&updatedRole); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	err := NewRolesService(db.GetDb()).Edit(updatedRole)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating role in db: %s", err))
		return
	}

	c.JSON(http.StatusOK, updatedRole)
}
