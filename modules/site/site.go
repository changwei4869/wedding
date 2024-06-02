package site

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/changwei4869/wedding/model"

	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// GetSiteById 根据id获取site
// @summary 根据id获取site
// @description 根据id获取site
// @tags site
// @param id path string true "id"
// @produce application/json
// @success 200 {object} model.SitesResp "成功获取site"
// @router /site/:id [get]
func GetSiteById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "site id empty")
		return
	}

	siteId, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}

	site, err := NewSiteService(db.GetDb()).Detail(siteId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("get site from db error: %s", err))
		return
	}
	c.JSON(http.StatusOK, site)
}

// AddSite 添加新site
// @summary 添加新site
// @description 添加新site
// @tags site
// @accept application/json
// @produce application/json
// @param site body model.SitesAddReq true "Site 信息"
// @success 201 {object} model.SitesAddReq "成功添加site"
// @router /site [post]
func AddSite(c *gin.Context) {
	site := model.SitesAddReq{}
	// 从请求中解析 JSON 数据到 site 结构体
	if err := c.BindJSON(&site); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if err := NewSiteService(db.GetDb()).Add(site); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding site to db: %s", err))
		return
	}

	c.JSON(http.StatusCreated, site) // 返回新创建的站点信息
}

// DeleteSite 删除指定id的site
// @summary 删除指定id的site。
// @description 删除指定id的site
// @tags site
// @param id path string true "id"
// @produce text/plain
// @success 200 {string} string "成功删除site"
// @router /site/:id [delete]
func DeleteSite(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "site id empty")
		return
	}

	siteID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	err = NewSiteService(db.GetDb()).Del(siteID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete site from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "site deleted successfully")
}

// GetAllSites 从数据库中获取所有站点记录
// @Summary 获取所有站点
// @Description 从数据库中获取所有站点记录
// @Tags site
// @Accept json
// @Produce json
// @Success 200 {array} model.SitesResp "成功获取所有站点"
// @Router /site [get]
func GetAllSites(c *gin.Context) {
	sites, err := NewSiteService(db.GetDb()).GetAll()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("get all sites from db error: %s", err))
		return
	}
	c.JSON(http.StatusOK, sites)
}

// EditSite 编辑site
// @summary 编辑site
// @description 编辑site
// @tags site
// @accept application/json
// @produce application/json
// @param site body model.SitesEditReq true "Site 信息"
// @success 200 {object} model.SitesEditReq "成功编辑site"
// @router /site [put]
func EditSite(c *gin.Context) {
	// 获取更新后的site
	var updatedSite model.SitesEditReq
	if err := c.BindJSON(&updatedSite); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	// 保存更新后的站点信息到数据库
	err := NewSiteService(db.GetDb()).Edit(updatedSite)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating site in db: %s", err))
		return
	}

	// 返回更新后的站点信息
	c.JSON(http.StatusOK, updatedSite)
}
