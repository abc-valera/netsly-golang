package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

type RoomMember struct {
	ent.Schema
}

func (RoomMember) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Immutable(),
	}
}

func (RoomMember) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", Room.Type).
			Ref("members").
			Unique().
			Required(),
		edge.From("user", User.Type).
			Ref("rooms").
			Unique().
			Required(),
	}
}

func (RoomMember) Indexes() []ent.Index {
	return []ent.Index{
		index.Edges("room", "user").
			Unique(),
	}
}
