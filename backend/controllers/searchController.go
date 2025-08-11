package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"yongdeng-ecology-visualization/config"
	"yongdeng-ecology-visualization/models"
)

// GetRiskUsageByGridcode 根据 Gridcode 获取风险数据
func GetRiskUsageByGridcode(c *gin.Context) {
	// 从上下文获取数据库连接
	dbConfig := c.MustGet("db").(*config.DBConfig)
	db := dbConfig.DB

	// 获取并验证查询参数
	gridcodeStr := c.Query("gridcode")
	if gridcodeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "gridcode参数不能为空",
		})
		return
	}

	gridcode, err := strconv.Atoi(gridcodeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的gridcode格式",
			"error":   err.Error(),
		})
		return
	}

	// 调用模型层查询
	riskUsages, err := models.GetRiskUsageByGridcode(db, gridcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询数据失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回标准化的响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"count":   len(riskUsages), // 增加数据条数
		"data":    riskUsages,
	})
}

// GetRisksByGridcode 根据 Gridcode 获取风险数据
func GetRisksByGridcode(c *gin.Context) {
	// 从上下文获取数据库连接
	dbConfig := c.MustGet("db").(*config.DBConfig)
	db := dbConfig.DB

	// 获取并验证查询参数
	gridcodeStr := c.Query("gridcode")
	if gridcodeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "gridcode参数不能为空",
		})
		return
	}

	gridcode, err := strconv.Atoi(gridcodeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的gridcode格式",
			"error":   err.Error(),
		})
		return
	}

	// 调用模型层查询
	riskUsages, err := models.GetRisksByGridcode(db, gridcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询数据失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回标准化的响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"count":   len(riskUsages), // 增加数据条数
		"data":    riskUsages,
	})
}

// GetUsagesByGridcode 根据 Gridcode 获取风险数据
func GetUsagesByGridcode(c *gin.Context) {
	// 从上下文获取数据库连接
	dbConfig := c.MustGet("db").(*config.DBConfig)
	db := dbConfig.DB

	// 获取并验证查询参数
	gridcodeStr := c.Query("gridcode")
	if gridcodeStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "gridcode参数不能为空",
		})
		return
	}

	gridcode, err := strconv.Atoi(gridcodeStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    400,
			"message": "无效的gridcode格式",
			"error":   err.Error(),
		})
		return
	}

	// 调用模型层查询
	riskUsages, err := models.GetUsagesByGridcode(db, gridcode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    500,
			"message": "查询数据失败",
			"error":   err.Error(),
		})
		return
	}

	// 返回标准化的响应
	c.JSON(http.StatusOK, gin.H{
		"code":    200,
		"message": "查询成功",
		"count":   len(riskUsages), // 增加数据条数
		"data":    riskUsages,
	})
}
