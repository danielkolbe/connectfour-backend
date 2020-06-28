package api

import (
	"github.com/gorilla/mux"
	"net/http"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/turn", Turn).Methods("POST").Schemes("http")
	r.Handle("/favicon.ico", http.NotFoundHandler())
	return r
}