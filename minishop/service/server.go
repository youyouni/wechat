package service

import (
	"encoding/json"
	"strconv"

	"github.com/silenceper/wechat/v2/minishop/context"
)

const (
	serverGetList      = "https://api.weixin.qq.com/product/service/get_list?access_token=%s"
	serverGetOrderList = "https://api.weixin.qq.com/product/service/get_order_list?access_token=%s"
)

//Service 服务商接口
type Service struct {
	*context.Context
}

//NewService 初始化服务商
func NewService(ctx *context.Context) *Service {
	return &Service{ctx}
}

// GetOrderList 获取用户购买的服务列表
func (s *Service) GetOrderList(start, end string, page, pageSize int) (*GetOrderListRsp, error) {
	info := &GetOrderListRsp{}
	req := map[string]string{
		"start_create_time": start,
		"end_create_time":   end,
		"page":              strconv.Itoa(page),
		"page_size":         strconv.Itoa(pageSize),
	}
	response, err := s.FetchData(serverGetOrderList, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetList 获取订单详情
func (s *Service) GetList() (*GetListRsp, error) {
	info := &GetListRsp{}
	req := map[string]string{}

	response, err := s.FetchData(serverGetList, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, &info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
