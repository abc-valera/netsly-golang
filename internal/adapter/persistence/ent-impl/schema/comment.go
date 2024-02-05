package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the User entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("user_id").
			NotEmpty().
			Immutable(),
		field.String("joke_id").
			NotEmpty().
			Immutable(),
		field.String("text").
			NotEmpty(),
		field.Time("created_at").
			Immutable(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("comments").
			Unique(),
		edge.From("commented_joke", Joke.Type).
			Ref("comments").
			Unique(),
	}
}
