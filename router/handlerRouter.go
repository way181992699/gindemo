package router

import (
	"aaa/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)


func handlerRouter(r *gin.Engine)  {

		//1.简单
		r.GET("/handler", func(c *gin.Context) {
			fmt.Println("第一步")
		},info1)
		//2.通用
		r.GET("/handler2",middleware.DebugWare(),middleware.CustomWare(),info2)
}

func info1(c *gin.Context) {
	c.JSON(http.StatusOK,gin.H{
		"info1":"info",
	})
}

func info2(c *gin.Context) {
	fmt.Println("info2")
	c.JSON(http.StatusOK,gin.H{
		"info2":"info2",
	})
}
