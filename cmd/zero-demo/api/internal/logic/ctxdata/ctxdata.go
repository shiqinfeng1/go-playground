package ctxdata

import (
	"context"
	"encoding/json"

	"github.com/zeromicro/go-zero/core/logx"
)

const CtxKeyJwtUserId = "userId"

func GetUserId(ctx context.Context) uint64 {
	var uid uint64
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = uint64(int64Uid)
		} else {
			logx.WithContext(ctx).Errorf("GetUidFromCtx err : %+v", err)
		}
	}
	return uid
}
