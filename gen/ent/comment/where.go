// Code generated by ent, DO NOT EDIT.

package comment

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc-valera/netsly-api-golang/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserID, v))
}

// JokeID applies equality check predicate on the "joke_id" field. It's identical to JokeIDEQ.
func JokeID(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldJokeID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldText, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldUserID, v))
}

// JokeIDEQ applies the EQ predicate on the "joke_id" field.
func JokeIDEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldJokeID, v))
}

// JokeIDNEQ applies the NEQ predicate on the "joke_id" field.
func JokeIDNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldJokeID, v))
}

// JokeIDIn applies the In predicate on the "joke_id" field.
func JokeIDIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldJokeID, vs...))
}

// JokeIDNotIn applies the NotIn predicate on the "joke_id" field.
func JokeIDNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldJokeID, vs...))
}

// JokeIDGT applies the GT predicate on the "joke_id" field.
func JokeIDGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldJokeID, v))
}

// JokeIDGTE applies the GTE predicate on the "joke_id" field.
func JokeIDGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldJokeID, v))
}

// JokeIDLT applies the LT predicate on the "joke_id" field.
func JokeIDLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldJokeID, v))
}

// JokeIDLTE applies the LTE predicate on the "joke_id" field.
func JokeIDLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldJokeID, v))
}

// JokeIDContains applies the Contains predicate on the "joke_id" field.
func JokeIDContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldJokeID, v))
}

// JokeIDHasPrefix applies the HasPrefix predicate on the "joke_id" field.
func JokeIDHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldJokeID, v))
}

// JokeIDHasSuffix applies the HasSuffix predicate on the "joke_id" field.
func JokeIDHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldJokeID, v))
}

// JokeIDEqualFold applies the EqualFold predicate on the "joke_id" field.
func JokeIDEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldJokeID, v))
}

// JokeIDContainsFold applies the ContainsFold predicate on the "joke_id" field.
func JokeIDContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldJokeID, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Comment {
	return predicate.Comment(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Comment {
	return predicate.Comment(sql.FieldContainsFold(FieldText, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Comment {
	return predicate.Comment(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCommentedJoke applies the HasEdge predicate on the "commented_joke" edge.
func HasCommentedJoke() predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CommentedJokeTable, CommentedJokeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCommentedJokeWith applies the HasEdge predicate on the "commented_joke" edge with a given conditions (other predicates).
func HasCommentedJokeWith(preds ...predicate.Joke) predicate.Comment {
	return predicate.Comment(func(s *sql.Selector) {
		step := newCommentedJokeStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Comment) predicate.Comment {
	return predicate.Comment(sql.NotPredicates(p))
}
