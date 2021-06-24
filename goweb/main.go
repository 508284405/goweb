package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	//获取访问路径
	fmt.Println(r.URL.Path)
	//获取访问?之后的字符串
	fmt.Println(r.URL.RawQuery)
	//获取请求头信息,map
	fmt.Println(r.Header)
	//获取请求头内具体的信息
	fmt.Println(r.Header.Get("User-Agent"))
	//获取请求体
	//获取请求体长度
	//length := r.ContentLength
	//defer r.Body.Close()
	//body := make([]byte, length)
	////此处不能有接收异常信息
	//a, _ := r.Body.Read(body)
	////请求体的长度
	//fmt.Println("body.length", a)
	//fmt.Println("body : ", string(body))

	//FORM 字段
	r.ParseForm()
	form := r.PostForm
	fmt.Println("name : ", form)
}

func main() {
	http.HandleFunc("/hello", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		return
	}
}
