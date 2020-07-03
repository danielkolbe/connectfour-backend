package main

import (
	"github.com/danielkolbe/boardgames/app/logger"
	"net/http"
	"github.com/danielkolbe/boardgames/api"
)

func main() {
	http.Handle("/", api.NewRouter())
	logger.Logger.Info("Starting to serve on port 8080.")
	err := http.ListenAndServe(":8080", nil)
	if(nil != err) {
		logger.Logger.Fatal(err)
	}
}
