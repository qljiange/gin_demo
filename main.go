package main

import (
	"html/template"

	"aita.com/gin_demo/model"
	"aita.com/gin_demo/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//自定义模板函数，放在加载templates模板前，勇于HTML渲染是调用该函数进行转换
	r.SetFuncMap(template.FuncMap{
		"UnixToTime": model.UnixToTime,
	})

	//指定HTML模板的位置
	r.LoadHTMLGlob("templates/*")

	//调用方法实现路由注册
	routers.DefaultRoutersInit(r)

	routers.ApiRoutersInit(r)

	routers.AdminRoutersInit(r)

    _ = r.Run()
}
