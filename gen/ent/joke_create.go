// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc-valera/flugo-api-golang/gen/ent/comment"
	"github.com/abc-valera/flugo-api-golang/gen/ent/joke"
	"github.com/abc-valera/flugo-api-golang/gen/ent/like"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
)

// JokeCreate is the builder for creating a Joke entity.
type JokeCreate struct {
	config
	mutation *JokeMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (jc *JokeCreate) SetUserID(s string) *JokeCreate {
	jc.mutation.SetUserID(s)
	return jc
}

// SetTitle sets the "title" field.
func (jc *JokeCreate) SetTitle(s string) *JokeCreate {
	jc.mutation.SetTitle(s)
	return jc
}

// SetText sets the "text" field.
func (jc *JokeCreate) SetText(s string) *JokeCreate {
	jc.mutation.SetText(s)
	return jc
}

// SetExplanation sets the "explanation" field.
func (jc *JokeCreate) SetExplanation(s string) *JokeCreate {
	jc.mutation.SetExplanation(s)
	return jc
}

// SetCreatedAt sets the "created_at" field.
func (jc *JokeCreate) SetCreatedAt(t time.Time) *JokeCreate {
	jc.mutation.SetCreatedAt(t)
	return jc
}

// SetID sets the "id" field.
func (jc *JokeCreate) SetID(s string) *JokeCreate {
	jc.mutation.SetID(s)
	return jc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (jc *JokeCreate) SetOwnerID(id string) *JokeCreate {
	jc.mutation.SetOwnerID(id)
	return jc
}

// SetOwner sets the "owner" edge to the User entity.
func (jc *JokeCreate) SetOwner(u *User) *JokeCreate {
	return jc.SetOwnerID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (jc *JokeCreate) AddCommentIDs(ids ...string) *JokeCreate {
	jc.mutation.AddCommentIDs(ids...)
	return jc
}

// AddComments adds the "comments" edges to the Comment entity.
func (jc *JokeCreate) AddComments(c ...*Comment) *JokeCreate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return jc.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (jc *JokeCreate) AddLikeIDs(ids ...int) *JokeCreate {
	jc.mutation.AddLikeIDs(ids...)
	return jc
}

// AddLikes adds the "likes" edges to the Like entity.
func (jc *JokeCreate) AddLikes(l ...*Like) *JokeCreate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return jc.AddLikeIDs(ids...)
}

// Mutation returns the JokeMutation object of the builder.
func (jc *JokeCreate) Mutation() *JokeMutation {
	return jc.mutation
}

// Save creates the Joke in the database.
func (jc *JokeCreate) Save(ctx context.Context) (*Joke, error) {
	return withHooks(ctx, jc.sqlSave, jc.mutation, jc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (jc *JokeCreate) SaveX(ctx context.Context) *Joke {
	v, err := jc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jc *JokeCreate) Exec(ctx context.Context) error {
	_, err := jc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jc *JokeCreate) ExecX(ctx context.Context) {
	if err := jc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (jc *JokeCreate) check() error {
	if _, ok := jc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Joke.user_id"`)}
	}
	if v, ok := jc.mutation.UserID(); ok {
		if err := joke.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Joke.user_id": %w`, err)}
		}
	}
	if _, ok := jc.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Joke.title"`)}
	}
	if v, ok := jc.mutation.Title(); ok {
		if err := joke.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Joke.title": %w`, err)}
		}
	}
	if _, ok := jc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Joke.text"`)}
	}
	if v, ok := jc.mutation.Text(); ok {
		if err := joke.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Joke.text": %w`, err)}
		}
	}
	if _, ok := jc.mutation.Explanation(); !ok {
		return &ValidationError{Name: "explanation", err: errors.New(`ent: missing required field "Joke.explanation"`)}
	}
	if _, ok := jc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Joke.created_at"`)}
	}
	if v, ok := jc.mutation.ID(); ok {
		if err := joke.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Joke.id": %w`, err)}
		}
	}
	if _, ok := jc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Joke.owner"`)}
	}
	return nil
}

func (jc *JokeCreate) sqlSave(ctx context.Context) (*Joke, error) {
	if err := jc.check(); err != nil {
		return nil, err
	}
	_node, _spec := jc.createSpec()
	if err := sqlgraph.CreateNode(ctx, jc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Joke.ID type: %T", _spec.ID.Value)
		}
	}
	jc.mutation.id = &_node.ID
	jc.mutation.done = true
	return _node, nil
}

func (jc *JokeCreate) createSpec() (*Joke, *sqlgraph.CreateSpec) {
	var (
		_node = &Joke{config: jc.config}
		_spec = sqlgraph.NewCreateSpec(joke.Table, sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString))
	)
	if id, ok := jc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := jc.mutation.UserID(); ok {
		_spec.SetField(joke.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := jc.mutation.Title(); ok {
		_spec.SetField(joke.FieldTitle, field.TypeString, value)
		_node.Title = value
	}
	if value, ok := jc.mutation.Text(); ok {
		_spec.SetField(joke.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := jc.mutation.Explanation(); ok {
		_spec.SetField(joke.FieldExplanation, field.TypeString, value)
		_node.Explanation = value
	}
	if value, ok := jc.mutation.CreatedAt(); ok {
		_spec.SetField(joke.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := jc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   joke.OwnerTable,
			Columns: []string{joke.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_jokes = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.mutation.CommentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joke.CommentsTable,
			Columns: []string{joke.CommentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := jc.mutation.LikesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   joke.LikesTable,
			Columns: []string{joke.LikesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// JokeCreateBulk is the builder for creating many Joke entities in bulk.
type JokeCreateBulk struct {
	config
	err      error
	builders []*JokeCreate
}

// Save creates the Joke entities in the database.
func (jcb *JokeCreateBulk) Save(ctx context.Context) ([]*Joke, error) {
	if jcb.err != nil {
		return nil, jcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(jcb.builders))
	nodes := make([]*Joke, len(jcb.builders))
	mutators := make([]Mutator, len(jcb.builders))
	for i := range jcb.builders {
		func(i int, root context.Context) {
			builder := jcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*JokeMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, jcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, jcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, jcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (jcb *JokeCreateBulk) SaveX(ctx context.Context) []*Joke {
	v, err := jcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (jcb *JokeCreateBulk) Exec(ctx context.Context) error {
	_, err := jcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (jcb *JokeCreateBulk) ExecX(ctx context.Context) {
	if err := jcb.Exec(ctx); err != nil {
		panic(err)
	}
}
