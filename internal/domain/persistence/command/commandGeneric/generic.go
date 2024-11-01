package commandGeneric

import "context"

// This file contains generic interfaces for command persistence.
//
// Every one of these interfaces contains a some combination
// of Create, Update, and Delete methods.

type ICreate[Model any] interface {
	Create(context.Context, Model) error
}

type IUpdate[Model any] interface {
	Update(context.Context, Model) error
}

type IDelete[Model any] interface {
	Delete(context.Context, Model) error
}

type ICreateDelete[Model any] interface {
	ICreate[Model]
	IDelete[Model]
}

type IUpdateDelete[Model any] interface {
	IUpdate[Model]
	IDelete[Model]
}

type ICreateUpdateDelete[Model any] interface {
	ICreate[Model]
	IUpdate[Model]
	IDelete[Model]
}
