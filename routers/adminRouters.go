package routers

import (
	"aita.com/gin_demo/controllers/admin"
	"aita.com/gin_demo/middlewares"
	"github.com/gin-gonic/gin"
)

func AdminRoutersInit(r *gin.Engine) {
	//方法一：直接在group后传入中间件
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	//方法二：使用gin.Engine.Use()传入中间件
	//r.Use(middlewares.InitMiddleware)

	{
		//IndexController
		adminRouters.GET("/", admin.IndexController{}.Index)

		//UserController
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)
		adminRouters.GET("/user/delete", admin.UserController{}.Delete)

		//ArticleController
		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.ArticleController{}.Edit)
		adminRouters.GET("/article/delete", admin.ArticleController{}.Delete)
	}

}
