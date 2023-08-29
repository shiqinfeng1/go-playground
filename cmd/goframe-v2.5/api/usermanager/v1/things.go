package v1

import "github.com/gogf/gf/v2/frame/g"

type UserReq struct {
	g.Meta `path:"/get" tags:"基础服务" method:"get" summary:"获取验证码"`
}
type UserRes struct {
	g.Meta `mime:"application/json"`
	Key    string `json:"key" dc:"随机码的key"`
	Img    string `json:"img" dc:"随机码图片数据"`
}
