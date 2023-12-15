package handler

import (
	"context"

	"github.com/abc-valera/flugo-api-golang/gen/ogen"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/domain"
	"github.com/abc-valera/flugo-api-golang/internal/core/domain/repository/query"
	"github.com/abc-valera/flugo-api-golang/internal/port/http/dto"
)

type MeHandler struct {
	userQuery  query.IUserQuery
	userDomain domain.UserDomain
}

func NewMeHandler(
	userQuery query.IUserQuery,
	userDomain domain.UserDomain,
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

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) error {
	return h.userDomain.Update(ctx, payloadUserID(ctx), domain.UserUpdateRequest{
		Password: dto.NewPointerString(req.Password),
		Fullname: dto.NewPointerString(req.Fullname),
		Status:   dto.NewPointerString(req.Status),
	})
}

func (h MeHandler) MeDel(ctx context.Context, req *ogen.MeDelReq) error {
	return h.userDomain.Delete(ctx, payloadUserID(ctx), domain.UserDeleteRequest{
		Password: req.Password,
	})
}
