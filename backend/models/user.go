package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"yongdeng-ecology-visualization/config"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(db *config.DBConfig, user *User) error {
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)

	result := db.DB.Create(user)
	return result.Error
}

// GetUserByUsername 根据用户名获取用户
func GetUserByUsername(db *config.DBConfig, username string) (User, error) {
	var user User
	result := db.DB.Where("username = ?", username).First(&user)
	return user, result.Error
}

// 检查输入密码与存储的密码是否匹配
func (user *User) VerifyPassword(inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(inputPassword))
	return err == nil
}
