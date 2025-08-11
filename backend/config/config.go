package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type DBConfig struct {
	DB *gorm.DB
}

func LoadConfig() (*DBConfig, error) {
	// 硬编码数据库连接信息
	host := "localhost"
	user := "postgres"
	password := "12345"
	dbname := "yongdeng"
	port := "5432"

	// 构建数据库连接字符串
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable options=-csearch_path=public,postgis",
		host, user, password, dbname, port,
	)

	// 注册序列化器：在GORM配置中通过Serializers字段注册
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	log.Println("Database connection established successfully.")

	return &DBConfig{DB: db}, nil
}
