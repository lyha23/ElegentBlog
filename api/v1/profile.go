package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"net/http"
	"strconv"
)

func GetProfile(c *gin.Context) {
	/*
		/api/v1/profile/1
	*/
	id, _ := strconv.Atoi(c.Param("id"))
	var ctl model.Profile
	ctl.ID = uint(id)
	if err := ctl.GetProfile(); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": ctl,
	})
}

func UpdateProfile(c *gin.Context) {
	/*
			PUT	/api/v1/profile/2
		"name":"Liyuhang",
		"desc":"myblog",
		"qq_chat":"83006652",
		"wechat":"lyha23",
		"bili":"safdasdf",
		"email":"lyha23@sjtu.edu.cn",
		"img":"ssssf",
		"avatar":"ss",

	*/
	var ctl model.Profile
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&ctl)
	ctl.ID = uint(id)
	if err := ctl.UpdateProfile(); err != nil {
		c.AbortWithStatusJSON(http.StatusOK, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": http.StatusOK,
		"data": ctl,
	})
}