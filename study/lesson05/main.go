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
		fmt.Printf("The http server start failed, err:%v\n", err)
		return
	}
}

type User struct {
	Name string
	Gender string
	Age int
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 1. 定义模板

	// 2. 解析模板
	t, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Printf("Parse template failed, err:%v\n", err)
		return
	}

	// 3. 渲染模板
	u1 := User{
		Name: "小王子",
		Gender: "男",
		Age: 18,
	}

	m1 := map[string] interface{}{
		"name": "小王子",
		"gender": "男",
		"age": 18,
	}
	err = t.Execute(w, map[string]interface{}{
		"u1": u1,
		"m1": m1,
	})
	if err != nil {
		fmt.Printf("Render template failed, err:%v\n", err)
		return
	}
}
