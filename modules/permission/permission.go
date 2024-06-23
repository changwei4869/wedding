package permission

import (
	"fmt"
	"net/http"

	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// ListPermission 展示所有权限
// @Summary 展示所有权限
// @Description 展示数据库中权限表中的所有权限
// @Tags permission
// @Produce application/json
// @Success 200 {array} model.PermissionsListReq "成功获取权限列表"
// @Failure 500 {string} string "内部服务器错误"
// @Router /permissions [get]
func ListPermission(c *gin.Context) {
	permissions, err := NewPermissionService(db.GetDb()).ListAll()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("get permissions from db error: %s", err))
		return
	}
	c.JSON(http.StatusOK, permissions)
}
