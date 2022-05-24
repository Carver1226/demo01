package impl

import (
	"demo01/dao"
	"demo01/model"
	"errors"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"testing"
)

type OrderServiceTestSuite struct {
	 suite.Suite
	 orderService *OrderServiceImpl
	 m *dao.MockIDao
}

func TestOrderServiceTestSuite(t *testing.T) {
	suite.Run(t, new(OrderServiceTestSuite))
}

func (o *OrderServiceTestSuite)BeforeTest(suiteName, testName string) {
	ctrl := gomock.NewController(o.T())
	ctrl.Finish()
	m := dao.NewMockIDao(ctrl)
	m.EXPECT().CreateOrder(model.NewDemoOrder("1","carver",999,"true","blog.carvermia.top")).Return(nil)
	m.EXPECT().CreateOrder(model.NewDemoOrder("1","carver",999,"true","blog.carvermia.top")).Return(errors.New("OrderNo已存在"))
	m.EXPECT().GetAllOrder().Return([]*model.DemoOrder{
		{
			OrderNo: "1",
			UserName: "carver",
			Amount: 100,
			Status: "true",
			FileUrl: "blog.carvermia.top",
		},
		{
			OrderNo: "2",
			UserName: "carver",
			Amount: 100,
			Status: "true",
			FileUrl: "blog.carvermia.top",
		},
	},nil)
	m.EXPECT().GetAllOrder().Return(nil, errors.New("error"))
	m.EXPECT().GetOrderByOrderNo("1").Return(model.NewDemoOrder("1","carver",999,"true","blog.carvermia.top"), nil)
	m.EXPECT().GetOrderByOrderNo("2").Return(nil, errors.New("error"))
	m.EXPECT().GetOrderByOrderNo(gomock.Any()).Return(&model.DemoOrder{}, nil)
	m.EXPECT().UpdatesOrder(gomock.Any(), gomock.Any()).Return(nil)
	m.EXPECT().GetOrderByOrderNo(gomock.Any()).Return(nil, errors.New("error"))
	m.EXPECT().GetOrderByOrderNo(gomock.Any()).Return(&model.DemoOrder{}, nil)
	m.EXPECT().UpdatesOrder(gomock.Any(), gomock.Any()).Return(errors.New("error"))
	m.EXPECT().DeleteOrderByOrderNo(gomock.Any()).Return(nil)
	m.EXPECT().DeleteOrderByOrderNo(gomock.Any()).Return(errors.New("error"))

	o.orderService = NewOrderServiceImpl(m)
}

func (o *OrderServiceTestSuite)TestOrderServiceImpl_CreateOrder() {
	err := o.orderService.CreateOrder("1","carver",999,"true","blog.carvermia.top")
	o.NoError(err)
	err = o.orderService.CreateOrder("1","carver",999,"true","blog.carvermia.top")
	o.Error(err)
}

func (o *OrderServiceTestSuite)TestOrderServiceImpl_GetAllOrder() {
	orderList, err := o.orderService.GetAllOrder()
	o.Equal(orderList, []*model.DemoOrder{
		{
			OrderNo: "1",
			UserName: "carver",
			Amount: 100,
			Status: "true",
			FileUrl: "blog.carvermia.top",
		},
		{
			OrderNo: "2",
			UserName: "carver",
			Amount: 100,
			Status: "true",
			FileUrl: "blog.carvermia.top",
		},
	})
	o.NoError(err)
	_, err = o.orderService.GetAllOrder()
	o.Error(err)
}

func (o *OrderServiceTestSuite)TestOrderServiceImpl_GetOrderByOrderNo() {
	order, err := o.orderService.GetOrderByOrderNo("1")
	o.NoError(err)
	o.Equal(order.UserName, "carver")
	_, err = o.orderService.GetOrderByOrderNo("2")
	o.Error(err)
}

func (o *OrderServiceTestSuite)TestOrderServiceImpl_UpdatesOrder() {
	err := o.orderService.UpdatesOrder("1",999,"true","blog.carvermia.top")
	o.NoError(err)
	err = o.orderService.UpdatesOrder("0000",0,"","")
	o.Error(err)
	err = o.orderService.UpdatesOrder("1",0,"","")
	o.Error(err)
}

func (o *OrderServiceTestSuite)TestOrderServiceImpl_DeleteOrderByOrderNo() {
	err := o.orderService.DeleteOrderByOrderNo("1")
	o.NoError(err)
	err = o.orderService.DeleteOrderByOrderNo("0")
	o.Error(err)
}