package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type ChatMember struct {
	ent.Schema
}

func (ChatMember) Fields() []ent.Field {
	return []ent.Field{
		field.String("chat_room_id").
			NotEmpty().
			Immutable(),
		field.String("user_id").
			NotEmpty().
			Immutable(),
		field.Time("created_at").
			Immutable(),
	}
}

func (ChatMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", ChatRoom.Type).
			Ref("members").
			Unique().
			Required(),
		edge.From("member", User.Type).
			Ref("chat_rooms").
			Unique().
			Required(),
	}
}

func (ChatMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("chat_room_id", "user_id").
			Unique(),
	}
}
