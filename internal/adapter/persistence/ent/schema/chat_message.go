package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type ChatMessage struct {
	ent.Schema
}

func (ChatMessage) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("chat_room_id").
			NotEmpty().
			Immutable(),
		field.String("user_id").
			NotEmpty().
			Immutable(),
		field.String("text").
			NotEmpty(),
		field.Time("created_at").
			Immutable(),
	}
}

func (ChatMessage) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("room", ChatRoom.Type).
			Ref("messages").
			Unique().
			Required(),
		edge.From("owner", User.Type).
			Ref("chat_messages").
			Unique().
			Required(),
	}
}
