package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
)

type MeHandler struct {
	user entity.IUser
}

func NewMeHandler(
	user entity.IUser,
) MeHandler {
	return MeHandler{
		user: user,
	}
}

func (h MeHandler) MeGet(ctx context.Context) (*ogen.User, error) {
	user, err := h.user.GetByID(ctx, payloadUserID(ctx))
	if err != nil {
		return nil, err
	}
	return restDto.NewUserResponse(user), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) (*ogen.User, error) {
	user, err := h.user.Update(ctx, payloadUserID(ctx), entity.UserUpdateRequest{
		Password: restDto.NewDomainOptionalString(req.Password),
		Fullname: restDto.NewDomainOptionalString(req.Fullname),
		Status:   restDto.NewDomainOptionalString(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewUserResponse(user), nil
}

func (h MeHandler) MeDel(ctx context.Context, req *ogen.MeDelReq) error {
	return h.user.Delete(ctx, payloadUserID(ctx), entity.UserDeleteRequest{
		Password: req.Password,
	})
}
