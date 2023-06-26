package logic

import (
	"context"

	"go-playground/cmd/zero-demo/rpc/internal/svc"
	"go-playground/cmd/zero-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type PushMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPushMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PushMessageLogic {
	return &PushMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义双向流式 rpc
func (l *PushMessageLogic) PushMessage(stream pb.Greet_PushMessageServer) error {
	// todo: add your logic here and delete this line

	return nil
}
