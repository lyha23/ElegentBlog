package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/model"
	"github.com/wejectchen/ginblog/utils"
	"net/http"
	"strconv"
)

func init() {
	var user model.User
	user.Username = utils.AdminName
	user.Password = utils.AdminPwd
	_ = user.AddUserRecordIfNeeded()
}

// SaveUser 编辑用户
func SaveUser(c *gin.Context) {
	/*
	   "username":"root1",
	   "password":"root"
	*/
	var ctl model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&ctl)
	ctl.ID = uint(id)
	if err := ctl.SaveUser(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": ctl,
		},
	)
}

// GetUserInfo 查询单个用户
func GetUserInfo(c *gin.Context) {
	/*
		GET /api/v1/user/1
	*/
	id, _ := strconv.Atoi(c.Param("id"))
	var ctl model.User
	ctl.ID = uint(id)
	if err := ctl.GetUserByID(); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
			"data": ctl,
		},
	)

}

// GetUsers 查询用户列表
func GetUsers(c *gin.Context) {
	/*
	 */
	username := c.Query("username")
	var ctl model.UserEvent
	ctl.Username = username
	if err := ctl.GetUsers(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
			"list": ctl.UserList,
		},
	)
}

// ChangeUserPassword 修改密码
func ChangeUserPassword(c *gin.Context) {
	/*
		POST /api/v1/admin/changepw/1
		   "username":"root1",
		   "password":"root"
	*/
	var ctl model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&ctl)
	ctl.ID = uint(id)
	if err := ctl.ChangePassword(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
		},
	)
}

// DeleteUser 删除用户
func DeleteUser(c *gin.Context) {
	/*
		DELETE	localhost:8848/api/v1/user/10
	*/
	id, _ := strconv.Atoi(c.Param("id"))
	var ctl model.User
	ctl.ID = uint(id)
	if err := ctl.DeleteUser(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	c.JSON(
		http.StatusOK, gin.H{
			"code": http.StatusOK,
		},
	)
}