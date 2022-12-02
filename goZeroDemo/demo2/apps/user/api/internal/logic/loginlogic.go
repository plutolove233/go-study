package logic

import (
	"context"
	"demo2/apps/user/api/internal/svc"
	"demo2/apps/user/api/internal/types"
	"demo2/apps/user/model"
	"demo2/common/errorx"
	"github.com/golang-jwt/jwt/v4"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginReply, err error) {
	if len(req.UserName) == 0 || len(req.Password) == 0 {
		return nil, errorx.NewCodeError(errorx.ParameterIllegal, "param error")
	}

	userInfo, err := l.svcCtx.UserModel.FindOneByNumber(l.ctx, req.UserName)
	switch err {
	case nil:
	case model.ErrNotFound:
		return nil, errorx.NewCodeError(errorx.NotData, "user not exist")
	default:
		return nil, err
	}
	if userInfo.Password != req.Password {
		return nil, errorx.NewCodeError(errorx.PasswordError, "wrong password")
	}

	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	jwtToken, err := l.getJwtToken(l.svcCtx.Config.Auth.AccessSecret, now, l.svcCtx.Config.Auth.AccessExpire, userInfo.Id)
	if err != nil {
		return nil, errorx.NewCodeError(errorx.InternalError, "token生成失败")
	}

	return &types.LoginReply{
		Id:           userInfo.Id,
		Name:         userInfo.Name,
		Gender:       userInfo.Gender,
		AccessToken:  jwtToken,
		AccessExpire: now + accessExpire,
		RefreshAfter: now + accessExpire,
	}, nil
}

func (l *LoginLogic) getJwtToken(secretKey string, iat, seconds, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = iat + seconds
	claims["iat"] = iat
	claims["userId"] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
