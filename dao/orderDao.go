package dao

import (
	"demo01/model"
	"github.com/jinzhu/gorm"
)

type OrderDao struct {
	db *gorm.DB
}

func NewOrderDao(db *gorm.DB) *OrderDao {
	return &OrderDao{db: db}
}

func (orderDao *OrderDao)CreateOrder(order *model.DemoOrder) (err error){
	if err = orderDao.db.Create(&order).Error; err != nil {
		return err
	}
	return nil
}

func (orderDao *OrderDao)GetAllOrder() (orderList []*model.DemoOrder, err error){
	if err = orderDao.db.Find(&orderList).Error; err != nil {
		return nil, err
	}
	return orderList, nil
}

func (orderDao *OrderDao)GetOrderByOrderNo(orderNo string) (order *model.DemoOrder, err error) {
	order = &model.DemoOrder{}
	if err := orderDao.db.First(&order, "order_no=?", orderNo).Error; err != nil {
		return nil, err
	}
	return order, nil
}

func (orderDao *OrderDao)UpdateOrder(order *model.DemoOrder) (err error){
	if err = orderDao.db.Save(&order).Error; err != nil {
		return err
	}
	return nil
}

func (orderDao *OrderDao)UpdatesOrder(order *model.DemoOrder,updateOrder *model.DemoOrder) (err error){
	if err = orderDao.db.Model(&order).Updates(updateOrder).Error; err != nil{
		return err
	}
	return nil
}

func (orderDao *OrderDao)DeleteOrderByOrderNo(order *model.DemoOrder) (err error){
	if err = orderDao.db.Delete(&order).Error; err != nil {
		return err
	}
	return nil
}