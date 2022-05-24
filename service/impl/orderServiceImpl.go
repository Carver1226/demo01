package impl

import (
	"demo01/dao"
	"demo01/model"
)

type OrderServiceImpl struct {
	orderDao dao.IDao
}

func NewOrderServiceImpl(orderDao dao.IDao) *OrderServiceImpl {
	return &OrderServiceImpl{orderDao: orderDao}
}

func (osi *OrderServiceImpl)CreateOrder(orderNo, userName string, amount float64, status, fileUrl string) (err error) {
	order := model.NewDemoOrder(orderNo, userName, amount, status, fileUrl)
	if err = osi.orderDao.CreateOrder(order); err != nil {
		return err
	}
	return nil
}

func (osi *OrderServiceImpl)GetAllOrder() (orderList []*model.DemoOrder, err error) {
	if orderList, err = osi.orderDao.GetAllOrder(); err != nil {
		return nil, err
	}
	return orderList, nil
}

func (osi *OrderServiceImpl)GetOrderByOrderNo(orderNo string) (order *model.DemoOrder, err error) {
	if order, err = osi.orderDao.GetOrderByOrderNo(orderNo); err != nil {
		return nil, err
	}
	return order, nil
}


func (osi *OrderServiceImpl)UpdatesOrder(orderNo string, amount float64, status string, fileUrl string) (err error){
	order, err := osi.orderDao.GetOrderByOrderNo(orderNo)
	if err != nil {
		return err
	}
	updateOrder := model.NewDemoOrder(orderNo, "", amount, status, fileUrl)
	if err = osi.orderDao.UpdatesOrder(order, updateOrder); err != nil {
		return err
	}
	return nil
}

func (osi *OrderServiceImpl)DeleteOrderByOrderNo(orderNo string) (err error) {
	order := model.NewDemoOrder(orderNo,"",0,"","")
	if err = osi.orderDao.DeleteOrderByOrderNo(order); err != nil {
		return err
	}
	return nil
}
