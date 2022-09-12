package v1

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/middleware"
	"github.com/wejectchen/ginblog/model"
	"net/http"
	"time"
)

// Login 后台登陆
func Login(c *gin.Context) {
	/*
		POST /login
	*/
	var ctl model.User
	_ = c.ShouldBindJSON(&ctl)
	if err := ctl.CheckLogin(); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}
	token := setToken(c, ctl)
	c.JSON(http.StatusOK, gin.H{
		"code":  http.StatusOK,
		"data":  ctl,
		"token": token,
	})
}

// token生成函数
func setToken(c *gin.Context, user model.User) string {
	j := middleware.NewJWT()
	claims := middleware.MyClaims{
		Username: user.Username,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 100,
			ExpiresAt: time.Now().Unix() + 604800,
			Issuer:    "GinBlog",
		},
	}
	token, _ := j.CreateToken(claims)
	return token
}