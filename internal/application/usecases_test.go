package application

import (
	"context"
	"os"
	"testing"

	"github.com/abc-valera/netsly-golang/gen/mocks"
	"github.com/abc-valera/netsly-golang/internal/domain/entity"
	"github.com/abc-valera/netsly-golang/internal/domain/global"
	"github.com/abc-valera/netsly-golang/internal/infrastructure/globals/logger/loggerNop"
	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/trace/noop"
)

func TestMain(m *testing.M) {
	// Init global variables. Note, thatgenerally we don't want to mock it,
	// just use noop variants and make sure it's not null.
	global.Init(
		global.ModeProduction,
		noop.NewTracerProvider().Tracer("testing"),
		loggerNop.New(),
	)

	os.Exit(m.Run())
}

// NewTest is a helper function that sets the test up.
//
// It returns:
//   - the default and instrumented background contexts
//     (the default should be passed to the usecase and the instrumented one should be used inside the EXPECT calls),
//   - instance of require.Assertions,
//   - usecase dependency initialized with mocks.
func NewTest(t *testing.T) (
	defaultContext context.Context,
	instrumentedContext context.Context,
	r *require.Assertions,
	dep MockDependency,
) {
	defaultContext = context.Background()
	instrumentedContext, span := global.NewSpan(defaultContext)
	defer span.End()

	return defaultContext, instrumentedContext, require.New(t), NewMockDependency(t)
}

type MockDependency struct {
	MockEntities MockEntities
}

var _ IDependency = (*MockDependency)(nil)

func NewMockDependency(t *testing.T) MockDependency {
	return MockDependency{
		MockEntities: NewMockEntities(t),
	}
}

func (m MockDependency) BeginTX(_ context.Context) (IDependencyTX, error) {
	return MockTX{
		MockDependency: m,
	}, nil
}

func (m MockDependency) RunInTX(
	ctx context.Context,
	fn func(context.Context, entity.Entities, Usecases) error,
) error {
	return fn(ctx, m.MockEntities.ToDomain(), NewUsecases(m))
}

func (m MockDependency) U() Usecases {
	return NewUsecases(m)
}

func (m MockDependency) E() entity.Entities {
	return m.MockEntities.ToDomain()
}

type MockTX struct {
	MockDependency
}

var _ IDependencyTX = (*MockTX)(nil)

func (m MockTX) BeginTX(_ context.Context) (IDependencyTX, error) {
	return m, nil
}

func (m MockTX) RunInTX(
	ctx context.Context,
	fn func(context.Context, entity.Entities, Usecases) error,
) error {
	return fn(ctx, m.MockEntities.ToDomain(), NewUsecases(m))
}

func (m MockTX) E() entity.Entities {
	return m.MockEntities.ToDomain()
}

func (MockTX) Commit() error {
	return nil
}

func (MockTX) Rollback() error {
	return nil
}

type MockEntities struct {
	User        *mocks.User
	Joke        *mocks.Joke
	Like        *mocks.Like
	Comment     *mocks.Comment
	Room        *mocks.Room
	RoomMember  *mocks.RoomMember
	RoomMessage *mocks.RoomMessage
	File        *mocks.File

	Passworder *mocks.Passworder
	Emailer    *mocks.Emailer
}

func NewMockEntities(t *testing.T) MockEntities {
	return MockEntities{
		User:        mocks.NewUser(t),
		Joke:        mocks.NewJoke(t),
		Like:        mocks.NewLike(t),
		Comment:     mocks.NewComment(t),
		Room:        mocks.NewRoom(t),
		RoomMember:  mocks.NewRoomMember(t),
		RoomMessage: mocks.NewRoomMessage(t),
		File:        mocks.NewFile(t),

		Passworder: mocks.NewPassworder(t),
		Emailer:    mocks.NewEmailer(t),
	}
}

func (m MockEntities) ToDomain() entity.Entities {
	return entity.Entities{
		User:        m.User,
		Joke:        m.Joke,
		Like:        m.Like,
		Comment:     m.Comment,
		Room:        m.Room,
		RoomMember:  m.RoomMember,
		RoomMessage: m.RoomMessage,
		File:        m.File,

		Passworder: m.Passworder,
		Emailer:    m.Emailer,
	}
}
