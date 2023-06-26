package user

import (
	"context"

	"go-playground/cmd/zero-demo/api/internal/svc"
	"go-playground/cmd/zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSetLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetLogic {
	return &SetLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SetLogic) Set(req *types.SetInfoReq) (resp *types.SetInfoResp, err error) {
	// todo: add your logic here and delete this line

	return
}
