package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", sayHelloName)
	http.HandleFunc("/login", login)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}
}

func sayHelloName(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form) //map类型
	fmt.Println("path", r.URL.Path)
	fmt.Println("schema", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprint(w, "hello mr.su")
}

func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("login.gtpl")
		t.Execute(w, nil)
	} else {
		fmt.Println(r.Form)                               // []
		r.ParseForm()                                     // 此时r.Form为url.Values类型
		fmt.Println("username:", r.Form["username"])      // 得到的是数组
		fmt.Println("username:", r.Form.Get("username"))  // url.Values类型的方法 所以需要先ParseForm
		fmt.Println("username:", r.FormValue("username")) // 自动ParseForm
		fmt.Println("password:", r.Form["password"])

		// 防xss测试
		fmt.Println("username:", template.HTMLEscapeString(r.Form.Get("username")))
		fmt.Println("password:", template.HTMLEscapeString(r.Form.Get("password")))
		template.HTMLEscape(w, []byte(r.Form.Get("username")))
		// fmt.Fprint(w, r.Form.Get("username"))
	}
}
