package goods

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/silenceper/wechat/v2/minishop/context"
)

const (
	spuListURL = "https://api.weixin.qq.com/product/spu/get_list?access_token=%s"
	spuAddURL  = "https://api.weixin.qq.com/product/spu/add?access_token=%s"
)

//Spu 商品接口
type Spu struct {
	*context.Context
}

//NewSpu new spu
func NewSpu(ctx *context.Context) *Spu {
	return &Spu{ctx}
}

// AddSpu 添加商品
func (s *Spu) AddSpu(req *SpuReq) (*SpuRsp, error) {
	response, err := s.FetchData(spuAddURL, req)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &SpuRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetSpuList 获取商品列表
func (s *Spu) GetSpuList(status, page, pageSize int) (*SpuListRsp, error) {
	response, err := s.GetSpuListByte(status, page, pageSize)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &SpuListRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetSpuListByte 获取商品列表
func (s *Spu) GetSpuListByte(status, page, pageSize int) ([]byte, error) {
	req := map[string]string{
		"status":    strconv.Itoa(status),
		"page":      strconv.Itoa(page),
		"page_size": strconv.Itoa(pageSize),
	}
	return s.request(spuListURL, req)
}

func (s *Spu) request(reqURL string, req map[string]string) ([]byte, error) {
	return s.FetchData(reqURL, req)
}
