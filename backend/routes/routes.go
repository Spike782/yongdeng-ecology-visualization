package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"yongdeng-ecology-visualization/config"
	"yongdeng-ecology-visualization/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//添加 CORS 跨域中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		// 允许的请求方法
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		// 允许的请求头
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		// 允许前端读取的响应头
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		// 允许携带 Cookie
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		// 处理预检请求（OPTIONS 方法）
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent) // 204 No Content
			return
		}

		c.Next() // 继续处理后续请求
	})

	dbConfig, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}
	db := dbConfig

	// 使用中间件注入数据库连接
	r.Use(func(c *gin.Context) {
		c.Set("db", db) // 将数据库连接保存到上下文中
		c.Next()        // 调用后续处理程序
	})

	r.POST("/pi/auth/register", controllers.RegisterUser)
	r.POST("/api/auth/login", controllers.Login)

	// 根据 Gridcode 获取数据
	r.GET("/api/riskusage/gridcode", controllers.GetRiskUsageByGridcode)
	r.GET("/api/risks/gridcode", controllers.GetRisksByGridcode)
	r.GET("/api/usages/gridcode", controllers.GetUsagesByGridcode)

	return r
}
