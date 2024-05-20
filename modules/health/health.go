package health

import (
	"github.com/changwei4869/wedding/model"
	"github.com/gin-gonic/gin"
)

// HealthCheck 健康检查
// @summary 健康检查
// @description 健康检查
// @tags   health
// @accept application/json
// @produce application/json
// @success 200 {object} model.HealthResp "健康检查成功"
// @router /health [get]
func HealthCheck(c *gin.Context) {
	c.JSON(200, model.HealthResp{
		Status: "up",
	})
}
