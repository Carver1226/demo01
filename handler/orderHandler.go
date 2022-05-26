package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"demo01/dao"
	"demo01/service"
	"demo01/service/impl"
	"demo01/util"
)

type OrderHandler struct {
	orderService service.OrderService
}

func NewOrderHandler(orderService service.OrderService) *OrderHandler {
	return &OrderHandler{orderService: orderService}
}

func OrderApi(router *gin.Engine) {
	orderHandler := NewOrderHandler(impl.NewOrderServiceImpl(dao.NewOrderDao(util.InitMySql())))

	orderGroup := router.Group("/order")
	{
		orderGroup.GET("", orderHandler.GetAllOrder)
		orderGroup.GET("/:orderNo", orderHandler.GetOrderByOrderNo)
		orderGroup.POST("", orderHandler.CreateOrder)
		orderGroup.PUT("", orderHandler.UpdateOrderByOrderNo)
		orderGroup.DELETE("/:orderNo", orderHandler.DeleteOrderByOrderNo)
	}
}

func (orderHandler *OrderHandler)CreateOrder(c *gin.Context) {
	orderNo := c.PostForm("orderNo")
	userName := c.PostForm("userName")
	amount,_ := strconv.ParseFloat(c.PostForm("amount"), 64)
	status := c.PostForm("status")
	fileUrl := c.PostForm("fileUrl")
	err := orderHandler.orderService.CreateOrder(orderNo, userName, amount, status, fileUrl)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
}

func (orderHandler *OrderHandler)GetOrderByOrderNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	order, err := orderHandler.orderService.GetOrderByOrderNo(orderNo)
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, order)
}

func (orderHandler *OrderHandler)GetAllOrder(c *gin.Context) {
	orderList, err := orderHandler.orderService.GetAllOrder()
	if err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
	c.JSON(http.StatusOK, orderList)
}

func (orderHandler *OrderHandler)UpdateOrderByOrderNo(c *gin.Context) {
	orderNo := c.PostForm("orderNo")
	amount,_ := strconv.ParseFloat(c.PostForm("amount"), 64)
	status := c.PostForm("status")
	fileUrl := c.PostForm("fileUrl")

	if err := orderHandler.orderService.UpdatesOrder(orderNo, amount, status, fileUrl); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
}

func (orderHandler *OrderHandler)DeleteOrderByOrderNo(c *gin.Context) {
	orderNo := c.Param("orderNo")
	if err := orderHandler.orderService.DeleteOrderByOrderNo(orderNo); err != nil {
		c.JSON(http.StatusBadRequest,gin.H{"error":err.Error()})
		return
	}
}


