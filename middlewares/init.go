package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//定义用于路由组内的中间件
func InitMiddleware(c *gin.Context) {
	//打印当前访问时间
	fmt.Println(time.Now())
	//打印访问的URL
	fmt.Println(c.Request.URL)

	c.Set("username", "张三")

	//创建新的协程打印访问路径
	cCp := c.Copy()
	go func() {
		time.Sleep(2 * time.Second)

		fmt.Printf("Done! in path %v", cCp.Request.URL)
	}()
}
