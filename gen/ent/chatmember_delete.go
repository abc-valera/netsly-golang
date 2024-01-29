// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc-valera/netsly-api-golang/gen/ent/chatmember"
	"github.com/abc-valera/netsly-api-golang/gen/ent/predicate"
)

// ChatMemberDelete is the builder for deleting a ChatMember entity.
type ChatMemberDelete struct {
	config
	hooks    []Hook
	mutation *ChatMemberMutation
}

// Where appends a list predicates to the ChatMemberDelete builder.
func (cmd *ChatMemberDelete) Where(ps ...predicate.ChatMember) *ChatMemberDelete {
	cmd.mutation.Where(ps...)
	return cmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (cmd *ChatMemberDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, cmd.sqlExec, cmd.mutation, cmd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (cmd *ChatMemberDelete) ExecX(ctx context.Context) int {
	n, err := cmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (cmd *ChatMemberDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(chatmember.Table, sqlgraph.NewFieldSpec(chatmember.FieldID, field.TypeInt))
	if ps := cmd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, cmd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	cmd.mutation.done = true
	return affected, err
}

// ChatMemberDeleteOne is the builder for deleting a single ChatMember entity.
type ChatMemberDeleteOne struct {
	cmd *ChatMemberDelete
}

// Where appends a list predicates to the ChatMemberDelete builder.
func (cmdo *ChatMemberDeleteOne) Where(ps ...predicate.ChatMember) *ChatMemberDeleteOne {
	cmdo.cmd.mutation.Where(ps...)
	return cmdo
}

// Exec executes the deletion query.
func (cmdo *ChatMemberDeleteOne) Exec(ctx context.Context) error {
	n, err := cmdo.cmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{chatmember.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (cmdo *ChatMemberDeleteOne) ExecX(ctx context.Context) {
	if err := cmdo.Exec(ctx); err != nil {
		panic(err)
	}
}
