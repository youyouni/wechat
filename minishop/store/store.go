package store

import (
	"encoding/json"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	storeInfoURL = "https://api.weixin.qq.com/product/store/get_info?access_token=%s"
)

// Store 商户
type Store struct {
	*context.Context
}

//NewStore new order
func NewStore(ctx *context.Context) *Store {
	return &Store{ctx}
}

// InfoRsp 店铺信息
type InfoRsp struct {
	util.CommonError
	Data struct {
		StoreName string `json:"store_name"`
		Logo      string `json:"logo"`
	} `json:"data"`
}

// GetStoreInfo 获取用户购买的服务列表
func (s *Store) GetStoreInfo(start, end string, page, pageSize int) (*InfoRsp, error) {
	info := &InfoRsp{}
	req := map[string]string{}
	response, err := s.FetchData(storeInfoURL, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
