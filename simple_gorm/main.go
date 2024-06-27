package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	ID           uint
	Name         string
	Email        *string
	Age          uint8
	Birthday     *time.Time
	MemberNumber sql.NullString // 使用 sql.NullString 处理可空字符串
	ActivatedAt  sql.NullTime   // 使用 sql.NullTime 处理可空时间字段
	CreatedAt    time.Time      // 创建时间（由 GORM 自动管理）
	UpdatedAt    time.Time
}

func main() {
	user := User{Name: "ellan", Age: 18}
	dsn := "host=localhost user=postgres password='Hhu123456&' dbname=gin_web port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dsn,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("无法连接数据库: %v", err)
	}

	fmt.Println("连接成功")
	fmt.Println(db)
	// 手动创建表
	err = db.Exec(`CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255),
		email VARCHAR(255),
		age INT,
		birthday DATE,
		member_number VARCHAR(50),
		activated_at TIMESTAMP,
		created_at TIMESTAMP DEFAULT NOW(),
		updated_at TIMESTAMP DEFAULT NOW()
	)`).Error
	if err != nil {
		panic("无法创建表")
	}

	result := db.Create(&user)
	if result.Error != nil {
		log.Fatalf("创建用户失败: %v", result.Error)
	}

	fmt.Println("用户创建成功")
	fmt.Println(result.RowsAffected)
}
