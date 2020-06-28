package main

import (
	"fmt"
	"net/http"
	"github.com/danielkolbe/connectfour/api"
)

func main() {
	http.Handle("/", api.NewRouter())
	err := http.ListenAndServe(":8080", nil)
	if(nil != err) {
		fmt.Print(err)
	}
}
