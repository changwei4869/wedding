package role

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/changwei4869/wedding/model"

	"github.com/gin-gonic/gin"
)

// ListRole 列出所有角色
// @summary 列出所有角色
// @description 列出所有角色
// @tags role
// @produce application/json
// @success 200 {object} response.PageResp "成功获取所有角色"
// @router /role [get]
func ListRole(c *gin.Context) {
	res, err := RoleIns.All()
	if err != nil {
		c.String(http.StatusInternalServerError, "error fetching roles from db")
		return
	}
	c.JSON(http.StatusOK, res)
}

// AddRole 添加新角色
// @summary 添加新角色
// @description 添加新角色
// @tags role
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

	if err := RoleIns.Add(role); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding role to db: %s", err))
		return
	}

	c.JSON(http.StatusCreated, role)
}

// DeleteRole 删除指定id的角色
// @summary 删除指定id的角色
// @description 删除指定id的角色
// @tags role
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
	err = RoleIns.Del(roleID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete role from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "role deleted successfully")
}

// EditRole 编辑角色
// @summary 编辑角色
// @description 编辑角色
// @tags role
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
	err := RoleIns.Edit(updatedRole)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating role in db: %s", err))
		return
	}

	c.JSON(http.StatusOK, updatedRole)
}
