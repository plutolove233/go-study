package user

import (
	"context"

	"go-looklook/app/usercenter/cmd/api/internal/svc"
	"go-looklook/app/usercenter/cmd/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type WxAuthLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewWxAuthLogic(ctx context.Context, svcCtx *svc.ServiceContext) *WxAuthLogic {
	return &WxAuthLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *WxAuthLogic) WxAuth(req *types.WXAuthReq) (resp *types.WXAuthResp, err error) {
	// todo: add your logic here and delete this line

	return
}
