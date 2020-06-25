package main

import (
	"fmt"
	"net/http"
	"github.com/danielkolbe/connectfour/api"
)

func main() {
	http.HandleFunc("/turn", handlers.Turn);
	http.Handle("/favicon.ico", http.NotFoundHandler())
	err := http.ListenAndServe(":8080", nil)
	if(nil != err) {
		fmt.Print(err)
	}
}
