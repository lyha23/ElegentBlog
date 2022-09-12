package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"net/http"
	"strconv"
)

// SaveArticle 保存文章
func SaveArticle(c *gin.Context) {
	/*
		POST /article/add
			"title":"新的文章",
			"desc":"文章描述",
			"content":"文章内容",
			"img":"https://hosizuki.github.io/pictures/dream/railgun.jpg",
	*/
	var ctl model.Article
	_ = c.ShouldBindJSON(&ctl)
	id, err := strconv.Atoi(c.Param("id"))
	if err == nil {
		ctl.ID = uint(id)
	}
	if err := ctl.SaveArt(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"id":   ctl.ID,
	})
}

// GetCateArt 查询分类下的所有文章
func GetCateArt(c *gin.Context) {
	/*
		POST /article/list/:id
	*/
	var ctl model.Event
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	id, _ := strconv.Atoi(c.Param("id"))
	pagination := model.Pagination{
		PageSize: pageSize,
		Page:     pageNum,
	}
	pagination.DefaultRequest()
	article := model.Article{
		Cid: id,
	}
	ctl = model.Event{
		Pagination: pagination,
		Article:    article,
	}
	if err := ctl.GetCateArt(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"list":       ctl.ArticleList,
		"pagination": ctl.Pagination,
	})
}

// GetArtInfo 查询单个文章信息
func GetArtInfo(c *gin.Context) {
	/*
	   POST /article/info/:id
	*/
	id, _ := strconv.Atoi(c.Param("id"))
	ctl := model.Article{}
	ctl.ID = uint(id)
	err := ctl.GetArtInfo()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": ctl,
	})
}

// GetArt 查询文章列表
func GetArt(c *gin.Context) {
	/*
	   GET /article/list
	*/
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	title := c.Query("title")
	var pagination = model.Pagination{
		Page:     pageNum,
		PageSize: pageSize,
	}
	pagination.DefaultRequest()
	var ctl model.Event
	ctl.Pagination = pagination
	ctl.Title = title
	err := ctl.GetList()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":       http.StatusOK,
		"list":       ctl.ArticleList,
		"pagination": ctl.Pagination,
	})
}

func GetArtTimeline(c *gin.Context) {
	/*
	   GET /article/timeline
	*/
	var ctl model.Event
	err := ctl.GetArtTimeline()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"list": ctl.ArticleList,
	})
}

// DeleteArt 删除文章
func DeleteArt(c *gin.Context) {
	/*
		localhost:8848/api/v1/article/3
	*/
	var ctl model.Article
	id, _ := strconv.Atoi(c.Param("id"))
	ctl.ID = uint(id)
	if err := ctl.DeleteArt(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
	})
}