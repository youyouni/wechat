package goods

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/silenceper/wechat/v2/minishop/context"
	"github.com/silenceper/wechat/v2/util"
)

const (
	spuListURL = "https://api.weixin.qq.com/product/spu/get_list?access_token=%s"
)

//Spu 商品接口
type Spu struct {
	*context.Context
}

//NewSpu new spu
func NewSpu(ctx *context.Context) *Spu {
	return &Spu{ctx}
}

// SpuListRsp 商品列表
type SpuListRsp struct {
	util.CommonError
	Data struct {
		Spu struct {
			Title    string   `json:"title"`
			SubTitle string   `json:"sub_title"`
			HeadImg  []string `json:"head_img"`
			DescInfo struct {
				Imgs []string `json:"imgs"`
			} `json:"desc_info"`
			OutProductID string `json:"out_product_id"`
			ProductID    int    `json:"product_id"`
			BrandID      int    `json:"brand_id"`
			Status       int    `json:"status"`
			EditStatus   int    `json:"edit_status"`
			MinPrice     int    `json:"min_price"`
			Cats         []struct {
				CatID int `json:"cat_id"`
				Level int `json:"level"`
			} `json:"cats"`
			Attrs []struct {
				AttrKey   string `json:"attr_key"`
				AttrValue string `json:"attr_value"`
			} `json:"attrs"`
			Model   string `json:"model"`
			Shopcat []struct {
				ShopcatID int `json:"shopcat_id"`
			} `json:"shopcat"`
			Skus []struct {
				SkuID int `json:"sku_id"`
			} `json:"skus"`
		} `json:"spu"`
	} `json:"data"`
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
