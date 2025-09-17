package main

import "yongdeng-ecology-visualization/backend/routes"

func main() {

	// 启动路由并传递数据库连接
	r := routes.SetupRouter()
	r.Run(":8080") // 启动服务器
}
