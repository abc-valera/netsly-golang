package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/json-rest-api/dto"
)

type MeHandler struct {
	userQuery  query.IUser
	userDomain entity.User
}

func NewMeHandler(
	userQuery query.IUser,
	userDomain entity.User,
) MeHandler {
	return MeHandler{
		userQuery:  userQuery,
		userDomain: userDomain,
	}
}

func (h MeHandler) MeGet(ctx context.Context) (*ogen.User, error) {
	user, err := h.userQuery.GetByID(ctx, payloadUserID(ctx))
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponse(user), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) (*ogen.User, error) {
	user, err := h.userDomain.Update(ctx, payloadUserID(ctx), entity.UserUpdateRequest{
		Password: dto.NewPointerString(req.Password),
		Fullname: dto.NewPointerString(req.Fullname),
		Status:   dto.NewPointerString(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return dto.NewUserResponse(user), nil
}

func (h MeHandler) MeDel(ctx context.Context, req *ogen.MeDelReq) error {
	return h.userDomain.Delete(ctx, payloadUserID(ctx), entity.UserDeleteRequest{
		Password: req.Password,
	})
}
