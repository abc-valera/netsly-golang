// Code generated by ent, DO NOT EDIT.

package chatmember

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/abc-valera/flugo-api-golang/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLTE(FieldID, id))
}

// ChatRoomID applies equality check predicate on the "chat_room_id" field. It's identical to ChatRoomIDEQ.
func ChatRoomID(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldChatRoomID, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldUserID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldCreatedAt, v))
}

// ChatRoomIDEQ applies the EQ predicate on the "chat_room_id" field.
func ChatRoomIDEQ(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldChatRoomID, v))
}

// ChatRoomIDNEQ applies the NEQ predicate on the "chat_room_id" field.
func ChatRoomIDNEQ(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNEQ(FieldChatRoomID, v))
}

// ChatRoomIDIn applies the In predicate on the "chat_room_id" field.
func ChatRoomIDIn(vs ...string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldIn(FieldChatRoomID, vs...))
}

// ChatRoomIDNotIn applies the NotIn predicate on the "chat_room_id" field.
func ChatRoomIDNotIn(vs ...string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNotIn(FieldChatRoomID, vs...))
}

// ChatRoomIDGT applies the GT predicate on the "chat_room_id" field.
func ChatRoomIDGT(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGT(FieldChatRoomID, v))
}

// ChatRoomIDGTE applies the GTE predicate on the "chat_room_id" field.
func ChatRoomIDGTE(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGTE(FieldChatRoomID, v))
}

// ChatRoomIDLT applies the LT predicate on the "chat_room_id" field.
func ChatRoomIDLT(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLT(FieldChatRoomID, v))
}

// ChatRoomIDLTE applies the LTE predicate on the "chat_room_id" field.
func ChatRoomIDLTE(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLTE(FieldChatRoomID, v))
}

// ChatRoomIDContains applies the Contains predicate on the "chat_room_id" field.
func ChatRoomIDContains(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldContains(FieldChatRoomID, v))
}

// ChatRoomIDHasPrefix applies the HasPrefix predicate on the "chat_room_id" field.
func ChatRoomIDHasPrefix(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldHasPrefix(FieldChatRoomID, v))
}

// ChatRoomIDHasSuffix applies the HasSuffix predicate on the "chat_room_id" field.
func ChatRoomIDHasSuffix(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldHasSuffix(FieldChatRoomID, v))
}

// ChatRoomIDEqualFold applies the EqualFold predicate on the "chat_room_id" field.
func ChatRoomIDEqualFold(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEqualFold(FieldChatRoomID, v))
}

// ChatRoomIDContainsFold applies the ContainsFold predicate on the "chat_room_id" field.
func ChatRoomIDContainsFold(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldContainsFold(FieldChatRoomID, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldContainsFold(FieldUserID, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.ChatMember {
	return predicate.ChatMember(sql.FieldLTE(FieldCreatedAt, v))
}

// HasRoom applies the HasEdge predicate on the "room" edge.
func HasRoom() predicate.ChatMember {
	return predicate.ChatMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RoomTable, RoomColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRoomWith applies the HasEdge predicate on the "room" edge with a given conditions (other predicates).
func HasRoomWith(preds ...predicate.ChatRoom) predicate.ChatMember {
	return predicate.ChatMember(func(s *sql.Selector) {
		step := newRoomStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMember applies the HasEdge predicate on the "member" edge.
func HasMember() predicate.ChatMember {
	return predicate.ChatMember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, MemberTable, MemberColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasMemberWith applies the HasEdge predicate on the "member" edge with a given conditions (other predicates).
func HasMemberWith(preds ...predicate.User) predicate.ChatMember {
	return predicate.ChatMember(func(s *sql.Selector) {
		step := newMemberStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ChatMember) predicate.ChatMember {
	return predicate.ChatMember(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ChatMember) predicate.ChatMember {
	return predicate.ChatMember(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ChatMember) predicate.ChatMember {
	return predicate.ChatMember(sql.NotPredicates(p))
}
