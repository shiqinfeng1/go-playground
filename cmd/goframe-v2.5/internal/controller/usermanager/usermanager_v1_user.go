package usermanager

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-playground/cmd/goframe-v2.5feature/api/usermanager/v1"
)

func (c *ControllerV1) User(ctx context.Context, req *v1.UserReq) (res *v1.UserRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
