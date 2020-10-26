package coupon

import (
	"encoding/json"

	"github.com/silenceper/wechat/v2/minishop/context"
)

const (
	couponGetListURL = "https://api.weixin.qq.com/product/coupon/get_list?access_token=%s"
	couponPushURL    = "https://api.weixin.qq.com/product/coupon/push?access_token=%s"
)

const (
	// GOODSDISCOUNT 商品条件折扣券
	GOODSDISCOUNT CStatus = 1
	// GOODSFULLDISCOUNT 商品满减券
	GOODSFULLDISCOUNT CStatus = 2
	// GOODSALLDISCOUNT 商品统一折扣券
	GOODSALLDISCOUNT CStatus = 3
	// GOODSDDISCOUNT 商品直减券
	GOODSDDISCOUNT CStatus = 4
	// SHOPSDISCOUNT 店铺条件折扣券
	SHOPSDISCOUNT CStatus = 101
	// SHOPSFULLDISCOUNT 店铺满减券
	SHOPSFULLDISCOUNT CStatus = 102
	// SHOPSALLDISCOUNT 店铺统一折扣券
	SHOPSALLDISCOUNT CStatus = 103
	// SHOPSDDISCOUNT 店铺直减券
	SHOPSDDISCOUNT CStatus = 104
)

// CStatus 折扣券类型
type CStatus int

//Coupon 优惠券接口
type Coupon struct {
	*context.Context
}

//NewCoupon new order
func NewCoupon(ctx *context.Context) *Coupon {
	return &Coupon{ctx}
}

// GetList 获取优惠券列表
func (c *Coupon) GetList(req *GetListParam) (*GetListResp, error) {
	info := &GetListResp{}
	response, err := c.FetchData(couponGetListURL, req)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(response, info)
	if err != nil {
		return nil, err
	}
	return info, err
}

// Push 发放优惠券
func (c *Coupon) Push(openID, couponID string) error {
	req := map[string]string{
		"openid":    openID,
		"coupon_id": couponID,
	}
	_, err := c.FetchData(couponPushURL, req)
	return err
}
