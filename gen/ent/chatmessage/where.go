// Code generated by ent, DO NOT EDIT.

package chatmessage

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc-valera/flugo-api-golang/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldID, id))
}

// IDEqualFold applies the EqualFold predicate on the ID field.
func IDEqualFold(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEqualFold(FieldID, id))
}

// IDContainsFold applies the ContainsFold predicate on the ID field.
func IDContainsFold(id string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContainsFold(FieldID, id))
}

// ChatRoomID applies equality check predicate on the "chat_room_id" field. It's identical to ChatRoomIDEQ.
func ChatRoomID(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldChatRoomID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldUserID, v))
}

// Text applies equality check predicate on the "text" field. It's identical to TextEQ.
func Text(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldText, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// ChatRoomIDEQ applies the EQ predicate on the "chat_room_id" field.
func ChatRoomIDEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldChatRoomID, v))
}

// ChatRoomIDNEQ applies the NEQ predicate on the "chat_room_id" field.
func ChatRoomIDNEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldChatRoomID, v))
}

// ChatRoomIDIn applies the In predicate on the "chat_room_id" field.
func ChatRoomIDIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldChatRoomID, vs...))
}

// ChatRoomIDNotIn applies the NotIn predicate on the "chat_room_id" field.
func ChatRoomIDNotIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldChatRoomID, vs...))
}

// ChatRoomIDGT applies the GT predicate on the "chat_room_id" field.
func ChatRoomIDGT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldChatRoomID, v))
}

// ChatRoomIDGTE applies the GTE predicate on the "chat_room_id" field.
func ChatRoomIDGTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldChatRoomID, v))
}

// ChatRoomIDLT applies the LT predicate on the "chat_room_id" field.
func ChatRoomIDLT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldChatRoomID, v))
}

// ChatRoomIDLTE applies the LTE predicate on the "chat_room_id" field.
func ChatRoomIDLTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldChatRoomID, v))
}

// ChatRoomIDContains applies the Contains predicate on the "chat_room_id" field.
func ChatRoomIDContains(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContains(FieldChatRoomID, v))
}

// ChatRoomIDHasPrefix applies the HasPrefix predicate on the "chat_room_id" field.
func ChatRoomIDHasPrefix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasPrefix(FieldChatRoomID, v))
}

// ChatRoomIDHasSuffix applies the HasSuffix predicate on the "chat_room_id" field.
func ChatRoomIDHasSuffix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasSuffix(FieldChatRoomID, v))
}

// ChatRoomIDEqualFold applies the EqualFold predicate on the "chat_room_id" field.
func ChatRoomIDEqualFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEqualFold(FieldChatRoomID, v))
}

// ChatRoomIDContainsFold applies the ContainsFold predicate on the "chat_room_id" field.
func ChatRoomIDContainsFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContainsFold(FieldChatRoomID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContainsFold(FieldUserID, v))
}

// TextEQ applies the EQ predicate on the "text" field.
func TextEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldText, v))
}

// TextNEQ applies the NEQ predicate on the "text" field.
func TextNEQ(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldText, v))
}

// TextIn applies the In predicate on the "text" field.
func TextIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldText, vs...))
}

// TextNotIn applies the NotIn predicate on the "text" field.
func TextNotIn(vs ...string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldText, vs...))
}

// TextGT applies the GT predicate on the "text" field.
func TextGT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldText, v))
}

// TextGTE applies the GTE predicate on the "text" field.
func TextGTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldText, v))
}

// TextLT applies the LT predicate on the "text" field.
func TextLT(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldText, v))
}

// TextLTE applies the LTE predicate on the "text" field.
func TextLTE(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldText, v))
}

// TextContains applies the Contains predicate on the "text" field.
func TextContains(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContains(FieldText, v))
}

// TextHasPrefix applies the HasPrefix predicate on the "text" field.
func TextHasPrefix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasPrefix(FieldText, v))
}

// TextHasSuffix applies the HasSuffix predicate on the "text" field.
func TextHasSuffix(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldHasSuffix(FieldText, v))
}

// TextEqualFold applies the EqualFold predicate on the "text" field.
func TextEqualFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEqualFold(FieldText, v))
}

// TextContainsFold applies the ContainsFold predicate on the "text" field.
func TextContainsFold(v string) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldContainsFold(FieldText, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChatMessage {
	return predicate.ChatMessage(sql.FieldLTE(FieldCreatedAt, v))
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.ChatMessage {
	return predicate.ChatMessage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.ChatRoom) predicate.ChatMessage {
	return predicate.ChatMessage(func(s *sql.Selector) {
		step := newRoomStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasOwner applies the HasEdge predicate on the "owner" edge.
func HasOwner() predicate.ChatMessage {
	return predicate.ChatMessage(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, OwnerTable, OwnerColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasOwnerWith applies the HasEdge predicate on the "owner" edge with a given conditions (other predicates).
func HasOwnerWith(preds ...predicate.User) predicate.ChatMessage {
	return predicate.ChatMessage(func(s *sql.Selector) {
		step := newOwnerStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChatMessage) predicate.ChatMessage {
	return predicate.ChatMessage(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChatMessage) predicate.ChatMessage {
	return predicate.ChatMessage(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChatMessage) predicate.ChatMessage {
	return predicate.ChatMessage(sql.NotPredicates(p))
}