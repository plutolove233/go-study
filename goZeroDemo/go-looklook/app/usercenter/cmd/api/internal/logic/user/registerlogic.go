package user

import (
	"context"
	"go-looklook/app/usercenter/cmd/rpc/usercenter"
	"go-looklook/app/usercenter/model"
	"go-looklook/common/errorx"

	"go-looklook/app/usercenter/cmd/api/internal/svc"
	"go-looklook/app/usercenter/cmd/api/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (*types.RegisterResp, error) {
	registerResp, err := l.svcCtx.UserCenterRpc.Register(l.ctx, &usercenter.RegisterReq{
		Mobile:   req.Mobile,
		Password: req.Password,
		AuthKey:  req.Mobile,
		AuthType: model.UserAuthTypeSystem,
	})
	if err != nil {
		return nil, errorx.NewErrCodeMsg(
			errorx.InternalError,
			"request rpc service failed, err="+err.Error(),
		)
	}
	var resp types.RegisterResp
	_ = copier.Copy(&resp, registerResp)
	return &resp, nil
}
