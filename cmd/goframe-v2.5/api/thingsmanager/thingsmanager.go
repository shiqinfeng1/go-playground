// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT. 
// =================================================================================

package thingsmanager

import (
	"context"
	
	"go-playground/cmd/goframe-v2.5feature/api/thingsmanager/v1"
)

type IThingsmanagerV1 interface {
	Things(ctx context.Context, req *v1.ThingsReq) (res *v1.ThingsRes, err error)
}


