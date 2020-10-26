package coupon

import "github.com/silenceper/wechat/v2/util"

// GetListParam 请求优惠券列表参数
type GetListParam struct {
	StartCreateTime string  `json:"start_create_time"` // (必填)优惠券创建时间的搜索开始时间
	EndCreateTime   string  `json:"end_create_time"`   // (必填)优惠券创建时间的搜索结束时间
	Status          CStatus `json:"status"`            // (必填)优惠券状态
	Page            int     `json:"page"`              // (必填)第几页（最小填1）
	PageSize        int     `json:"page_size"`         // (必填)每页数量(不超过10,000)
}

// GetListResp 获取优惠券列表
type GetListResp struct {
	util.CommonError
	Coupons []struct {
		CouponID   int    `json:"coupon_id"`
		Type       int    `json:"type"`
		Status     int    `json:"status"`
		CreateTime string `json:"create_time"`
		UpdateTime string `json:"update_time"`
		CouponInfo struct {
			Name      string `json:"name"`
			ValidInfo struct {
				ValidType   int    `json:"valid_type"`
				ValidDayNum int    `json:"valid_day_num"`
				StartTime   string `json:"start_time"`
				EndTime     string `json:"end_time"`
			} `json:"valid_info"`
		} `json:"coupon_info"`
		StockInfo struct {
			IssuedNum  int `json:"issued_num"`
			ReceiveNum int `json:"receive_num"`
			UsedNum    int `json:"used_num"`
		} `json:"stock_info"`
	} `json:"coupons"`
}
