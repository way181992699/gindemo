package main

import "aaa/router"

func main() {

	//初始化路由
	r := router.InitRouters()

	_ = r.Run(":8081")
}

