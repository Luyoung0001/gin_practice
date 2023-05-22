package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func f(w http.ResponseWriter, r *http.Request) {
	// 定义模板
	k := func(name string) (string, error) {
		return name + "太棒了!", nil
	}
	t := template.New("f.tmpl")
	t.Funcs(template.FuncMap{
		"kua": k,
	})
	// 解析模板
	_, err := t.ParseFiles("/Users/luliang/GoLand/gin_practice/chap3/f.tmpl")
	if err != nil {
		return
	}
	// 渲染模板

	err = t.Execute(w, "小王子")
	if err != nil {
		return
	}

}

func main() {
	http.HandleFunc("/hello", f)
	err := http.ListenAndServe(":9002", nil)
	if err != nil {
		fmt.Println("http server failed:%V", err)
		return
	}

}
