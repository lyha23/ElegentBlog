package routes

import (
	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
	"github.com/wejectchen/ginblog/middleware"
	"github.com/wejectchen/ginblog/routes/api"
	"github.com/wejectchen/ginblog/utils"
)

func createMyRender() multitemplate.Renderer {
	p := multitemplate.NewRenderer()
	p.AddFromFiles("admin", "web/admin/dist/index.html")
	p.AddFromFiles("front", "web/front/dist/index.html")
	return p
}

func InitRouter() *gin.Engine {
	gin.SetMode(utils.AppMode)
	r := gin.Default()
	// 设置信任网络 []string
	// nil 为不计算，避免性能消耗，上线应当设置
	_ = r.SetTrustedProxies(nil)

	r.HTMLRender = createMyRender()
	r.Use(middleware.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	r.Static("/assets", "./web/front/dist/assets")
	r.Static("/static", "./web/front/dist/static")
	r.Static("/admin", "./web/admin/dist")
	r.StaticFile("/favicon.ico", "/web/front/dist/favicon.ico")
	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "front", nil)
	})
	r.GET("/admin", func(c *gin.Context) {
		c.HTML(200, "admin", nil)
	})

	// 注册路由
	api.RegisterBlogRouters(r)
	// 注册管理路由
	api.RegisterBlogManageRouter(r)

	return r

}