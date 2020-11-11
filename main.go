package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe(":3389", nil)
	if err != nil {
		fmt.Println("HTTP server failed,err:", err)
		return
	}
}

func sayHello(w http.ResponseWriter, r *http.Request) {
	// 解析指定文件生成模板对象
	tmpl, err := template.ParseFiles("./hello.tmpl")
	if err != nil {
		fmt.Println("create template failed, err:", err)
		return
	}
	// 利用给定数据渲染模板，并将结果写入w
	errTmpl := tmpl.Execute(w, "dough")
	if errTmpl != nil {
		fmt.Println("errTmpl failed, err:", err)
		return
	}
}
