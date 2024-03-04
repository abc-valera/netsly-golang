package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type RoomMessage struct {
	ent.Schema
}

func (RoomMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("text").
			NotEmpty(),
		field.Time("created_at").
			Immutable(),
	}
}

func (RoomMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("messages").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("room_messages").
			Unique().
			Required(),
	}
}
