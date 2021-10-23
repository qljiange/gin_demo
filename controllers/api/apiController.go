package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ApiController struct{}

func (con ApiController) Index(C *gin.Context) {
	C.String(http.StatusOK, "这是一个api总览页接口")
}

func (con ApiController) Ulist(C *gin.Context) {
	C.String(http.StatusOK, "这是一个api接口用户列表页接口")
}

func (con ApiController) Plist(C *gin.Context) {
	C.String(http.StatusOK, "这是一个api接口列表总览页接口")
}

func (con ApiController) Cart(C *gin.Context) {
	C.String(http.StatusOK, "这是一个api接口新增接口")
}
