package main

import (
	"net/http"
	"os"
)

func NotFoundHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(2)

}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	os.Exit(1)
}
