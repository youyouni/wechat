package customservice

import "github.com/silenceper/wechat/v2/officialaccount/context"

//CustomService 客服消息管理
type CustomService struct {
	*context.Context
}

//NewCustomService new
func NewCustomService(ctx *context.Context) *CustomService {
	return &CustomService{ctx}
}

//Add 添加客服帐号
func (customService *CustomService) Add(kfAccount, nickname, password string) error {
	//TODO 接口实现
	return nil
}

//TODO 其他接口
