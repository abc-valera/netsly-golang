package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ChatRoom struct {
	ent.Schema
}

func (ChatRoom) Fields() []ent.Field {
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

func (ChatRoom) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("members", ChatMember.Type),
		edge.To("messages", ChatMessage.Type),
	}
}
