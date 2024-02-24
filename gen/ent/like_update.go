// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc-valera/netsly-api-golang/gen/ent/joke"
	"github.com/abc-valera/netsly-api-golang/gen/ent/like"
	"github.com/abc-valera/netsly-api-golang/gen/ent/predicate"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
)

// LikeUpdate is the builder for updating Like entities.
type LikeUpdate struct {
	config
	hooks    []Hook
	mutation *LikeMutation
}

// Where appends a list predicates to the LikeUpdate builder.
func (lu *LikeUpdate) Where(ps ...predicate.Like) *LikeUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUserID sets the "user" edge to the User entity by ID.
func (lu *LikeUpdate) SetUserID(id string) *LikeUpdate {
	lu.mutation.SetUserID(id)
	return lu
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (lu *LikeUpdate) SetNillableUserID(id *string) *LikeUpdate {
	if id != nil {
		lu = lu.SetUserID(*id)
	}
	return lu
}

// SetUser sets the "user" edge to the User entity.
func (lu *LikeUpdate) SetUser(u *User) *LikeUpdate {
	return lu.SetUserID(u.ID)
}

// SetJokeID sets the "joke" edge to the Joke entity by ID.
func (lu *LikeUpdate) SetJokeID(id string) *LikeUpdate {
	lu.mutation.SetJokeID(id)
	return lu
}

// SetNillableJokeID sets the "joke" edge to the Joke entity by ID if the given value is not nil.
func (lu *LikeUpdate) SetNillableJokeID(id *string) *LikeUpdate {
	if id != nil {
		lu = lu.SetJokeID(*id)
	}
	return lu
}

// SetJoke sets the "joke" edge to the Joke entity.
func (lu *LikeUpdate) SetJoke(j *Joke) *LikeUpdate {
	return lu.SetJokeID(j.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (lu *LikeUpdate) Mutation() *LikeMutation {
	return lu.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (lu *LikeUpdate) ClearUser() *LikeUpdate {
	lu.mutation.ClearUser()
	return lu
}

// ClearJoke clears the "joke" edge to the Joke entity.
func (lu *LikeUpdate) ClearJoke() *LikeUpdate {
	lu.mutation.ClearJoke()
	return lu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *LikeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, lu.sqlSave, lu.mutation, lu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (lu *LikeUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *LikeUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *LikeUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lu *LikeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if lu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.JokeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.JokeTable,
			Columns: []string{like.JokeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.JokeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.JokeTable,
			Columns: []string{like.JokeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	lu.mutation.done = true
	return n, nil
}

// LikeUpdateOne is the builder for updating a single Like entity.
type LikeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *LikeMutation
}

// SetUserID sets the "user" edge to the User entity by ID.
func (luo *LikeUpdateOne) SetUserID(id string) *LikeUpdateOne {
	luo.mutation.SetUserID(id)
	return luo
}

// SetNillableUserID sets the "user" edge to the User entity by ID if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableUserID(id *string) *LikeUpdateOne {
	if id != nil {
		luo = luo.SetUserID(*id)
	}
	return luo
}

// SetUser sets the "user" edge to the User entity.
func (luo *LikeUpdateOne) SetUser(u *User) *LikeUpdateOne {
	return luo.SetUserID(u.ID)
}

// SetJokeID sets the "joke" edge to the Joke entity by ID.
func (luo *LikeUpdateOne) SetJokeID(id string) *LikeUpdateOne {
	luo.mutation.SetJokeID(id)
	return luo
}

// SetNillableJokeID sets the "joke" edge to the Joke entity by ID if the given value is not nil.
func (luo *LikeUpdateOne) SetNillableJokeID(id *string) *LikeUpdateOne {
	if id != nil {
		luo = luo.SetJokeID(*id)
	}
	return luo
}

// SetJoke sets the "joke" edge to the Joke entity.
func (luo *LikeUpdateOne) SetJoke(j *Joke) *LikeUpdateOne {
	return luo.SetJokeID(j.ID)
}

// Mutation returns the LikeMutation object of the builder.
func (luo *LikeUpdateOne) Mutation() *LikeMutation {
	return luo.mutation
}

// ClearUser clears the "user" edge to the User entity.
func (luo *LikeUpdateOne) ClearUser() *LikeUpdateOne {
	luo.mutation.ClearUser()
	return luo
}

// ClearJoke clears the "joke" edge to the Joke entity.
func (luo *LikeUpdateOne) ClearJoke() *LikeUpdateOne {
	luo.mutation.ClearJoke()
	return luo
}

// Where appends a list predicates to the LikeUpdate builder.
func (luo *LikeUpdateOne) Where(ps ...predicate.Like) *LikeUpdateOne {
	luo.mutation.Where(ps...)
	return luo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *LikeUpdateOne) Select(field string, fields ...string) *LikeUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated Like entity.
func (luo *LikeUpdateOne) Save(ctx context.Context) (*Like, error) {
	return withHooks(ctx, luo.sqlSave, luo.mutation, luo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (luo *LikeUpdateOne) SaveX(ctx context.Context) *Like {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *LikeUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *LikeUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (luo *LikeUpdateOne) sqlSave(ctx context.Context) (_node *Like, err error) {
	_spec := sqlgraph.NewUpdateSpec(like.Table, like.Columns, sqlgraph.NewFieldSpec(like.FieldID, field.TypeInt))
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Like.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, like.FieldID)
		for _, f := range fields {
			if !like.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != like.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if luo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.UserTable,
			Columns: []string{like.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.JokeCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.JokeTable,
			Columns: []string{like.JokeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.JokeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   like.JokeTable,
			Columns: []string{like.JokeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Like{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{like.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	luo.mutation.done = true
	return _node, nil
}
