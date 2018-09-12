package main

import (
	"crypto/md5"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认不解析
	fmt.Println(r.Form) //输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key: ", k)
		fmt.Println("val: ", strings.Join(v, "v"))
	}
	fmt.Fprintf(w, "%s", "hello astaxie!") //这个写入到w的是输出到客户端的
}

// 处理/login逻辑
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) //获取请求方法
	if r.Method == "GET" {
		cruTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cruTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("http_web_login.gtpl")
		t.Execute(w, token)
	} else {
		//请求的是登录数据，那么执行登录的逻辑判断
		r.ParseForm()
		token := r.Form.Get("token")
		if token != "" {
			//验证token的合法性
		} else {
			//不存在token报错
		}

		fmt.Println("usename length:", len(r.Form["username"][10]))
		fmt.Println("usename: ", template.HTMLEscapeString(r.Form.Get("username"))) //输出到服务器端
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username"))) //输出到客户端
	}
}

// 处理/upload逻辑
func upload(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method: ", r.Method) //获取请求方法
	if r.Method == "GET" {
		cruTime := time.Now().Unix()
		h := md5.New()
		io.WriteString(h, strconv.FormatInt(cruTime, 10))
		token := fmt.Sprintf("%x", h.Sum(nil))

		t, _ := template.ParseFiles("http_web_upload.gtpl")
		t.Execute(w, token)
	} else {
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("uploadfile")

		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()

		fmt.Fprintf(w, "%v", handler.Header)
		f, err := os.OpenFile("./trans/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		io.Copy(f, file)
	}
}

func main() {
	http.HandleFunc("/", sayHelloName) //设置访问路由
	http.HandleFunc("/login", login)   //设置访问路由
	http.HandleFunc("/upload", upload) //设置访问路由

	err := http.ListenAndServe(":9090", nil) //监听端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
