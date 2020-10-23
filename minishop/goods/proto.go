package goods

import "github.com/silenceper/wechat/v2/util"

// CategoryRsp 类目详情
type CategoryRsp struct {
	util.CommonError
	CatList []struct {
		CatID  int    `json:"cat_id"`
		FCatID int    `json:"f_cat_id"`
		Name   string `json:"name"`
	} `json:"cat_list"`
}

// BrandRsp 品牌列表
type BrandRsp struct {
	util.CommonError
	Brands []struct {
		FirstCatID  int `json:"first_cat_id"`
		SecondCatID int `json:"second_cat_id"`
		ThirdCatID  int `json:"third_cat_id"`
		BrandInfo   struct {
			BrandID   int    `json:"brand_id"`
			BrandName string `json:"brand_name"`
		} `json:"brand_info"`
	} `json:"brands"`
}

// DeliveryRsp 运费模板
type DeliveryRsp struct {
	util.CommonError
	TemplateList []struct {
		TemplateID    int    `json:"template_id"`
		Name          string `json:"name"`
		ValuationType int    `json:"valuation_type"`
	} `json:"template_list"`
}

// ShopcatRsp 店铺的商品分类
type ShopcatRsp struct {
	util.CommonError
	ShopcatList []struct {
		ShopcatID   int    `json:"shopcat_id"`
		ShopcatName string `json:"shopcat_name"`
		FShopcatID  int    `json:"f_shopcat_id"`
		CatLevel    int    `json:"cat_level"`
	} `json:"shopcat_list"`
}

// SpuReq 添加商品
type SpuReq struct {
	OutProductID string   `json:"out_product_id"`
	Title        string   `json:"title"`
	SubTitle     string   `json:"sub_title"`
	HeadImg      []string `json:"head_img"`
	DescInfo     struct {
		Imgs []string `json:"imgs"`
	} `json:"desc_info"`
	BrandID int `json:"brand_id"`
	Cats    []struct {
		CatID int `json:"cat_id"`
		Level int `json:"level"`
	} `json:"cats"`
	Attrs []struct {
		AttrKey   string `json:"attr_key"`
		AttrValue string `json:"attr_value"`
	} `json:"attrs"`
	Model       string `json:"model"`
	ExpressInfo struct {
		TemplateID int `json:"template_id"`
	} `json:"express_info"`
	Skus []struct {
		OutProductID string `json:"out_product_id"`
		OutSkuID     string `json:"out_sku_id"`
		ThumbImg     string `json:"thumb_img"`
		SalePrice    int    `json:"sale_price"`
		MarketPrice  int    `json:"market_price"`
		StockNum     int    `json:"stock_num"`
		SkuCode      string `json:"sku_code"`
		Barcode      string `json:"barcode"`
		SkuAttrs     []struct {
			AttrKey   string `json:"attr_key"`
			AttrValue string `json:"attr_value"`
		} `json:"sku_attrs"`
	} `json:"skus"`
}

// SpuRsp 添加商品返回
type SpuRsp struct {
	util.CommonError
	Data struct {
		ProductID    int64  `json:"product_id"`
		OutProductID string `json:"out_product_id"`
		CreateTime   string `json:"create_time"`
	} `json:"data"`
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
