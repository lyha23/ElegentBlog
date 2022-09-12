package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"net/http"
)

// UpLoad 上传图片接口
func UpLoad(c *gin.Context) {
	file, fileHeader, _ := c.Request.FormFile("file")
	url, err := model.Upload(file, fileHeader)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"url":  url,
		"name": fileHeader.Filename,
	})
}