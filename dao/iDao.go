package dao

import "demo01/model"

type IDao interface {
	CreateOrder(order *model.DemoOrder) (err error)
	GetAllOrder() (orderList []*model.DemoOrder, err error)
	GetOrderByOrderNo(orderNo string) (order *model.DemoOrder, err error)
	UpdateOrder(order *model.DemoOrder) (err error)
	UpdatesOrder(order *model.DemoOrder,updateOrder *model.DemoOrder) (err error)
	DeleteOrderByOrderNo(order *model.DemoOrder) (err error)
}
