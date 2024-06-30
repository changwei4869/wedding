package banner

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/changwei4869/wedding/model"
	"github.com/gin-gonic/gin"
)

// AddBanner 添加轮播图
// @summary 添加轮播图
// @description 添加轮播图
// @tags banner
// @accept application/json
// @produce application/json
// @param Banner body model.BannersAddReq true "Banner 信息"
// @success 201 {string} string "成功添加轮播图"
// @router /banner [post]
func AddBanner(c *gin.Context) {
	req := model.BannersAddReq{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	displayOrder, err := BannerIns.GetMaxDisplayOrder()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error getting max display order: %s", err))
		return
	}
	req.DisplayOrder = displayOrder + 1

	if err := BannerIns.Add(req); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding Banner to db: %s", err))
		return
	}

	c.String(http.StatusCreated, "created Banner successfully")
}

// ListBanner 列出所有轮播图
// @summary 列出所有轮播图
// @description 列出所有轮播图
// @tags banner
// @produce application/json
// @success 200 {array} model.Banners "成功获取所有轮播图"
// @router /banner [get]
func ListBanner(c *gin.Context) {
	res, err := BannerIns.GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error getting banners from db: %s", err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// EditBanner 编辑轮播图
// @summary 编辑轮播图
// @description 编辑轮播图
// @tags banner
// @accept application/json
// @produce application/json
// @param Banner body model.BannersEditReq true "Banner 信息"
// @success 200 {string} string "success"
// @router /banner [put]
func EditBanner(c *gin.Context) {
	req := model.BannersEditReq{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	if err := BannerIns.Edit(req); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error editing Banner to db: %s", err))
		return
	}
	c.String(http.StatusOK, "success")
}

// DeleteBanner 删除指定id的轮播图
// @summary 删除指定id的轮播图
// @description 删除指定id的轮播图
// @tags banner
// @param id path string true "id"
// @produce text/plain
// @success 200 {string} string "成功删除轮播图"
// @router /banner/:id [delete]
func DeleteBanner(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "banner id empty")
		return
	}

	bannerID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	if err := BannerIns.Del(bannerID); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error deleting Banner to db: %s", err))
		return
	}
	c.String(http.StatusOK, "banner deleted successfully")
}

// BannerMoveUp 上移轮播图
// @summary 上移轮播图
// @description 上移轮播图
// @tags banner
// @accept application/json
// @produce application/json
// @success 200 {string} string "success"
// @router /banner/up/:id [put]
func BannerMoveUp(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "banner id empty")
		return
	}

	bannerID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	if err := BannerIns.Up(bannerID); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error moving up Banner to db: %s", err))
		return
	}
	c.String(http.StatusOK, "banner moved up successfully")
}

// BannerMoveDown 下移轮播图
// @summary 下移轮播图
// @description 下移轮播图
// @tags banner
// @accept application/json
// @produce application/json
// @success 200 {string} string "success"
// @router /banner/down/:id [put]
func BannerMoveDown(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "banner id empty")
		return
	}
	bannerID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	if err := BannerIns.Down(bannerID); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error moving down Banner to db: %s", err))
		return
	}
	c.String(http.StatusOK, "banner moved down successfully")
}