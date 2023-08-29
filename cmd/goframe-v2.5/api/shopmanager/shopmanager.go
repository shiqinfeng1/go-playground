// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package shopmanager

import (
	"context"
	
	"go-playground/cmd/goframe-v2.5feature/api/shopmanager/v1"
)

type IShopmanagerV1 interface {
	Shop(ctx context.Context, req *v1.ShopReq) (res *v1.ShopRes, err error)
	ShopList(ctx context.Context, req *v1.ShopListReq) (res *v1.ShopListRes, err error)
	Sreet(ctx context.Context, req *v1.SreetReq) (res *v1.SreetRes, err error)
}


