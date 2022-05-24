package service

import (
	"demo01/model"
)

type OrderService interface {
	CreateOrder(orderNo, userName string, amount float64, status, fileUrl string) (err error)
	GetAllOrder() (orderList []*model.DemoOrder, err error)
	GetOrderByOrderNo(orderNo string) (orderList *model.DemoOrder, err error)
	UpdatesOrder(orderNo string, amount float64, status string, fileUrl string) (err error)
	DeleteOrderByOrderNo(orderNo string) (err error)
}