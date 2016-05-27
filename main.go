package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func usage(port string) {
	fmt.Println("========================")
	fmt.Println("./app [port]")
	fmt.Println("default port is ", port)
	fmt.Println("========================")
}

func show(w http.ResponseWriter, r *http.Request) {
	fmt.Println("*************")
	// 解析参数, 默认是不会解析的
	r.ParseForm()

	// 显示所有输出目录
	var ret string
	ret = "<html><body>"
	err := filepath.Walk("./html/profile/", func(path string, f os.FileInfo, err error) error {
		if f == nil {
			return err
		}
		if f.IsDir() {
			return nil
		}

		fmt.Println(path)
		ret += fmt.Sprintf("<a href='%s'>sdfsdf</a>", path)
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
	ret += "</body></html>"

	// 这个写入到w的信息是输出到客户端的
	fmt.Fprintf(w, ret)
}

func main() {

	port := "30300"
	usage(port)

	if len(os.Args) > 1 {
		port = os.Args[1]
	}

	// 设置访问的路由
	http.Handle("/", http.FileServer(http.Dir("html")))

	// 设置监听的端口
	fmt.Println("golang scatter http server working at port: ", port)
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
