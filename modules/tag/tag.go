package tag

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/changwei4869/wedding/model"

	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// GetTagById 根据id获取tag
// @summary 根据id获取tag
// @description 根据id获取tag
// @tags tag
// @param id path string true "id"
// @produce application/json
// @success 200 {object} model.TagsResp "成功获取tag"
// @router /tag/:id [get]
func GetTagById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "tag id empty")
		return
	}

	tagId, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}

	tag, err := NewTagsService(db.GetDb()).Detail(tagId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("get tag from db error: %s", err))
		return
	}
	c.JSON(http.StatusOK, tag)
}

// GetAllTags 获取数据库中所有标签
// @Summary 获取所有标签
// @Description 获取数据库中所有标签记录
// @Tags tag
// @Accept json
// @Produce json
// @Success 200 {array} model.TagsResp "成功获取所有标签"
// @Router /tag [get]
func GetAllTags(c *gin.Context) {
	// 创建标签服务

	tags, err := NewTagsService(db.GetDb()).GetAll()

	if err != nil {
		// 如果发生错误，返回错误响应
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// 返回获取的所有标签
	c.JSON(http.StatusOK, tags)
}

// AddTag 添加新tag
// @summary 添加新tag
// @description 添加新tag
// @tags tag
// @accept application/json
// @produce application/json
// @param tag body model.TagsAddReq true "Tag 信息"
// @success 201 {object} model.TagsAddReq "成功添加tag"
// @router /tag [post]
func AddTag(c *gin.Context) {
	tag := model.TagsAddReq{}
	// 从请求中解析 JSON 数据到 tag 结构体
	if err := c.BindJSON(&tag); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if err := NewTagsService(db.GetDb()).Add(tag); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding tag to db: %s", err))
		return
	}

	c.JSON(http.StatusCreated, tag) // 返回新创建的标签信息
}

// DeleteTag 删除指定id的tag
// @summary 删除指定id的tag
// @description 删除指定id的tag
// @tags tag
// @param id path string true "id"
// @produce text/plain
// @success 200 {string} string "成功删除tag"
// @router /tag/:id [delete]
func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "tag id empty")
		return
	}

	tagID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	err = NewTagsService(db.GetDb()).Del(tagID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete tag from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "tag deleted successfully")
}

// EditTag 编辑tag
// @summary 编辑tag
// @description 编辑tag
// @tags tag
// @accept application/json
// @produce application/json
// @param tag body model.TagsEditReq true "Tag 信息"
// @success 200 {object} model.TagsEditReq "成功编辑tag"
// @router /tag [put]
func EditTag(c *gin.Context) {
	// 获取更新后的tag
	var updatedTag model.TagsEditReq
	if err := c.BindJSON(&updatedTag); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	// 保存更新后的标签信息到数据库
	err := NewTagsService(db.GetDb()).Edit(updatedTag)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating tag in db: %s", err))
		return
	}

	// 返回更新后的标签信息
	c.JSON(http.StatusOK, updatedTag)
}
