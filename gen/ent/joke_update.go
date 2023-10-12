// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/abc-valera/flugo-api-golang/gen/ent/comment"
	"github.com/abc-valera/flugo-api-golang/gen/ent/joke"
	"github.com/abc-valera/flugo-api-golang/gen/ent/like"
	"github.com/abc-valera/flugo-api-golang/gen/ent/predicate"
	"github.com/abc-valera/flugo-api-golang/gen/ent/user"
)

// JokeUpdate is the builder for updating Joke entities.
type JokeUpdate struct {
	config
	hooks    []Hook
	mutation *JokeMutation
}

// Where appends a list predicates to the JokeUpdate builder.
func (ju *JokeUpdate) Where(ps ...predicate.Joke) *JokeUpdate {
	ju.mutation.Where(ps...)
	return ju
}

// SetTitle sets the "title" field.
func (ju *JokeUpdate) SetTitle(s string) *JokeUpdate {
	ju.mutation.SetTitle(s)
	return ju
}

// SetText sets the "text" field.
func (ju *JokeUpdate) SetText(s string) *JokeUpdate {
	ju.mutation.SetText(s)
	return ju
}

// SetExplanation sets the "explanation" field.
func (ju *JokeUpdate) SetExplanation(s string) *JokeUpdate {
	ju.mutation.SetExplanation(s)
	return ju
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (ju *JokeUpdate) SetOwnerID(id string) *JokeUpdate {
	ju.mutation.SetOwnerID(id)
	return ju
}

// SetOwner sets the "owner" edge to the User entity.
func (ju *JokeUpdate) SetOwner(u *User) *JokeUpdate {
	return ju.SetOwnerID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (ju *JokeUpdate) AddCommentIDs(ids ...string) *JokeUpdate {
	ju.mutation.AddCommentIDs(ids...)
	return ju
}

// AddComments adds the "comments" edges to the Comment entity.
func (ju *JokeUpdate) AddComments(c ...*Comment) *JokeUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ju.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (ju *JokeUpdate) AddLikeIDs(ids ...int) *JokeUpdate {
	ju.mutation.AddLikeIDs(ids...)
	return ju
}

// AddLikes adds the "likes" edges to the Like entity.
func (ju *JokeUpdate) AddLikes(l ...*Like) *JokeUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ju.AddLikeIDs(ids...)
}

// Mutation returns the JokeMutation object of the builder.
func (ju *JokeUpdate) Mutation() *JokeMutation {
	return ju.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (ju *JokeUpdate) ClearOwner() *JokeUpdate {
	ju.mutation.ClearOwner()
	return ju
}

// ClearComments clears all "comments" edges to the Comment entity.
func (ju *JokeUpdate) ClearComments() *JokeUpdate {
	ju.mutation.ClearComments()
	return ju
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (ju *JokeUpdate) RemoveCommentIDs(ids ...string) *JokeUpdate {
	ju.mutation.RemoveCommentIDs(ids...)
	return ju
}

// RemoveComments removes "comments" edges to Comment entities.
func (ju *JokeUpdate) RemoveComments(c ...*Comment) *JokeUpdate {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return ju.RemoveCommentIDs(ids...)
}

// ClearLikes clears all "likes" edges to the Like entity.
func (ju *JokeUpdate) ClearLikes() *JokeUpdate {
	ju.mutation.ClearLikes()
	return ju
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (ju *JokeUpdate) RemoveLikeIDs(ids ...int) *JokeUpdate {
	ju.mutation.RemoveLikeIDs(ids...)
	return ju
}

// RemoveLikes removes "likes" edges to Like entities.
func (ju *JokeUpdate) RemoveLikes(l ...*Like) *JokeUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ju.RemoveLikeIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ju *JokeUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ju.sqlSave, ju.mutation, ju.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ju *JokeUpdate) SaveX(ctx context.Context) int {
	affected, err := ju.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ju *JokeUpdate) Exec(ctx context.Context) error {
	_, err := ju.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ju *JokeUpdate) ExecX(ctx context.Context) {
	if err := ju.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ju *JokeUpdate) check() error {
	if v, ok := ju.mutation.Title(); ok {
		if err := joke.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Joke.title": %w`, err)}
		}
	}
	if v, ok := ju.mutation.Text(); ok {
		if err := joke.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Joke.text": %w`, err)}
		}
	}
	if _, ok := ju.mutation.OwnerID(); ju.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Joke.owner"`)
	}
	return nil
}

