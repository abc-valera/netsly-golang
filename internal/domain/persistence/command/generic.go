package command

import "context"

type ICreate[DomainModel any] interface {
	Create(ctx context.Context, req DomainModel) error
}

type ICreateDelete[DomainModel any] interface {
	Create(ctx context.Context, req DomainModel) error
	Delete(ctx context.Context, req DomainModel) error
}

type IUpdateDelete[DomainModel any] interface {
	Update(ctx context.Context, req DomainModel) error
	Delete(ctx context.Context, req DomainModel) error
}

type ICreateUpdateDelete[DomainModel any] interface {
	Create(ctx context.Context, req DomainModel) error
	Update(ctx context.Context, req DomainModel) error
	Delete(ctx context.Context, req DomainModel) error
}
