package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID             string    `bun:",pk,type:uuid"`
	Username       string    `bun:",unique,notnull"`
	Email          string    `bun:",unique,notnull"`
	HashedPassword string    `bun:",notnull"`
	Fullname       string    `bun:",notnull"`
	Status         string    `bun:",notnull"`
	UpdatedAt      time.Time `bun:",notnull"`
	CreatedAt      time.Time `bun:",notnull"`
	DeletedAt      time.Time `bun:",notnull"`

	Jokes        Jokes        `bun:"rel:has-many,join:id=user_id"`
	Comments     Comments     `bun:"rel:has-many,join:id=user_id"`
	Likes        Likes        `bun:"rel:has-many,join:id=user_id"`
	CreatedRooms Rooms        `bun:"rel:has-many,join:id=creator_user_id"`
	RoomMembers  RoomMembers  `bun:"rel:has-many,join:id=user_id"`
	RoomMessages RoomMessages `bun:"rel:has-many,join:id=user_id"`
}

func (u User) ToDomain() model.User {
	return model.User{
		ID:             u.ID,
		Username:       u.Username,
		Email:          u.Email,
		HashedPassword: u.HashedPassword,
		Fullname:       u.Fullname,
		Status:         u.Status,
		UpdatedAt:      u.UpdatedAt,
		CreatedAt:      u.CreatedAt,
		DeletedAt:      u.DeletedAt,
	}
}

type Users []User

func (u Users) ToDomain() model.Users {
	users := make(model.Users, 0, len(u))
	for _, user := range u {
		users = append(users, user.ToDomain())
	}
	return users
}