func (ju *JokeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ju.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(joke.Table, joke.Columns, sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString))
	if ps := ju.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ju.mutation.Title(); ok {
		_spec.SetField(joke.FieldTitle, field.TypeString, value)
	}
	if value, ok := ju.mutation.Text(); ok {
		_spec.SetField(joke.FieldText, field.TypeString, value)
	}
	if value, ok := ju.mutation.Explanation(); ok {
		_spec.SetField(joke.FieldExplanation, field.TypeString, value)
	}
	if ju.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ju.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !ju.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.CommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ju.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.RemovedLikesIDs(); len(nodes) > 0 && !ju.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ju.mutation.LikesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ju.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{joke.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ju.mutation.done = true
	return n, nil
}

// JokeUpdateOne is the builder for updating a single Joke entity.
type JokeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *JokeMutation
}

// SetTitle sets the "title" field.
func (juo *JokeUpdateOne) SetTitle(s string) *JokeUpdateOne {
	juo.mutation.SetTitle(s)
	return juo
}

// SetText sets the "text" field.
func (juo *JokeUpdateOne) SetText(s string) *JokeUpdateOne {
	juo.mutation.SetText(s)
	return juo
}

// SetExplanation sets the "explanation" field.
func (juo *JokeUpdateOne) SetExplanation(s string) *JokeUpdateOne {
	juo.mutation.SetExplanation(s)
	return juo
}

// SetOwnerID sets the "owner" edge to the User entity by ID.
func (juo *JokeUpdateOne) SetOwnerID(id string) *JokeUpdateOne {
	juo.mutation.SetOwnerID(id)
	return juo
}

// SetOwner sets the "owner" edge to the User entity.
func (juo *JokeUpdateOne) SetOwner(u *User) *JokeUpdateOne {
	return juo.SetOwnerID(u.ID)
}

// AddCommentIDs adds the "comments" edge to the Comment entity by IDs.
func (juo *JokeUpdateOne) AddCommentIDs(ids ...string) *JokeUpdateOne {
	juo.mutation.AddCommentIDs(ids...)
	return juo
}

// AddComments adds the "comments" edges to the Comment entity.
func (juo *JokeUpdateOne) AddComments(c ...*Comment) *JokeUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return juo.AddCommentIDs(ids...)
}

// AddLikeIDs adds the "likes" edge to the Like entity by IDs.
func (juo *JokeUpdateOne) AddLikeIDs(ids ...int) *JokeUpdateOne {
	juo.mutation.AddLikeIDs(ids...)
	return juo
}

// AddLikes adds the "likes" edges to the Like entity.
func (juo *JokeUpdateOne) AddLikes(l ...*Like) *JokeUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return juo.AddLikeIDs(ids...)
}

// Mutation returns the JokeMutation object of the builder.
func (juo *JokeUpdateOne) Mutation() *JokeMutation {
	return juo.mutation
}

// ClearOwner clears the "owner" edge to the User entity.
func (juo *JokeUpdateOne) ClearOwner() *JokeUpdateOne {
	juo.mutation.ClearOwner()
	return juo
}

// ClearComments clears all "comments" edges to the Comment entity.
func (juo *JokeUpdateOne) ClearComments() *JokeUpdateOne {
	juo.mutation.ClearComments()
	return juo
}

