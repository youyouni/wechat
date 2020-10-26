package service

import "github.com/silenceper/wechat/v2/util"

// GetOrderListRsp 用户购买的服务列表
type GetOrderListRsp struct {
	util.CommonError
	ServiceOrderList []struct {
		ServiceOrderID  int    `json:"service_order_id"`
		ServiceID       int    `json:"service_id"`
		ServiceName     string `json:"service_name"`
		CreateTime      string `json:"create_time"`
		ExpireTime      string `json:"expire_time"`
		ServiceType     int    `json:"service_type"`
		SpecificationID string `json:"specification_id"`
		TotalPrice      int    `json:"total_price"`
	} `json:"service_order_list"`
}

// GetListRsp 用户购买的在有效期内的服务列表
type GetListRsp struct {
	Errcode     int    `json:"errcode"`
	Errmsg      string `json:"errmsg"`
	ServiceList []struct {
		ServiceID   int    `json:"service_id"`
		ServiceName string `json:"service_name"`
		ExpireTime  string `json:"expire_time"`
		ServiceType int    `json:"service_type"`
	} `json:"service_list"`
}
