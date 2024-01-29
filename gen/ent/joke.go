// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/abc-valera/netsly-api-golang/gen/ent/joke"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
)

// Joke is the model entity for the Joke schema.
type Joke struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID string `json:"user_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Text holds the value of the "text" field.
	Text string `json:"text,omitempty"`
	// Explanation holds the value of the "explanation" field.
	Explanation string `json:"explanation,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the JokeQuery when eager-loading is set.
	Edges        JokeEdges `json:"edges"`
	user_jokes   *string
	selectValues sql.SelectValues
}

// JokeEdges holds the relations/edges for other nodes in the graph.
type JokeEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// Comments holds the value of the comments edge.
	Comments []*Comment `json:"comments,omitempty"`
	// Likes holds the value of the likes edge.
	Likes []*Like `json:"likes,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e JokeEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// CommentsOrErr returns the Comments value or an error if the edge
// was not loaded in eager-loading.
func (e JokeEdges) CommentsOrErr() ([]*Comment, error) {
	if e.loadedTypes[1] {
		return e.Comments, nil
	}
	return nil, &NotLoadedError{edge: "comments"}
}

// LikesOrErr returns the Likes value or an error if the edge
// was not loaded in eager-loading.
func (e JokeEdges) LikesOrErr() ([]*Like, error) {
	if e.loadedTypes[2] {
		return e.Likes, nil
	}
	return nil, &NotLoadedError{edge: "likes"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Joke) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case joke.FieldID, joke.FieldUserID, joke.FieldTitle, joke.FieldText, joke.FieldExplanation:
			values[i] = new(sql.NullString)
		case joke.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case joke.ForeignKeys[0]: // user_jokes
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Joke fields.
func (j *Joke) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case joke.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				j.ID = value.String
			}
		case joke.FieldUserID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				j.UserID = value.String
			}
		case joke.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				j.Title = value.String
			}
		case joke.FieldText:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field text", values[i])
			} else if value.Valid {
				j.Text = value.String
			}
		case joke.FieldExplanation:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field explanation", values[i])
			} else if value.Valid {
				j.Explanation = value.String
			}
		case joke.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				j.CreatedAt = value.Time
			}
		case joke.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_jokes", values[i])
			} else if value.Valid {
				j.user_jokes = new(string)
				*j.user_jokes = value.String
			}
		default:
			j.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Joke.
// This includes values selected through modifiers, order, etc.
func (j *Joke) Value(name string) (ent.Value, error) {
	return j.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the Joke entity.
func (j *Joke) QueryOwner() *UserQuery {
	return NewJokeClient(j.config).QueryOwner(j)
}

// QueryComments queries the "comments" edge of the Joke entity.
func (j *Joke) QueryComments() *CommentQuery {
	return NewJokeClient(j.config).QueryComments(j)
}

// QueryLikes queries the "likes" edge of the Joke entity.
func (j *Joke) QueryLikes() *LikeQuery {
	return NewJokeClient(j.config).QueryLikes(j)
}

// Update returns a builder for updating this Joke.
// Note that you need to call Joke.Unwrap() before calling this method if this Joke
// was returned from a transaction, and the transaction was committed or rolled back.
func (j *Joke) Update() *JokeUpdateOne {
	return NewJokeClient(j.config).UpdateOne(j)
}

// Unwrap unwraps the Joke entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (j *Joke) Unwrap() *Joke {
	_tx, ok := j.config.driver.(*txDriver)
	if !ok {
		panic("ent: Joke is not a transactional entity")
	}
	j.config.driver = _tx.drv
	return j
}

// String implements the fmt.Stringer.
func (j *Joke) String() string {
	var builder strings.Builder
	builder.WriteString("Joke(")
	builder.WriteString(fmt.Sprintf("id=%v, ", j.ID))
	builder.WriteString("user_id=")
	builder.WriteString(j.UserID)
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(j.Title)
	builder.WriteString(", ")
	builder.WriteString("text=")
	builder.WriteString(j.Text)
	builder.WriteString(", ")
	builder.WriteString("explanation=")
	builder.WriteString(j.Explanation)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(j.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// Jokes is a parsable slice of Joke.
type Jokes []*Joke
