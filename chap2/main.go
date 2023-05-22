package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type User struct {
	Name   string
	Gander string // 首字母是否大小写,作为是否对外暴露的标识
	Age    int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	u1 := User{
		Name:   "小王子",
		Gander: "男",
		Age:    19,
	}
	// 解析模板
	t, err := template.ParseFiles("/Users/luliang/GoLand/gin_practice/chap2/test.tmpl")
	if err != nil {
		fmt.Println("ParseFiles failed:%V", err)
		return
	}
	err = t.Execute(w, u1)
	if err != nil {
		return
	}
	// 渲染模板

}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server failed:%V", err)
		return
	}

}
