package main

import (
	"fmt"
	"net/http"
	"strings"
	"log"
	"html/template"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()  //解析参数，默认是不会解析的
	fmt.Println(r.Form)  //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!") //这个写入到w的是输出到客户端的
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //获取请求的方法
	if r.Method == "GET" {
		t, _ := template.ParseFiles("src/web/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		//请求的是登陆数据，那么执行登陆的逻辑判断
		r.ParseForm();
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		username := r.Form.Get("username")
		password := r.Form.Get("password")
		if username == "" || password == "" {
			w.Write([]byte("username or password is error"));
		} else {
			w.Write([]byte("login success"));
		}
		// 验证中文
		/*if m, _ := regexp.MatchString("^\\p{Han}+$", username); !m {
			return false
		}*/
		// 验证英文
		/*if m, _ := regexp.MatchString("^[a-zA-Z]+$",username); !m {
			return false
		}*/
		// 验证邮箱
		/*if m, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, r.Form.Get("email")); !m {
			fmt.Println("no")
		}else{
			fmt.Println("yes")
		}*/
		// 验证手机
		/*if m, _ := regexp.MatchString(`^(1[3|4|5|8][0-9]\d{4,8})$`, r.Form.Get("mobile")); !m {
			return false
		}*/

	}
}

func main() {
	http.HandleFunc("/", sayhelloName) //设置访问的路由
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}