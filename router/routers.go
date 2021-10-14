package router

import (
	"aaa/middleware"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine {

	r := gin.Default()
	//全局使用中间件 Use()方法
	r.Use(middleware.DebugWare())

	//路由模块化，不同的路由处理不同的业务
	//1.julyRouter
	julyRouter(r)
	//2.cindyRouter
	cindyRouter(r)
	//3.handlerRouter
	handlerRouter(r)
return r

}

