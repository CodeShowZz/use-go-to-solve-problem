package main

import (
	"fmt"
	"log"
	"net/http"
)

//写一个简单的web应用,模拟登陆
func main() {
 	http.HandleFunc("/login",login)
 	err := http.ListenAndServe(":8080",nil)
 	if err !=nil {
 		log.Fatal("ListenAndServe:",err)
	}
}

func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	userName := r.Form.Get("userName")
	password := r.Form.Get("password")
	//模拟数据库操作
	if userName == "root"  && password == "123456" {
		fmt.Fprintf(w,"登陆成功")
	} else {
		fmt.Fprintf(w,"请传入正确的用户名和密码")
	}
}
