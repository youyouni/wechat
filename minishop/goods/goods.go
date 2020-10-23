package goods

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/silenceper/wechat/v2/minishop/context"
)

const (
	goodsCategory = "https://api.weixin.qq.com/product/category/get?access_token=%s"
	goodsBrand    = "https://api.weixin.qq.com/product/brand/get?access_token=%s"
	goodsDelivery = "https://api.weixin.qq.com/product/delivery/get_freight_template?access_token=%s"
	goodsShopcat  = "https://api.weixin.qq.com/product/store/get_shopcat?access_token=%s"
)

//Goods 商品前接口
type Goods struct {
	*context.Context
}

//NewGoods new spu
func NewGoods(ctx *context.Context) *Goods {
	return &Goods{ctx}
}

// GetGoodsCategory 获取类目详情
func (g *Goods) GetGoodsCategory(fCatID int) (*CategoryRsp, error) {
	req := map[string]string{
		"f_cat_id": strconv.Itoa(fCatID),
	}

	response, err := g.FetchData(goodsCategory, req)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &CategoryRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetGoodsBrand 获取品牌列表
func (g *Goods) GetGoodsBrand() (*BrandRsp, error) {
	req := map[string]string{}

	response, err := g.FetchData(goodsBrand, req)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &BrandRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetDelivery 获取运费模板
func (g *Goods) GetDelivery() (*DeliveryRsp, error) {
	req := map[string]string{}

	response, err := g.FetchData(goodsDelivery, req)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &DeliveryRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}

// GetShopcat 获取店铺的商品分类
func (g *Goods) GetShopcat() (*ShopcatRsp, error) {
	req := map[string]string{}

	response, err := g.FetchData(goodsShopcat, req)
	if err != nil {
		return nil, fmt.Errorf("empty params")
	}
	info := &ShopcatRsp{}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, nil
}
