package logic

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go-looklook/common/ctxdata"
	"go-looklook/common/errorx"
	"time"

	"go-looklook/app/usercenter/cmd/rpc/internal/svc"
	"go-looklook/app/usercenter/cmd/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateTokenLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateTokenLogic {
	return &GenerateTokenLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateTokenLogic) GenerateToken(in *pb.GenerateTokenReq) (*pb.GenerateTokenResp, error) {
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := l.getJwtToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, in.Id)
	if err != nil {
		return nil, errorx.NewErrCodeMsg(
			errorx.InternalError,
			"generate token failed, err="+err.Error(),
		)
	}

	return &pb.GenerateTokenResp{
		AccessExpire: now + accessExpire,
		AccessToken:  accessToken,
		RefreshAfter: now + accessExpire/2,
	}, nil
}

func (l *GenerateTokenLogic) getJwtToken(secretKey string, issueAt, second, userId int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = issueAt + second
	claims["issueAt"] = issueAt
	claims[ctxdata.CtxKeyJwtUserId] = userId
	token := jwt.New(jwt.SigningMethodHS256)
	token.Claims = claims
	return token.SignedString([]byte(secretKey))
}
