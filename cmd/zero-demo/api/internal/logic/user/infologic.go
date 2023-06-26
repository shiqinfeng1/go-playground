package user

import (
	"context"

	"go-playground/cmd/zero-demo/api/internal/logic/ctxdata"
	"go-playground/cmd/zero-demo/api/internal/svc"
	"go-playground/cmd/zero-demo/api/internal/types"
	"go-playground/cmd/zero-demo/rpc/greet"

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
	smr := new(greet.SendMessageReq)
	smr.Message = "hello rpc greet"
	smrsp, err := l.svcCtx.Greet.SendMessage(l.ctx, smr)
	if err != nil {
		return nil, err
	}

	resp = &types.InfoResp{}
	resp.Id = int64(smrsp.Status)
	resp.Address = "addr"
	resp.Name = "name"
	return
}
