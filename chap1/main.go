package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析模板
	// 请勿刻舟求剑,用绝对地址
	t, err := template.ParseFiles("/Users/luliang/GoLand/gin_practice/chap1/hello.tmpl")
	if err != nil {
		fmt.Println("http server failed:%V", err)
		return
	}
	// 渲染模板
	err = t.Execute(w, "小王子!")
	if err != nil {
		fmt.Println("http server failed:%V", err)
		return
	}

}
func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("http server failed:%V", err)
		return
	}

}
