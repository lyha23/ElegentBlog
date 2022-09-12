package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"net/http"
	"strconv"
)

// AddCategory 添加分类
func SaveCategory(c *gin.Context) {
	/*
	   "name":"新增分类"
	*/
	var ctl model.Category
	_ = c.ShouldBindJSON(&ctl)
	id, _ := strconv.Atoi(c.Param("id"))
	ctl.ID = uint(id)
	err := ctl.SaveCate()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
			"id":   ctl.ID,
		},
	)
}

// GetCate 查询分类列表
func GetCateList(c *gin.Context) {
	/*
		"name":"获取分类列表"
	*/
	var ctl model.CatEvent
	err := ctl.GetList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
			"list": ctl.CategoryList,
		},
	)
}

// DeleteCate 删除分类
func DeleteCate(c *gin.Context) {
	/*
	   DELETE  /category/:id
	*/
	var ctl model.Category
	id, _ := strconv.Atoi(c.Param("id"))
	ctl.ID = uint(id)
	if err := ctl.DeleteCate(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
		},
	)
}