package order

import (
	"encoding/json"

	"github.com/silenceper/wechat/v2/minishop/context"
)

const (
	orderGetListURL = "https://api.weixin.qq.com/product/order/get_list?access_token=%s"
	orderGetURL     = "https://api.weixin.qq.com/product/order/get?access_token=%s"
	orderSearchURL  = "https://api.weixin.qq.com/product/order/search?access_token=%s"
)

const (
	// UNPAY 未支付
	UNPAY OStatus = 10
	// UNDELIVE 待发货
	UNDELIVE OStatus = 20
	// UNRECEVIED 待收货
	UNRECEVIED OStatus = 30
	// COMPLETE 完成
	COMPLETE OStatus = 100
	// CANCEL 全部商品售后之后，订单取消
	CANCEL OStatus = 200
	// OVERTIME 用户主动取消或待付款超时取消
	OVERTIME OStatus = 250
)

// OStatus 订单状态
type OStatus int

//Order 订单接口
type Order struct {
	*context.Context
}

//NewOrder new order
func NewOrder(ctx *context.Context) *Order {
	return &Order{ctx}
}

// GetList 获取订单列表
func (o *Order) GetList(req *GetListParam) ([]byte, error) {
	return o.FetchData(orderGetListURL, req)
}

// ParseOrderList 解析订单列表
func (o *Order) ParseOrderList(t []byte) (*GetListRsp, error) {
	rsp := &GetListRsp{}
	err := json.Unmarshal(t, rsp)
	if err != nil {
		return nil, err
	}
	return rsp, nil
}

// Get 获取订单详情
func (o *Order) Get(orderID string) ([]byte, error) {
	req := map[string]string{
		"order_id": orderID,
	}
	return o.FetchData(orderGetURL, req)
}

// Search 搜索订单
func (o *Order) Search(req SearchParam) ([]byte, error) {
	return o.FetchData(orderSearchURL, req)
}
