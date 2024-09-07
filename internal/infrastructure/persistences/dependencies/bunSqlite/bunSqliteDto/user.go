package bunSqliteDto

import (
	"time"

	"github.com/abc-valera/netsly-golang/internal/domain/model"
	"github.com/abc-valera/netsly-golang/internal/domain/persistence/command"
	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users"`

	ID             string    `bun:"id,pk,type:uuid"`
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

func NewUserUpdate(ids model.User, req command.UserUpdateRequest) (User, []string) {
	user := User{
		ID: ids.ID,
	}
	var columns []string

	user.UpdatedAt = req.UpdatedAt
	columns = append(columns, "updated_at")
	if req.HashedPassword != nil {
		user.HashedPassword = *req.HashedPassword
		columns = append(columns, "hashed_password")
	}
	if req.Fullname != nil {
		user.Fullname = *req.Fullname
		columns = append(columns, "fullname")
	}
	if req.Status != nil {
		user.Status = *req.Status
		columns = append(columns, "status")
	}

	return user, columns
}

func (dto User) ToDomain() model.User {
	return model.User{
		ID:             dto.ID,
		Username:       dto.Username,
		Email:          dto.Email,
		HashedPassword: dto.HashedPassword,
		Fullname:       dto.Fullname,
		Status:         dto.Status,
		UpdatedAt:      dto.UpdatedAt,
		CreatedAt:      dto.CreatedAt,
		DeletedAt:      dto.DeletedAt,
	}
}

type Users []User

func NewUsers(users model.Users) Users {
	dtos := make(Users, 0, len(users))
	for _, user := range users {
		dtos = append(dtos, NewUser(user))
	}
	return dtos
}

func (dtos Users) ToDomain() model.Users {
	users := make(model.Users, 0, len(dtos))
	for _, user := range dtos {
		users = append(users, user.ToDomain())
	}
	return users
}
