package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintln(w, "Hello, Golang!")
}

func main() {
	http.HandleFunc("/hello", sayHello)
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Printf("http server failed:%v\n", err)
		return
	}

}
