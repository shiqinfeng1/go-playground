package user

import (
	"context"
	"errors"
	"time"

	"go-playground/cmd/zero-demo/api/internal/svc"
	"go-playground/cmd/zero-demo/api/internal/types"

	"github.com/dgrijalva/jwt-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func getJwtToken(secretKey string, iat, seconds int64, payload string) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["payload"] = payload
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	if req.Username != "aa" || req.Password != "bb" {
		err = errors.New("user is not aa")
		return
	}
	resp = &types.LoginResp{}
	iat := time.Now().Unix()
	resp.Token, err = getJwtToken(l.svcCtx.Config.Auth.AccessSecret, iat, iat+1000, req.Username)
	return
}
