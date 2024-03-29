package persistence_test

import (
	"testing"

	"github.com/abc-valera/netsly-api-golang/pkg/domain/persistence/query/selectParams"
	"github.com/stretchr/testify/require"
)

func TestUserQuery(t *testing.T) {
	ctx, _, _, queries := initTest(t)

	t.Run("GetByID", func(t *testing.T) {
		r := require.New(t)

		user, err := queries.User.GetByID(ctx, "00000000-0000-0000-0000-000000000000")
		r.NoError(err)
		r.NotEmpty(user)
	})

	t.Run("GetByUsername", func(t *testing.T) {
		r := require.New(t)

		user, err := queries.User.GetByUsername(ctx, "abc-valera")
		r.NoError(err)
		r.NotEmpty(user)
	})

	t.Run("GetByEmail", func(t *testing.T) {
		r := require.New(t)

		user, err := queries.User.GetByEmail(ctx, "abc-valera@gmail.com")
		r.NoError(err)
		r.NotEmpty(user)
	})

	t.Run("SearchAllByUsername", func(t *testing.T) {
		r := require.New(t)

		users, err := queries.User.SearchAllByUsername(ctx, "", selectParams.NewSelectParams(selectParams.OrderDesc, 3, 0))
		r.NoError(err)
		r.NotEmpty(users)
		r.Equal(3, len(users))
	})
}
