package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"go-looklook/app/usercenter/model"
	"go-looklook/common/errorx"
	"go-looklook/common/tools"
	"golang.org/x/crypto/bcrypt"

	"go-looklook/app/usercenter/cmd/rpc/internal/svc"
	"go-looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *pb.RegisterReq) (*pb.RegisterResp, error) {
	user, err := l.svcCtx.UserModel.FindOneByMobile(l.ctx, in.Mobile)
	if err != nil && err != model.ErrNotFound {
		return nil, errorx.NewErrCodeMsg(errorx.DataError, "数据库异常")
	}

	if user != nil {
		return nil, errorx.NewErrCodeMsg(errorx.DataExist, "用户已经注册")
	}
	var userId int64
	if err := l.svcCtx.UserModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		var user model.User
		user.Mobile = in.Mobile
		user.Nickname = tools.Krand(8, tools.KC_RAND_KIND_ALL)
		hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return errorx.NewErrCodeMsg(errorx.InternalError, "password encode failed")
		}
		user.Password = string(hash)
		insertRes, err := l.svcCtx.UserModel.Insert(l.ctx, session, &user)
		if err != nil {
			return errorx.NewErrCodeMsg(
				errorx.DBError,
				"数据插入异常",
			)
		}
		lastId, err := insertRes.LastInsertId()
		if err != nil {
			return errorx.NewErrCodeMsg(errorx.DBError, "get last insert id failed, err = "+err.Error())
		}

		userId = lastId

		userAuth := new(model.UserAuth)
		userAuth.UserId = userId
		userAuth.AuthKey = in.AuthKey
		userAuth.AuthType = in.AuthType
		if _, err := l.svcCtx.UserAuthModel.Insert(ctx, session, userAuth); err != nil {
			return errorx.NewErrCodeMsg(
				errorx.DBError,
				"Register db user_auth Insert err",
			)
		}
		return nil
	}); err != nil {
		return nil, err
	}
	generateTokenLogic := NewGenerateTokenLogic(l.ctx, l.svcCtx)
	tokenResp, err := generateTokenLogic.GenerateToken(&pb.GenerateTokenReq{
		Id: userId,
	})
	if err != nil {
		return nil, errorx.NewErrCodeMsg(
			errorx.InternalError,
			"generate token failed",
		)
	}
	return &pb.RegisterResp{
		AccessToken:  tokenResp.AccessToken,
		AccessExpire: tokenResp.AccessExpire,
		RefreshAfter: tokenResp.RefreshAfter,
	}, nil
}
