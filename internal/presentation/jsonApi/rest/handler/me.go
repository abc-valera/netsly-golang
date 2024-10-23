package handler

import (
	"context"

	"github.com/abc-valera/netsly-golang/gen/ogen"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query/queryUtil/filter"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/contexts"
	"github.com/abc-valera/netsly-golang/internal/presentation/jsonApi/rest/restDto"
	"go.opentelemetry.io/otel/trace"
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
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := h.user.GetOne(
		ctx,
		filter.By(model.User{ID: userID}),
	)
	if err != nil {
		return nil, err
	}
	return restDto.NewUser(user), nil
}

func (h MeHandler) MePut(ctx context.Context, req *ogen.MePutReq) (*ogen.User, error) {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return nil, err
	}

	user, err := h.user.Update(ctx, userID, entity.UserUpdateRequest{
		Password: restDto.NewString(req.Password),
		Fullname: restDto.NewString(req.Fullname),
		Status:   restDto.NewString(req.Status),
	})
	if err != nil {
		return nil, err
	}
	return restDto.NewUser(user), nil
}

func (h MeHandler) MeDel(ctx context.Context, req *ogen.MeDelReq) error {
	var span trace.Span
	ctx, span = global.NewSpan(ctx)
	defer span.End()

	userID, err := contexts.GetUserID(ctx)
	if err != nil {
		return err
	}

	return h.user.Delete(ctx, userID, entity.UserDeleteRequest{
		Password: req.Password,
	})
}
