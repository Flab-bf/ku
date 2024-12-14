package api

import "github.com/cloudwego/hertz/pkg/app/server"

func NewRouter() *server.Hertz {
	h := server.New()
	h.POST("/register", Register)
	h.POST("/login", Login)
	h.POST("/comment/write", Writer)
	h.POST("/comment/delete", Delete)
	h.GET("/comment/readall", GetAll)
	return h
}
