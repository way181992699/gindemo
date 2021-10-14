package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func DebugWare() gin.HandlerFunc  {

	return debug
}

func debug(c *gin.Context)   {

	fmt.Println("1.打印日志...")

}
