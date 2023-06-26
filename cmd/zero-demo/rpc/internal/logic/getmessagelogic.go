package logic

import (
	"context"

	"go-playground/cmd/zero-demo/rpc/internal/svc"
	"go-playground/cmd/zero-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义服务端流式 rpc
func (l *GetMessageLogic) GetMessage(in *pb.GetMessageReq, stream pb.Greet_GetMessageServer) error {
	// todo: add your logic here and delete this line

	return nil
}
