/*
	说明，参数1表示端口号，如果没有则默认8888
*/

package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	http.Handle("/test/", http.FileServer(http.Dir("template")))
	http.Handle("/css/", http.FileServer(http.Dir("template")))
	http.Handle("/js/", http.FileServer(http.Dir("template")))

	http.HandleFunc("/login/", loginHandler)
	http.HandleFunc("/", NotFoundHandler)

	var port string
	if len(os.Args) > 1 {
		port = os.Args[1]
	} else {
		port = "8888"
	}

	fmt.Println("http opened at", port)

	addr := ":" + port
	http.ListenAndServe(addr, nil)

}
