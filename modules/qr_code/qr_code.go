package qrcode

import (
	"fmt"
	"net/http"

	"github.com/changwei4869/wedding/model"
	"github.com/gin-gonic/gin"
)

// AddQrCode 添加新二维码
// @summary 添加新二维码
// @description 添加新二维码
// @tags qrcode
// @accept application/json
// @produce application/json
// @param qrcode body model.QrCodesAddReq true "QrCode 信息"
// @success 201 {string} string "成功添加二维码"
// @router /qrcode [post]
func AddQrCode(c *gin.Context) {
	req := model.QrCodesAddReq{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if err := QrCodeIns.Add(req); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding qrcode to db: %s", err))
		return
	}

	c.String(http.StatusCreated, "created qrcode successfully")
}

// GetAllQrCode 列出所有二维码
// @summary 列出所有二维码
// @description 列出所有二维码
// @tags qrcode
// @produce application/json
// @success 200 {array} model.QrCodes "成功获取所有二维码"
// @router /qrcodes [get]
func GetAllQrCode(c *gin.Context) {
	res, err := QrCodeIns.All()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error getting qrcodes from db: %s", err))
		return
	}
	c.JSON(http.StatusOK, res)
}

// EditQrCode 编辑二维码
// @summary 编辑二维码
// @description 编辑二维码
// @tags qrcode
// @accept application/json
// @produce application/json
// @param qrcode body model.QrCodesEditReq true "Qrcode 信息"
// @success 200 {object} model.QrCodesEditReq "成功编辑二维码"
// @router /qrcode [put]
func EditQrCode(c *gin.Context) {
	req := model.QrCodesEditReq{}
	if err := c.BindJSON(&req); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	if err := QrCodeIns.Edit(req); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error editing qrcode to db: %s", err))
		return
	}

	c.JSON(http.StatusOK, req)
}
