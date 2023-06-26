package user

import (
	"context"

	"go-playground/cmd/zero-demo/api/internal/logic/ctxdata"
	"go-playground/cmd/zero-demo/api/internal/svc"
	"go-playground/cmd/zero-demo/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type InfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *InfoLogic {
	return &InfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *InfoLogic) Info(req *types.InfoReq) (resp *types.InfoResp, err error) {
	logx.WithContext(l.ctx).Errorf("GetUidFromCtx : %+v", ctxdata.GetUserId(l.ctx))
	resp = &types.InfoResp{}
	resp.Id = 1
	resp.Address = "addr"
	resp.Name = "name"
	return
}
