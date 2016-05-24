package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var g_echo string

func sayhelloName(w http.ResponseWriter, r *http.Request) {

	// 解析参数, 默认是不会解析的
	r.ParseForm()

	// 这个写入到w的信息是输出到客户端的
	fmt.Fprintf(w, g_echo)
}

func usage() {
	fmt.Println("========================")
	fmt.Println("src port [args]")
}

func main() {

	var port string

	g_echo = "server feedback:"

	if len(os.Args) < 2 {
		usage()
		os.Exit(1)
	}
	port = os.Args[1]

	for i := 2; i < len(os.Args); i++ {
		g_echo = g_echo + " " + os.Args[i]
	}

	// 设置访问的路由
	http.HandleFunc("/", sayhelloName)

	// 设置监听的端口
	fmt.Println("listening at port ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
