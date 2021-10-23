package admin

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleController struct{}

func (con ArticleController) Index(C *gin.Context) {
	C.String(http.StatusOK, "新闻列表")
}

func (con ArticleController) Add(C *gin.Context) {
	C.String(http.StatusOK, "新增新闻")
}

func (con ArticleController) Edit(C *gin.Context) {
	C.String(http.StatusOK, "修改新闻")
}

func (con ArticleController) Delete(C *gin.Context) {
	C.String(http.StatusOK, "删除新闻")
}
