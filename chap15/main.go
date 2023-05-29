package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
// 将会在数据库中帮我们生成一个表单
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func main() {
	// 连接数据库
	// 默认端口为 3306
	db, err := gorm.Open("mysql", "root:791975457@qq.com@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	// 最后关闭数据库
	defer func(db *gorm.DB) {
		err := db.Close()
		if err != nil {
		}
	}(db)

	// 自动迁移
	// 自动迁移,本质上就是把结构体映射成 table
	db.AutoMigrate(&UserInfo{})

	//u1 := UserInfo{1, "七米", "男", "篮球"}
	//u2 := UserInfo{2, "沙河娜扎", "女", "足球"}

	// 创建记录
	//db.Create(&u1)
	//db.Create(&u2)
	// 更新
	var u UserInfo
	db.First(&u)
	fmt.Printf("u:%v\n", u)
	db.Model(&u).Update("hobby", "双色球")
	db.Delete(&u)

	// 查询
	//var u = new(UserInfo)
	//db.First(u)
	//fmt.Printf("%#v\n", u)
	//
	//var uu UserInfo
	//db.Find(&uu, "hobby=?", "足球")
	//fmt.Printf("%#v\n", uu)
	//
	//// 更新
	//db.Model(&u).Update("hobby", "双色球")
	//// 删除
	//db.Delete(&u)
}
