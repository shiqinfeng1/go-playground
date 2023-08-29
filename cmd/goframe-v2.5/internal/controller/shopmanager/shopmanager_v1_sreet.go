package shopmanager

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"go-playground/cmd/goframe-v2.5feature/api/shopmanager/v1"
)

func (c *ControllerV1) Sreet(ctx context.Context, req *v1.SreetReq) (res *v1.SreetRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
