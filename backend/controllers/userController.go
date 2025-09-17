package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
	"yongdeng-ecology-visualization/backend/config"
	"yongdeng-ecology-visualization/backend/models"
)

// JWT 密钥，用于签名和验证
var jwtKey = []byte("your_secret_key")

// 登录请求结构体
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func RegisterUser(c *gin.Context) {
	db := c.MustGet("db").(*config.DBConfig)
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	if err := models.CreateUser(db, &user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":  "Registration successful",
		"username": user.Username,
		"email":    user.Email,
	})
}

func Login(c *gin.Context) {
	var request LoginRequest
	db := c.MustGet("db").(*config.DBConfig)

	// 绑定请求数据
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 查找用户
	user, err := models.GetUserByUsername(db, request.Username)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 验证密码
	if !user.VerifyPassword(request.Password) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	// 生成 JWT
	token, err := generateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	// 返回 JWT
	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}

// 生成 JWT
func generateJWT(user models.User) (string, error) {
	claims := &jwt.StandardClaims{
		Issuer:    fmt.Sprintf("%d", user.ID),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(), // 设置 token 过期时间为 24 小时
	}

	// 创建 JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey)
	if err != nil {
		log.Printf("Error signing token: %v", err)
		return "", err
	}

	return signedToken, nil
}
