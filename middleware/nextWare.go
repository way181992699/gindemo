package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func NextWare() gin.HandlerFunc {

	return next
}

func next(c *gin.Context) {
	//Next()放可以阔以在请求处理完成后，在执行的中间件，也称为后置处理
	c.Next()
	fmt.Println("后置处理...")
}
