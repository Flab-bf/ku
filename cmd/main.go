package main

import (
	"first/api"
	"first/utils"
)

func main() {
	utils.ConnectDB()
	r := api.NewRouter()
	r.Spin()
}
