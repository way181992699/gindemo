package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func CustomWare() gin.HandlerFunc  {
	return custom

}

func custom(c *gin.Context) {
	fmt.Println("2.自定义中间件...")
}