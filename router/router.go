package router

import (
	"demo01/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() {
	r := gin.Default()

	handler.OrderApi(r)

	r.Run(":8080")
}