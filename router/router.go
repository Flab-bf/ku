package router

import (
	"first/api"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func NewRouter() *server.Hertz {
	h := server.New()
	h.POST("/register", api.Register)
	h.POST("/login", api.Login)
	h.POST("/comment/write", api.Writer)
	h.POST("/comment/delete", api.Delete)
	h.GET("/comment/readall", api.GetAll)
	return h
}
