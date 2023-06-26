package user

import (
	"context"
	"errors"
	"fmt"
	"time"

	"go-playground/cmd/zero-demo/api/internal/svc"
	"go-playground/cmd/zero-demo/api/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

type MyClaims struct {
	Userid               uint64 `json:"userId"`
	jwt.RegisteredClaims        // 注意!这是jwt-go的v4版本新增的，原先是jwt.StandardClaims
}

// @secretKey: JWT 加解密密钥
// @iat: 时间戳
// @seconds: 过期时间，单位秒
// @payload: 数据载体
func getJwtToken(secretKey []byte, userId uint64) (string, error) {
	claim := MyClaims{
		Userid: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Second * time.Duration(1))), // 过期时间3小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                                          // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                                          // 生效时间
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim) // 使用HS256算法
	resp, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("token sign:", err)
		return "", err
	}
	return resp, nil
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
	resp.Token, err = getJwtToken([]byte(l.svcCtx.Config.Auth.AccessSecret), 1254444)
	return
}
