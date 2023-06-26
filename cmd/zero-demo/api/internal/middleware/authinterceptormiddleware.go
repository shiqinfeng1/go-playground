package middleware

import (
	"go-playground/cmd/zero-demo/api/internal/logic/ctxdata"
	"net/http"

	"github.com/zeromicro/go-zero/core/logx"
)

type AuthCheck struct {
}

func NewAuthCheck() *AuthCheck {
	return &AuthCheck{}
}

func (m *AuthCheck) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		uid := ctxdata.GetUserId(ctx)
		logx.WithContext(ctx).Infof("get userid by AuthCheck: %v", uid)
		next(w, r)
	}
}
