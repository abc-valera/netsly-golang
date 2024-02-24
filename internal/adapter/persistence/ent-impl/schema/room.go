package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type Room struct {
	ent.Schema
}

func (Room) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("name").
			NotEmpty().
			Unique(),
		field.String("description"),
		field.Time("created_at").
			Immutable(),
	}
}

func (Room) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", RoomMember.Type),
		edge.To("messages", RoomMessage.Type),
	}
}
