package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Printf("HTTP Server start failed, err:%v\n", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}
	// 3. 渲染模板
	err = t.Execute(w, "小王子")
	if err != nil {
		fmt.Printf("render template failed, err:%v\n", err)
		return
	}
}
