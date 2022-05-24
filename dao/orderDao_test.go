package dao

import (
	"demo01/model"
	"demo01/util"
	"github.com/jinzhu/gorm"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OrderDaoTestSuite struct {
	suite.Suite
	orderDao *OrderDao
	db *gorm.DB
}

//测试初始化
func (o *OrderDaoTestSuite) SetupTest() {
	db, _ := util.InitMySql()
	orderDao := NewOrderDao(db)
	o.orderDao = orderDao
	o.db = db
	db.Delete(&model.DemoOrder{})
}

func TestOrderDaoTestSuite(t *testing.T) {
	suite.Run(t, new(OrderDaoTestSuite))
}

func (o *OrderDaoTestSuite)TestOrderDao_CreateOrder() {
	err := o.orderDao.CreateOrder(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver",
		Amount:   100,
		Status:   "true",
	})
	o.NoError(err)
	var order model.DemoOrder
	err = o.db.Model(&model.DemoOrder{}).Where("order_no=?", "1").First(&order).Error
	o.NoError(err)
	o.Equal("carver", order.UserName)
	o.Equal(100.0, order.Amount)
	o.Equal("true", order.Status)
	o.Equal("", "")
	err = o.orderDao.CreateOrder(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver",
		Amount:   100,
		Status:   "true",
	})
	o.Error(err)
}

func (o *OrderDaoTestSuite)TestOrderDao_GetAllOrder() {
	o.db.Create(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver01",
		Amount:   100,
		Status:   "true",
	})
	o.db.Create(&model.DemoOrder{
		OrderNo:  "2",
		UserName: "carver02",
		Amount:   100,
		Status:   "true",
	})
	orderList, err := o.orderDao.GetAllOrder()
	o.NoError(err)
	o.Equal("carver01", orderList[0].UserName)
	o.Equal("carver02", orderList[1].UserName)
}

func (o *OrderDaoTestSuite)TestOrderDao_GetOrderByOrderNo() {
	err := o.db.Create(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver01",
		Amount:   100,
		Status:   "true",
	}).Error
	o.NoError(err)
	err = o.db.Create(&model.DemoOrder{
		OrderNo:  "2",
		UserName: "carver02",
		Amount:   100,
		Status:   "true",
	}).Error
	o.NoError(err)
	order01, err := o.orderDao.GetOrderByOrderNo("1")
	o.NoError(err)
	order02, err := o.orderDao.GetOrderByOrderNo("2")
	o.NoError(err)
	_, err = o.orderDao.GetOrderByOrderNo("3")
	o.Error(err)
	o.Equal("carver01", order01.UserName)
	o.Equal("carver02", order02.UserName)
}

func (o *OrderDaoTestSuite)TestOrderDao_UpdateOrder() {
	err := o.db.Create(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver01",
		Amount:   100,
		Status:   "true",
	}).Error
	o.NoError(err)
	err = o.orderDao.UpdateOrder(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver02",
		Amount:   120,
		Status:   "false",
	})
	o.NoError(err)
	var updateOrder model.DemoOrder
	err = o.db.Model(&model.DemoOrder{}).Where("order_no=?", "1").First(&updateOrder).Error

	o.Equal("1", updateOrder.OrderNo)
	o.Equal("carver02", updateOrder.UserName)
	o.Equal(120.0, updateOrder.Amount)
	o.Equal("false", updateOrder.Status)
}

func (o *OrderDaoTestSuite)TestOrderDao_DeleteOrderByOrderNo() {
	err := o.db.Create(&model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver01",
		Amount:   100,
		Status:   "true",
	}).Error
	o.NoError(err)
	var order model.DemoOrder
	err = o.db.Model(&model.DemoOrder{}).Where("order_no=?", "1").First(&order).Error
	o.NoError(err)
	o.orderDao.DeleteOrderByOrderNo(&order)
	err = o.db.Model(&model.DemoOrder{}).Where("order_no=?", "1").First(&order).Error
	o.Error(err)
}

func (o *OrderDaoTestSuite)TestOrderDao_UpdatesOrder() {
	order := &model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver01",
		Amount:   100,
		Status:   "true",
	}
	err := o.orderDao.CreateOrder(order)
	o.NoError(err)
	update := &model.DemoOrder{
		OrderNo:  "1",
		UserName: "carver02",
		Amount:   200,
		Status:   "false",
	}
	err = o.orderDao.UpdatesOrder(order, update)
	o.NoError(err)
	result, err := o.orderDao.GetOrderByOrderNo("1")
	o.NoError(err)
	o.Equal(order.OrderNo, result.OrderNo)
	o.Equal(order.UserName, result.UserName)
	o.Equal(update.UserName, result.UserName)
	o.Equal(update.Amount, result.Amount)
	o.Equal(update.Status, result.Status)
	o.Equal(order, result)
}