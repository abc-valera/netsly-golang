package entity

import (
	"context"

	"github.com/abc-valera/netsly-golang/internal/domain/persistence"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/query"
	"github.com/abc-valera/netsly-golang/internal/domain/service"
)

type Entities struct {
	User        IUser
	Joke        IJoke
	Like        ILike
	Comment     IComment
	Room        IRoom
	RoomMember  IRoomMember
	RoomMessage IRoomMessage
	File        IFile

	Passworder IPassworder
	Emailer    IEmailer
}

func NewEntities(dep IDependency) Entities {
	return Entities{
		User:        newUser(dep),
		Joke:        newJoke(dep),
		Like:        newLike(dep),
		Comment:     newComment(dep),
		Room:        newRoom(dep),
		RoomMember:  newRoomMember(dep),
		RoomMessage: newRoomMessage(dep),
		File:        newFile(dep),

		Passworder: newPassworder(dep),
		Emailer:    newEmailer(dep),
	}
}

type IDependency interface {
	BeginTX(ctx context.Context) (IDependencyTX, error)
	RunInTX(
		ctx context.Context,
		fn func(context.Context, command.Commands, query.Queries, Entities) error,
	) error

	C() command.Commands
	Q() query.Queries
	S() service.Services
	E() Entities
}

type IDependencyTX interface {
	IDependency

	Commit() error
	Rollback() error
}

type dependency struct {
	db       persistence.IDB
	services service.Services
}

func NewDependency(db persistence.IDB, services service.Services) IDependency {
	return dependency{
		db:       db,
		services: services,
	}
}

func (dep dependency) BeginTX(ctx context.Context) (IDependencyTX, error) {
	tx, err := dep.db.BeginTX(ctx)
	if err != nil {
		return nil, err
	}

	return dependencyTX{
		tx:       tx,
		services: dep.services,
	}, nil
}

func (dep dependency) RunInTX(
	ctx context.Context,
	fn func(context.Context, command.Commands, query.Queries, Entities) error,
) error {
	tx, err := dep.db.BeginTX(ctx)
	if err != nil {
		return err
	}

	if err := fn(ctx, tx.Commands(), tx.Queries(), NewEntities(NewDependency(tx, dep.services))); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()
}

func (dep dependency) C() command.Commands {
	return dep.db.Commands()
}

func (dep dependency) Q() query.Queries {
	return dep.db.Queries()
}

func (dep dependency) S() service.Services {
	return dep.services
}

func (dep dependency) E() Entities {
	return NewEntities(dep)
}

type dependencyTX struct {
	tx       persistence.ITX
	services service.Services
}

func (depTX dependencyTX) BeginTX(ctx context.Context) (IDependencyTX, error) {
	nestedTX, err := depTX.tx.BeginTX(ctx)
	if err != nil {
		return nil, err
	}

	return dependencyTX{
		tx:       nestedTX,
		services: depTX.services,
	}, nil
}

func (depTX dependencyTX) RunInTX(
	ctx context.Context,
	fn func(context.Context, command.Commands, query.Queries, Entities) error,
) error {
	nestedTX, err := depTX.tx.BeginTX(ctx)
	if err != nil {
		return err
	}

	if err := fn(ctx, nestedTX.Commands(), nestedTX.Queries(), NewEntities(NewDependency(nestedTX, depTX.services))); err != nil {
		nestedTX.Rollback()
		return err
	}

	return nestedTX.Commit()
}

func (depTX dependencyTX) C() command.Commands {
	return depTX.tx.Commands()
}

func (depTX dependencyTX) Q() query.Queries {
	return depTX.tx.Queries()
}

func (depTX dependencyTX) S() service.Services {
	return depTX.services
}

func (depTX dependencyTX) E() Entities {
	return NewEntities(depTX)
}

func (depTX dependencyTX) Commit() error {
	return depTX.tx.Commit()
}

func (depTX dependencyTX) Rollback() error {
	return depTX.tx.Rollback()
}
