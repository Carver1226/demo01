package model

type DemoOrder struct {
	OrderNo string `gorm:"primary_key"`
	UserName string
	Amount float64
	Status string
	FileUrl string
}

func NewDemoOrder(orderNo string, userName string, amount float64, status string, fileUrl string) *DemoOrder {
	return &DemoOrder{OrderNo: orderNo, UserName: userName, Amount: amount, Status: status, FileUrl: fileUrl}
}