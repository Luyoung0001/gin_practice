package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	db, err := gorm.Open("mysql", "root:791975457@qq.com@tcp(127.0.0.1:3306)/db3?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}

	// 最后关闭数据库
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)
	db.AutoMigrate(&User{})
	//db.Create(&User{Name: "Lucy", Age: 19})
	//db.Create(&User{Name: "LuL", Age: 29})

	//进行查询

	// 根据主键查询第一条数据
	var user1 User
	db.First(&user1)
	fmt.Println(user1)

	// 随机获取一条记录
	var user2 User
	db.Take(&user2)
	fmt.Println(user2)

	// 根据主键查询最后一条数据
	var user3 User
	db.Last(&user3)
	fmt.Println(user3)
	// 第二个参数
	var user4 User
	db.First(&user4, 2)
	fmt.Println(user4)

}
