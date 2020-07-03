package main

import (
	"github.com/danielkolbe/connectfour/app/logger"
	"net/http"
	"github.com/danielkolbe/connectfour/api"
)

func main() {
	http.Handle("/", api.NewRouter())
	logger.Logger.Info("Starting to serve on port 8080.")
	err := http.ListenAndServe(":8080", nil)
	if(nil != err) {
		logger.Logger.Fatal(err)
	}
}
