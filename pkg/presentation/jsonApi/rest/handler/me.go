package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/entity"
	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/pkg/presentation/jsonApi/rest/restDto"
)

type MeHandler struct {
	userQuery  query.IUser
	userEntity entity.IUser
}

func NewMeHandler(
	userQuery query.IUser,
	userEntity entity.IUser,
) MeHandler {
	return MeHandler{
		userQuery:  userQuery,
		userEntity: userEntity,
	}
}

func (h MeHandler) MeGet(ctx context.Context) (*ogen.User, error) {
	user, err := h.userQuery.GetByID(ctx, payloadUserID(ctx))
	if err != nil {
		return nil, err
	}
	return restDto.NewUserResponse(user), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) (*ogen.User, error) {
	user, err := h.userEntity.Update(ctx, payloadUserID(ctx), entity.UserUpdateRequest{
		Password: restDto.NewPointerString(req.Password),
		Fullname: restDto.NewPointerString(req.Fullname),
		Status:   restDto.NewPointerString(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewUserResponse(user), nil
}

func (h MeHandler) MeDel(ctx context.Context, req *ogen.MeDelReq) error {
	return h.userEntity.Delete(ctx, payloadUserID(ctx), entity.UserDeleteRequest{
		Password: req.Password,
	})
}
