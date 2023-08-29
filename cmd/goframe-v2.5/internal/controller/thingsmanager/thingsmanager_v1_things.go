package thingsmanager

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-playground/cmd/goframe-v2.5feature/api/thingsmanager/v1"
)

func (c *ControllerV1) Things(ctx context.Context, req *v1.ThingsReq) (res *v1.ThingsRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
