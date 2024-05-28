package handler

import (
	"context"

	"github.com/abc-valera/netsly-api-golang/gen/ogen"
	"github.com/abc-valera/netsly-api-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-api-golang/internal/presentation/jsonApi/rest/restDto"
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
	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := h.user.GetByID(ctx, userID)
	if err != nil {
		return nil, err
	}
	return restDto.NewUser(user), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) (*ogen.User, error) {
	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := h.user.Update(ctx, userID, entity.UserUpdateRequest{
		Password: restDto.NewDomainOptionalString(req.Password),
		Fullname: restDto.NewDomainOptionalString(req.Fullname),
		Status:   restDto.NewDomainOptionalString(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewUser(user), nil
}

func (h MeHandler) MeDel(ctx context.Context, req *ogen.MeDelReq) error {
	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return err
	}

	return h.user.Delete(ctx, userID, entity.UserDeleteRequest{
		Password: req.Password,
	})
}
