package router

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string
	Age int
}


func cindyRouter(r *gin.Engine)  {

	r.GET("/",func(c *gin.Context){
		c.String(200,"hello!")
	})

	r.GET("/ginTest",func(c *gin.Context){
		c.String(200,"good  ginTest")
	})

	//get 路径参数 "/user/hank"
	r.GET("/user/:name", func(c *gin.Context) {
		name:=c.Param("name")
		c.String(http.StatusOK,"Hello  %s",name) //Hello hank
	})

	//get query : 路径后缀加参数  ?name=hank&&role=student
	r.GET("/user/get", func(c *gin.Context) {
		name := c.Query("name")
		role := c.DefaultQuery("role","teacher")
		c.String(http.StatusOK,"%s is %s",name,role) //hank is student
	})

	//Post format 参数 (postman中的 Body -> form-data)
	r.POST("/form", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.DefaultPostForm("username","zhangsan")//设置默认值

		c.JSON(http.StatusOK,gin.H{
			"username":username,
			"password":password,
		})
	})

	//post 字典参数 (postman中的 Body -> raw)
	r.POST("/post", func(c *gin.Context) {
		var user User
		c.BindJSON(&user)
		fmt.Println("name" ,user.Name)
		fmt.Println("age" ,user.Age)
		c.JSON(http.StatusOK,gin.H{
			"name": user.Name,
			"age": user.Age,
		})
	})

}