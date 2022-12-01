package logic

import (
	"context"
	"demo1/user/rpc/types/user"
	"errors"

	"demo1/order/api/internal/svc"
	"demo1/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetOrderLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetOrderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetOrderLogic {
	return &GetOrderLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetOrderLogic) GetOrder(req *types.OrderReq) (resp *types.OrderReply, err error) {
	// todo: add your logic here and delete this line
	_user, err := l.svcCtx.UserRpc.GetUser(l.ctx, &user.IdRequest{ //调用user的rpc服务
		Id: "1",
	})
	if err != nil {
		return nil, err
	}

	if _user.Name != "test" {
		return nil, errors.New("用户不存在")
	}

	return &types.OrderReply{
		Id:   req.Id,
		Name: "test order",
	}, nil
}
