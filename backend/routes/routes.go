package routes

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"yongdeng-ecology-visualization/backend/config"
	"yongdeng-ecology-visualization/backend/controllers"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 添加 CORS 跨域中间件
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
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

	// 载入数据库配置
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

	// 提供静态文件服务（页面和静态资源）
	r.Static("/static", "./frontend/static") // 映射静态文件夹
	r.GET("/", func(c *gin.Context) {
		// 返回 page.html 文件
		c.File("./frontend/page.html")
	})
	//加载边界
	r.GET("/yongdeng_boundary.json", func(c *gin.Context) {
		c.File("./frontend/static/yongdeng_boundary.json") // 提供文件
	})

	// 提供 PDF 文件下载（usage.pdf）
	r.GET("/usage.pdf", func(c *gin.Context) {
		// 检查文件是否存在
		if _, err := os.Stat("./frontend/usage.pdf"); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "usage.pdf not found"})
			return
		}
		// 提供文件下载
		c.File("./frontend/usage.pdf")
	})

	// 提供 PDF 文件下载（risk.pdf）
	r.GET("/risk.pdf", func(c *gin.Context) {
		// 检查文件是否存在
		if _, err := os.Stat("./frontend/risk.pdf"); os.IsNotExist(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "risk.pdf not found"})
			return
		}
		// 提供文件下载
		c.File("./frontend/risk.pdf")
	})

	// 提供用户注册和登录接口
	r.POST("/api/auth/register", controllers.RegisterUser)
	r.POST("/api/auth/login", controllers.Login)

	// 根据 Gridcode 获取数据接口
	r.GET("/api/riskusage/gridcode", controllers.GetRiskUsageByGridcode)
	r.GET("/api/risks/gridcode", controllers.GetRisksByGridcode)
	r.GET("/api/usages/gridcode", controllers.GetUsagesByGridcode)

	return r
}
