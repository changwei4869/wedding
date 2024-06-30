package category

import (
	"net/http"

	"github.com/changwei4869/wedding/model"
	"github.com/gin-gonic/gin"
)

// AddCategory 添加栏目
// @summary 添加栏目
// @description 添加栏目
// @tags category
// @accept application/json
// @produce application/json
// @param Category body model.CategoriesAddReq true "Category 信息"
// @success 201 {string} string "成功添加栏目"
// @router /category [post]
func AddCategory(c *gin.Context) {
	req := model.CategoriesAddReq{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if err := CategoriesIns.Add(req); err != nil {
		c.String(http.StatusBadRequest, err.Error())
		return
	}

	c.String(http.StatusOK, "created category successfully")
}
