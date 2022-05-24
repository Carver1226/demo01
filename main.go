package main

import (
	"demo01/util"
	"demo01/router"
)

func main() {
	util.InitMySql()
	defer util.Close()
	router.InitRouter()
}