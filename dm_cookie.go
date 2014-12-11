package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", testCookie) // http.HandleFunc 第二个参数 期待一个以 http.ResponseWriter, *http.Request 类型为参数的函数
	http.HandleFunc("/g", getCookie)
	err := http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("ListenAndServer: ", err)
	}

	// fmt.Println('s')
}

func testCookie(w http.ResponseWriter, r *http.Request) {
	// 设置cookie
	expiration := time.Now() // Time
	expiration = expiration.AddDate(1, 0, 0)
	cookie := http.Cookie{Name: "username", Value: "mario", Expires: expiration}
	http.SetCookie(w, &cookie)
}

func getCookie(w http.ResponseWriter, r *http.Request) {
	// 获得cookie
	cookie, _ := r.Cookie("username")
	fmt.Println(cookie)

	// 获取所有cookie
	for _, cookie := range r.Cookies() {
		fmt.Println(cookie)
	}

	fmt.Fprint(w, cookie)
}

/*
1.cookie 对象
type Cookie struct {
    Name       string
    Value      string
    Path       string
    Domain     string
    Expires    time.Time
    RawExpires string

// MaxAge=0 means no 'Max-Age' attribute specified.
// MaxAge<0 means delete cookie now, equivalently 'Max-Age: 0'
// MaxAge>0 means Max-Age attribute present and given in seconds
    MaxAge   int
    Secure   bool
    HttpOnly bool
    Raw      string
    Unparsed []string // Raw text of unparsed attribute-value pairs
}
*/
