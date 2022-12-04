// Package ctxdata
/*
@Coding : utf-8
@time : 2022/12/4 10:34
@Author : yizhigopher
@Software : GoLand
*/
package ctxdata

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
)

var CtxKeyJwtUserId = "jwtUserId"

func GetUidFromCtx(ctx context.Context) (uid int64) {
	if jsonUid, ok := ctx.Value(CtxKeyJwtUserId).(json.Number); ok {
		if int64Uid, err := jsonUid.Int64(); err == nil {
			uid = int64Uid
		} else {
			logx.WithContext(ctx).Errorf("get uid from context failed, err:%v", err)
		}
	}
	return
}
