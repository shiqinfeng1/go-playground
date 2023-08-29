// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package usermanager

import (
	"context"
	
	"go-playground/cmd/goframe-v2.5feature/api/usermanager/v1"
)

type IUsermanagerV1 interface {
	User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error)
}


