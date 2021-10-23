package itying

import (
	"fmt"
	"net/http"

	"aita.com/gin_demo/model"
	"github.com/gin-gonic/gin"
)

type DefaultController struct{}

func (con DefaultController) Index(C *gin.Context) {
	fmt.Printf("当前时间戳：%v", model.GetUnix())
	C.HTML(http.StatusOK, "index.html", gin.H{
		"title": "首页",
		"t":     1631726951,
	})
}

func (con DefaultController) News(C *gin.Context) {
	C.String(http.StatusOK, "新闻")
}
