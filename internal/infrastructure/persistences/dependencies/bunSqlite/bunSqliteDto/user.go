package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel

	ID             string    `bun:"id,pk,type:uuid"`
	Username       string    `bun:",unique,notnull"`
	Email          string    `bun:",unique,notnull"`
	HashedPassword string    `bun:",notnull"`
	Fullname       string    `bun:",notnull"`
	Status         string    `bun:",notnull"`
	UpdatedAt      time.Time `bun:",notnull"`
	CreatedAt      time.Time `bun:",notnull"`
	DeletedAt      time.Time `bun:",notnull"`

	Jokes        []Joke        `bun:"rel:has-many,join:id=user_id"`
	Comments     []Comment     `bun:"rel:has-many,join:id=user_id"`
	Likes        []Like        `bun:"rel:has-many,join:id=user_id"`
	CreatedRooms []Room        `bun:"rel:has-many,join:id=creator_user_id"`
	RoomMembers  []RoomMember  `bun:"rel:has-many,join:id=user_id"`
	RoomMessages []RoomMessage `bun:"rel:has-many,join:id=user_id"`
}

func NewUser(user model.User) User {
	return User{
		ID:             user.ID,
		Username:       user.Username,
		Email:          user.Email,
		HashedPassword: user.HashedPassword,
		Fullname:       user.Fullname,
		Status:         user.Status,
		CreatedAt:      user.CreatedAt,
		UpdatedAt:      user.UpdatedAt,
		DeletedAt:      user.DeletedAt,
	}
}

func (dto User) ToDomain() model.User {
	return model.User{
		ID:             dto.ID,
		Username:       dto.Username,
		Email:          dto.Email,
		HashedPassword: dto.HashedPassword,
		Fullname:       dto.Fullname,
		Status:         dto.Status,
		CreatedAt:      dto.CreatedAt,
		UpdatedAt:      dto.UpdatedAt,
		DeletedAt:      dto.DeletedAt,
	}
}
