package admin

import (
	"fmt"
	"net/http"

	"aita.com/gin_demo/model"
	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Index(C *gin.Context) {
	//直接调用方法在控制台转换时间戳
	fmt.Println(model.UnixToTime(1631725575907))
	C.String(http.StatusOK, "后台首页")
}
