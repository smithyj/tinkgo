package account

import (
	"context"
	"tinkgo/service/app/api/internal/svc"
)

type UpdateProfileRequest struct{}

type UpdateProfileResponse struct{}

type UpdateProfileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateProfileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProfileLogic {
	return &UpdateProfileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (logic *UpdateProfileLogic) UpdateProfile(req *UpdateProfileRequest) (*UpdateProfileResponse, error) {
	return &UpdateProfileResponse{}, nil
}
