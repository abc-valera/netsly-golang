package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Like holds the schema definition for the Like entity.
type Like struct {
	ent.Schema
}

// Fields of the Like.
func (Like) Fields() []ent.Field {
	return []ent.Field{
		field.String("user_id").
			NotEmpty().
			Immutable(),
		field.String("joke_id").
			NotEmpty().
			Immutable(),
		field.Time("created_at").
			Immutable(),
	}
}

// Edges of the Like.
func (Like) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("likes").
			Unique().
			Required(),
		edge.From("liked_joke", Joke.Type).
			Ref("likes").
			Unique().
			Required(),
	}
}

// Indexes of the Like.
func (Like) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "joke_id").
			Unique(),
	}

}
