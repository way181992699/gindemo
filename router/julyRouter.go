package router

import (
	"aaa/middleware"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func julyRouter(r *gin.Engine) {
	//Group分组，请求都在/july下,同时分组也能添加中间件(拦截器)
	j := r.Group("/july",middleware.NextWare(),gin.Logger(),gin.ErrorLogger())
	{
		j.GET("/info", info)
		j.GET("/birthday", birthday)
	}

}

func info(c *gin.Context) {
	fmt.Println("/july/info")
	c.JSON(http.StatusOK, gin.H{
		"name": "七月",
	})
}
func birthday(c *gin.Context) {
	fmt.Println("/july/birthday")
	c.JSON(http.StatusOK, gin.H{
		"date": "12月20号",
	})
}
