package passport

import (
	"context"
	"errors"
	"tinkgo/service/app/api/internal/svc"
)

type LoginRequest struct{}

type LoginResponse struct{}

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *LoginLogic) Login(req *LoginRequest) (*LoginResponse, error) {
	return &LoginResponse{}, errors.New("dsfsdafsf")
}
