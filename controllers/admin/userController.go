package admin

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (con UserController) Index(C *gin.Context) {

	//对应控制器获取值
	//返回值：(value interface{}, exists bool)
	username, _ := C.Get("username")
	fmt.Println(username)

	//类型断言
	v, ok := username.(string)
	if ok {
		C.String(http.StatusOK, "用户列表"+v)
	} else {
		C.String(http.StatusOK, "用户列表获取失败")
	}

}

func (con UserController) Add(C *gin.Context) {
	C.String(http.StatusOK, "新增用户")
}

func (con UserController) Edit(C *gin.Context) {
	C.String(http.StatusOK, "修改用户")
}

func (con UserController) Delete(C *gin.Context) {
	C.String(http.StatusOK, "删除用户")
}