// RemoveCommentIDs removes the "comments" edge to Comment entities by IDs.
func (juo *JokeUpdateOne) RemoveCommentIDs(ids ...string) *JokeUpdateOne {
	juo.mutation.RemoveCommentIDs(ids...)
	return juo
}

// RemoveComments removes "comments" edges to Comment entities.
func (juo *JokeUpdateOne) RemoveComments(c ...*Comment) *JokeUpdateOne {
	ids := make([]string, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return juo.RemoveCommentIDs(ids...)
}

// ClearLikes clears all "likes" edges to the Like entity.
func (juo *JokeUpdateOne) ClearLikes() *JokeUpdateOne {
	juo.mutation.ClearLikes()
	return juo
}

// RemoveLikeIDs removes the "likes" edge to Like entities by IDs.
func (juo *JokeUpdateOne) RemoveLikeIDs(ids ...int) *JokeUpdateOne {
	juo.mutation.RemoveLikeIDs(ids...)
	return juo
}

// RemoveLikes removes "likes" edges to Like entities.
func (juo *JokeUpdateOne) RemoveLikes(l ...*Like) *JokeUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return juo.RemoveLikeIDs(ids...)
}

// Where appends a list predicates to the JokeUpdate builder.
func (juo *JokeUpdateOne) Where(ps ...predicate.Joke) *JokeUpdateOne {
	juo.mutation.Where(ps...)
	return juo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (juo *JokeUpdateOne) Select(field string, fields ...string) *JokeUpdateOne {
	juo.fields = append([]string{field}, fields...)
	return juo
}

// Save executes the query and returns the updated Joke entity.
func (juo *JokeUpdateOne) Save(ctx context.Context) (*Joke, error) {
	return withHooks(ctx, juo.sqlSave, juo.mutation, juo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (juo *JokeUpdateOne) SaveX(ctx context.Context) *Joke {
	node, err := juo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (juo *JokeUpdateOne) Exec(ctx context.Context) error {
	_, err := juo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (juo *JokeUpdateOne) ExecX(ctx context.Context) {
	if err := juo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (juo *JokeUpdateOne) check() error {
	if v, ok := juo.mutation.Title(); ok {
		if err := joke.TitleValidator(v); err != nil {
			return &ValidationError{Name: "title", err: fmt.Errorf(`ent: validator failed for field "Joke.title": %w`, err)}
		}
	}
	if v, ok := juo.mutation.Text(); ok {
		if err := joke.TextValidator(v); err != nil {
			return &ValidationError{Name: "text", err: fmt.Errorf(`ent: validator failed for field "Joke.text": %w`, err)}
		}
	}
	if _, ok := juo.mutation.OwnerID(); juo.mutation.OwnerCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "Joke.owner"`)
	}
	return nil
}

func (juo *JokeUpdateOne) sqlSave(ctx context.Context) (_node *Joke, err error) {
	if err := juo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(joke.Table, joke.Columns, sqlgraph.NewFieldSpec(joke.FieldID, field.TypeString))
	id, ok := juo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Joke.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := juo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, joke.FieldID)
		for _, f := range fields {
			if !joke.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != joke.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := juo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := juo.mutation.Title(); ok {
		_spec.SetField(joke.FieldTitle, field.TypeString, value)
	}
	if value, ok := juo.mutation.Text(); ok {
		_spec.SetField(joke.FieldText, field.TypeString, value)
	}
	if value, ok := juo.mutation.Explanation(); ok {
		_spec.SetField(joke.FieldExplanation, field.TypeString, value)
	}
	if juo.mutation.OwnerCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.OwnerIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if juo.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.RemovedCommentsIDs(); len(nodes) > 0 && !juo.mutation.CommentsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.CommentsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if juo.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.RemovedLikesIDs(); len(nodes) > 0 && !juo.mutation.LikesCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := juo.mutation.LikesIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Joke{config: juo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, juo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{joke.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	juo.mutation.done = true
	return _node, nil
}
