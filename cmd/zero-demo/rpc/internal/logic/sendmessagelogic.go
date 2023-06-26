package logic

import (
	"context"

	"go-playground/cmd/zero-demo/rpc/internal/svc"
	"go-playground/cmd/zero-demo/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 定义客户端流式 rpc
func (l *SendMessageLogic) SendMessage(in *pb.SendMessageReq) (*pb.SendMessageResp, error) {
	// todo: add your logic here and delete this line

	return &pb.SendMessageResp{Status: 123789}, nil
}
