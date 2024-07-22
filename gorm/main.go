/*
 * @Author:
 * @Date: 2024-07-22 19:05:23
 * @LastEditors:
 * @LastEditTime: 2024-07-22 19:42:30
 * @Description:
 */
package main

import (
	"database/sql"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// type User struct {
// 	ID           uint                 //可直接使用
// 	Name         string               //可直接使用
// 	Email        *string              // 表示可空字段
// 	Age          uint8                //可直接使用
// 	Burthday     *time.Time           // 表示可空字段
// 	MemberNumber sql.NullStringchan   // 用于具有更多控制的可控字段
// 	ActiveatedAt sql.NullTimechanchan // 用于具有更多控制的可控字段
// 	CreatedAt    time.Time            //创建时间Gorm自动管理
// 	UpdatedAt    time.Time            //最后一次更新时间 gorm自动管理
// }

// tips：约定
// 主键：gorm使用一个名为ID的字段作为每个模型默认的主键
// 表名：默认情况下 gorm将结构体名称转换为snake_case 并未表名加上复数形式 例如 以恶个User结构体在数据库中的表名变为Users
// 列名：gorm自动将结构体字段名称转换为snake_case作为数据库中的列名
// 时间戳字段： gorm使用字段CreatedAt和UpdatedAt来自动跟踪记录的创建和更新时间

func main() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=Trye&loc=Local"

	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库打开失败")
	}
	gorm, err := gorm.Open(mysql.New(mysql.Config{
		Con: sqlDB,
	}), &gorm.Config{})

	fmt.Printf("数据库", gorm)

}
