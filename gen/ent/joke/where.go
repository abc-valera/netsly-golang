// Code generated by ent, DO NOT EDIT.

package joke

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc-valera/flugo-api-golang/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.Joke {
	return predicate.Joke(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.Joke {
	return predicate.Joke(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.Joke {
	return predicate.Joke(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.Joke {
	return predicate.Joke(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.Joke {
	return predicate.Joke(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.Joke {
	return predicate.Joke(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.Joke {
	return predicate.Joke(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.Joke {
	return predicate.Joke(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.Joke {
	return predicate.Joke(sql.FieldContainsFold(FieldID, id))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldUserID, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldTitle, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldText, v))
}

// Explanation applies equality check predicate on the "explanation" field. It's identical to ExplanationEQ.
func Explanation(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldExplanation, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldCreatedAt, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContainsFold(FieldUserID, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContainsFold(FieldTitle, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContainsFold(FieldText, v))
}

// ExplanationEQ applies the EQ predicate on the "explanation" field.
func ExplanationEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldExplanation, v))
}

// ExplanationNEQ applies the NEQ predicate on the "explanation" field.
func ExplanationNEQ(v string) predicate.Joke {
	return predicate.Joke(sql.FieldNEQ(FieldExplanation, v))
}

// ExplanationIn applies the In predicate on the "explanation" field.
func ExplanationIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldIn(FieldExplanation, vs...))
}

// ExplanationNotIn applies the NotIn predicate on the "explanation" field.
func ExplanationNotIn(vs ...string) predicate.Joke {
	return predicate.Joke(sql.FieldNotIn(FieldExplanation, vs...))
}

// ExplanationGT applies the GT predicate on the "explanation" field.
func ExplanationGT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGT(FieldExplanation, v))
}

// ExplanationGTE applies the GTE predicate on the "explanation" field.
func ExplanationGTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldGTE(FieldExplanation, v))
}

// ExplanationLT applies the LT predicate on the "explanation" field.
func ExplanationLT(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLT(FieldExplanation, v))
}

// ExplanationLTE applies the LTE predicate on the "explanation" field.
func ExplanationLTE(v string) predicate.Joke {
	return predicate.Joke(sql.FieldLTE(FieldExplanation, v))
}

// ExplanationContains applies the Contains predicate on the "explanation" field.
func ExplanationContains(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContains(FieldExplanation, v))
}

// ExplanationHasPrefix applies the HasPrefix predicate on the "explanation" field.
func ExplanationHasPrefix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasPrefix(FieldExplanation, v))
}

// ExplanationHasSuffix applies the HasSuffix predicate on the "explanation" field.
func ExplanationHasSuffix(v string) predicate.Joke {
	return predicate.Joke(sql.FieldHasSuffix(FieldExplanation, v))
}

// ExplanationEqualFold applies the EqualFold predicate on the "explanation" field.
func ExplanationEqualFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldEqualFold(FieldExplanation, v))
}

// ExplanationContainsFold applies the ContainsFold predicate on the "explanation" field.
func ExplanationContainsFold(v string) predicate.Joke {
	return predicate.Joke(sql.FieldContainsFold(FieldExplanation, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Joke {
	return predicate.Joke(sql.FieldLTE(FieldCreatedAt, v))
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.Joke {
	return predicate.Joke(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.Joke {
	return predicate.Joke(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Joke) predicate.Joke {
	return predicate.Joke(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Joke) predicate.Joke {
	return predicate.Joke(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Joke) predicate.Joke {
	return predicate.Joke(sql.NotPredicates(p))
}
