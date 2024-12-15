package main

import (
	"first/dao"
	"first/router"
)

func main() {
	dao.ConnectDB()
	r := router.NewRouter()
	r.Spin()
}
