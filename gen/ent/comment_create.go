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
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
)

// CommentCreate is the builder for creating a Comment entity.
type CommentCreate struct {
	config
	mutation *CommentMutation
	hooks    []Hook
}

// SetUserID sets the "user_id" field.
func (cc *CommentCreate) SetUserID(s string) *CommentCreate {
	cc.mutation.SetUserID(s)
	return cc
}

// SetJokeID sets the "joke_id" field.
func (cc *CommentCreate) SetJokeID(s string) *CommentCreate {
	cc.mutation.SetJokeID(s)
	return cc
}

// SetText sets the "text" field.
func (cc *CommentCreate) SetText(s string) *CommentCreate {
	cc.mutation.SetText(s)
	return cc
}

// SetCreatedAt sets the "created_at" field.
func (cc *CommentCreate) SetCreatedAt(t time.Time) *CommentCreate {
	cc.mutation.SetCreatedAt(t)
	return cc
}

// SetID sets the "id" field.
func (cc *CommentCreate) SetID(s string) *CommentCreate {
	cc.mutation.SetID(s)
	return cc
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (cc *CommentCreate) SetOwnerID(id string) *CommentCreate {
	cc.mutation.SetOwnerID(id)
	return cc
}

// SetOwner sets the "owner" edge to the User entity.
func (cc *CommentCreate) SetOwner(u *User) *CommentCreate {
	return cc.SetOwnerID(u.ID)
}

// SetCommentedJokeID sets the "commented_joke" edge to the Joke entity by ID.
func (cc *CommentCreate) SetCommentedJokeID(id string) *CommentCreate {
	cc.mutation.SetCommentedJokeID(id)
	return cc
}

// SetCommentedJoke sets the "commented_joke" edge to the Joke entity.
func (cc *CommentCreate) SetCommentedJoke(j *Joke) *CommentCreate {
	return cc.SetCommentedJokeID(j.ID)
}

// Mutation returns the CommentMutation object of the builder.
func (cc *CommentCreate) Mutation() *CommentMutation {
	return cc.mutation
}

// Save creates the Comment in the database.
func (cc *CommentCreate) Save(ctx context.Context) (*Comment, error) {
	return withHooks(ctx, cc.sqlSave, cc.mutation, cc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (cc *CommentCreate) SaveX(ctx context.Context) *Comment {
	v, err := cc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (cc *CommentCreate) Exec(ctx context.Context) error {
	_, err := cc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cc *CommentCreate) ExecX(ctx context.Context) {
	if err := cc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cc *CommentCreate) check() error {
	if _, ok := cc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "Comment.user_id"`)}
	}
	if v, ok := cc.mutation.UserID(); ok {
		if err := comment.UserIDValidator(v); err != nil {
			return &ValidationError{Name: "user_id", err: fmt.Errorf(`ent: validator failed for field "Comment.user_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.JokeID(); !ok {
		return &ValidationError{Name: "joke_id", err: errors.New(`ent: missing required field "Comment.joke_id"`)}
	}
	if v, ok := cc.mutation.JokeID(); ok {
		if err := comment.JokeIDValidator(v); err != nil {
			return &ValidationError{Name: "joke_id", err: fmt.Errorf(`ent: validator failed for field "Comment.joke_id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "Comment.text"`)}
	}
	if v, ok := cc.mutation.Text(); ok {
		if err := comment.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Comment.text": %w`, err)}
		}
	}
	if _, ok := cc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Comment.created_at"`)}
	}
	if v, ok := cc.mutation.ID(); ok {
		if err := comment.IDValidator(v); err != nil {
			return &ValidationError{Name: "id", err: fmt.Errorf(`ent: validator failed for field "Comment.id": %w`, err)}
		}
	}
	if _, ok := cc.mutation.OwnerID(); !ok {
		return &ValidationError{Name: "owner", err: errors.New(`ent: missing required edge "Comment.owner"`)}
	}
	if _, ok := cc.mutation.CommentedJokeID(); !ok {
		return &ValidationError{Name: "commented_joke", err: errors.New(`ent: missing required edge "Comment.commented_joke"`)}
	}
	return nil
}

func (cc *CommentCreate) sqlSave(ctx context.Context) (*Comment, error) {
	if err := cc.check(); err != nil {
		return nil, err
	}
	_node, _spec := cc.createSpec()
	if err := sqlgraph.CreateNode(ctx, cc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Comment.ID type: %T", _spec.ID.Value)
		}
	}
	cc.mutation.id = &_node.ID
	cc.mutation.done = true
	return _node, nil
}

func (cc *CommentCreate) createSpec() (*Comment, *sqlgraph.CreateSpec) {
	var (
		_node = &Comment{config: cc.config}
		_spec = sqlgraph.NewCreateSpec(comment.Table, sqlgraph.NewFieldSpec(comment.FieldID, field.TypeString))
	)
	if id, ok := cc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := cc.mutation.UserID(); ok {
		_spec.SetField(comment.FieldUserID, field.TypeString, value)
		_node.UserID = value
	}
	if value, ok := cc.mutation.JokeID(); ok {
		_spec.SetField(comment.FieldJokeID, field.TypeString, value)
		_node.JokeID = value
	}
	if value, ok := cc.mutation.Text(); ok {
		_spec.SetField(comment.FieldText, field.TypeString, value)
		_node.Text = value
	}
	if value, ok := cc.mutation.CreatedAt(); ok {
		_spec.SetField(comment.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if nodes := cc.mutation.OwnerIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.OwnerTable,
			Columns: []string{comment.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.user_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := cc.mutation.CommentedJokeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   comment.CommentedJokeTable,
			Columns: []string{comment.CommentedJokeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.joke_comments = &nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// CommentCreateBulk is the builder for creating many Comment entities in bulk.
type CommentCreateBulk struct {
	config
	err      error
	builders []*CommentCreate
}

// Save creates the Comment entities in the database.
func (ccb *CommentCreateBulk) Save(ctx context.Context) ([]*Comment, error) {
	if ccb.err != nil {
		return nil, ccb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(ccb.builders))
	nodes := make([]*Comment, len(ccb.builders))
	mutators := make([]Mutator, len(ccb.builders))
	for i := range ccb.builders {
		func(i int, root context.Context) {
			builder := ccb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*CommentMutation)
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
					_, err = mutators[i+1].Mutate(root, ccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, ccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ccb *CommentCreateBulk) SaveX(ctx context.Context) []*Comment {
	v, err := ccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ccb *CommentCreateBulk) Exec(ctx context.Context) error {
	_, err := ccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ccb *CommentCreateBulk) ExecX(ctx context.Context) {
	if err := ccb.Exec(ctx); err != nil {
		panic(err)
	}
}
