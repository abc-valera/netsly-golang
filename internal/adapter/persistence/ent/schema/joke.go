package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Joke holds the schema definition for the Joke entity.
type Joke struct {
	ent.Schema
}

// Fields of the Joke.
func (Joke) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").
			NotEmpty().
			Unique().
			Immutable(),
		field.String("user_id").
			NotEmpty().
			Immutable(),
		field.String("title").
			NotEmpty(),
		field.String("text").
			NotEmpty(),
		field.String("explanation"),
		field.Time("created_at").
			Immutable(),
	}
}

// Edges of the Joke.
func (Joke) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).
			Ref("jokes").
			Unique().
			Required(),
		edge.To("comments", Comment.Type),
	}
}

func (Joke) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("user_id", "title").
			Unique(),
	}
}
