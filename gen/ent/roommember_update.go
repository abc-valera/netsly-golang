// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc-valera/netsly-api-golang/gen/ent/predicate"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommember"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
)

// RoomMemberUpdate is the builder for updating RoomMember entities.
type RoomMemberUpdate struct {
	config
	hooks    []Hook
	mutation *RoomMemberMutation
}

// Where appends a list predicates to the RoomMemberUpdate builder.
func (rmu *RoomMemberUpdate) Where(ps ...predicate.RoomMember) *RoomMemberUpdate {
	rmu.mutation.Where(ps...)
	return rmu
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (rmu *RoomMemberUpdate) SetRoomID(id string) *RoomMemberUpdate {
	rmu.mutation.SetRoomID(id)
	return rmu
}

// SetRoom sets the "room" edge to the Room entity.
func (rmu *RoomMemberUpdate) SetRoom(r *Room) *RoomMemberUpdate {
	return rmu.SetRoomID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rmu *RoomMemberUpdate) SetUserID(id string) *RoomMemberUpdate {
	rmu.mutation.SetUserID(id)
	return rmu
}

// SetUser sets the "user" edge to the User entity.
func (rmu *RoomMemberUpdate) SetUser(u *User) *RoomMemberUpdate {
	return rmu.SetUserID(u.ID)
}

// Mutation returns the RoomMemberMutation object of the builder.
func (rmu *RoomMemberUpdate) Mutation() *RoomMemberMutation {
	return rmu.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (rmu *RoomMemberUpdate) ClearRoom() *RoomMemberUpdate {
	rmu.mutation.ClearRoom()
	return rmu
}

// ClearUser clears the "user" edge to the User entity.
func (rmu *RoomMemberUpdate) ClearUser() *RoomMemberUpdate {
	rmu.mutation.ClearUser()
	return rmu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rmu *RoomMemberUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, rmu.sqlSave, rmu.mutation, rmu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmu *RoomMemberUpdate) SaveX(ctx context.Context) int {
	affected, err := rmu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rmu *RoomMemberUpdate) Exec(ctx context.Context) error {
	_, err := rmu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmu *RoomMemberUpdate) ExecX(ctx context.Context) {
	if err := rmu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmu *RoomMemberUpdate) check() error {
	if _, ok := rmu.mutation.RoomID(); rmu.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoomMember.room"`)
	}
	if _, ok := rmu.mutation.UserID(); rmu.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoomMember.user"`)
	}
	return nil
}

func (rmu *RoomMemberUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := rmu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(roommember.Table, roommember.Columns, sqlgraph.NewFieldSpec(roommember.FieldID, field.TypeInt))
	if ps := rmu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rmu.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.RoomTable,
			Columns: []string{roommember.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.RoomTable,
			Columns: []string{roommember.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmu.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.UserTable,
			Columns: []string{roommember.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmu.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.UserTable,
			Columns: []string{roommember.UserColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, rmu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roommember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	rmu.mutation.done = true
	return n, nil
}

// RoomMemberUpdateOne is the builder for updating a single RoomMember entity.
type RoomMemberUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RoomMemberMutation
}

// SetRoomID sets the "room" edge to the Room entity by ID.
func (rmuo *RoomMemberUpdateOne) SetRoomID(id string) *RoomMemberUpdateOne {
	rmuo.mutation.SetRoomID(id)
	return rmuo
}

// SetRoom sets the "room" edge to the Room entity.
func (rmuo *RoomMemberUpdateOne) SetRoom(r *Room) *RoomMemberUpdateOne {
	return rmuo.SetRoomID(r.ID)
}

// SetUserID sets the "user" edge to the User entity by ID.
func (rmuo *RoomMemberUpdateOne) SetUserID(id string) *RoomMemberUpdateOne {
	rmuo.mutation.SetUserID(id)
	return rmuo
}

// SetUser sets the "user" edge to the User entity.
func (rmuo *RoomMemberUpdateOne) SetUser(u *User) *RoomMemberUpdateOne {
	return rmuo.SetUserID(u.ID)
}

// Mutation returns the RoomMemberMutation object of the builder.
func (rmuo *RoomMemberUpdateOne) Mutation() *RoomMemberMutation {
	return rmuo.mutation
}

// ClearRoom clears the "room" edge to the Room entity.
func (rmuo *RoomMemberUpdateOne) ClearRoom() *RoomMemberUpdateOne {
	rmuo.mutation.ClearRoom()
	return rmuo
}

// ClearUser clears the "user" edge to the User entity.
func (rmuo *RoomMemberUpdateOne) ClearUser() *RoomMemberUpdateOne {
	rmuo.mutation.ClearUser()
	return rmuo
}

// Where appends a list predicates to the RoomMemberUpdate builder.
func (rmuo *RoomMemberUpdateOne) Where(ps ...predicate.RoomMember) *RoomMemberUpdateOne {
	rmuo.mutation.Where(ps...)
	return rmuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rmuo *RoomMemberUpdateOne) Select(field string, fields ...string) *RoomMemberUpdateOne {
	rmuo.fields = append([]string{field}, fields...)
	return rmuo
}

// Save executes the query and returns the updated RoomMember entity.
func (rmuo *RoomMemberUpdateOne) Save(ctx context.Context) (*RoomMember, error) {
	return withHooks(ctx, rmuo.sqlSave, rmuo.mutation, rmuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (rmuo *RoomMemberUpdateOne) SaveX(ctx context.Context) *RoomMember {
	node, err := rmuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rmuo *RoomMemberUpdateOne) Exec(ctx context.Context) error {
	_, err := rmuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rmuo *RoomMemberUpdateOne) ExecX(ctx context.Context) {
	if err := rmuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rmuo *RoomMemberUpdateOne) check() error {
	if _, ok := rmuo.mutation.RoomID(); rmuo.mutation.RoomCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoomMember.room"`)
	}
	if _, ok := rmuo.mutation.UserID(); rmuo.mutation.UserCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "RoomMember.user"`)
	}
	return nil
}

func (rmuo *RoomMemberUpdateOne) sqlSave(ctx context.Context) (_node *RoomMember, err error) {
	if err := rmuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(roommember.Table, roommember.Columns, sqlgraph.NewFieldSpec(roommember.FieldID, field.TypeInt))
	id, ok := rmuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RoomMember.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rmuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, roommember.FieldID)
		for _, f := range fields {
			if !roommember.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != roommember.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rmuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if rmuo.mutation.RoomCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.RoomTable,
			Columns: []string{roommember.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.RoomIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.RoomTable,
			Columns: []string{roommember.RoomColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(room.FieldID, field.TypeString),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if rmuo.mutation.UserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.UserTable,
			Columns: []string{roommember.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(user.FieldID, field.TypeString),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := rmuo.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   roommember.UserTable,
			Columns: []string{roommember.UserColumn},
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
	_node = &RoomMember{config: rmuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rmuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{roommember.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	rmuo.mutation.done = true
	return _node, nil
}