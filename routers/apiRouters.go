package routers

import (
	"aita.com/gin_demo/controllers/api"
	"github.com/gin-gonic/gin"
)

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userslist", api.ApiController{}.Ulist)
		apiRouters.GET("/plist", api.ApiController{}.Plist)
		apiRouters.GET("/cart", api.ApiController{}.Cart)
	}
}
