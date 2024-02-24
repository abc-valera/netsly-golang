// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/abc-valera/netsly-api-golang/gen/ent/room"
	"github.com/abc-valera/netsly-api-golang/gen/ent/roommember"
	"github.com/abc-valera/netsly-api-golang/gen/ent/user"
)

// RoomMember is the model entity for the RoomMember schema.
type RoomMember struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the RoomMemberQuery when eager-loading is set.
	Edges        RoomMemberEdges `json:"edges"`
	room_members *string
	user_rooms   *string
	selectValues sql.SelectValues
}

// RoomMemberEdges holds the relations/edges for other nodes in the graph.
type RoomMemberEdges struct {
	// Room holds the value of the room edge.
	Room *Room `json:"room,omitempty"`
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// RoomOrErr returns the Room value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomMemberEdges) RoomOrErr() (*Room, error) {
	if e.loadedTypes[0] {
		if e.Room == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: room.Label}
		}
		return e.Room, nil
	}
	return nil, &NotLoadedError{edge: "room"}
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e RoomMemberEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.User == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RoomMember) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case roommember.FieldID:
			values[i] = new(sql.NullInt64)
		case roommember.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case roommember.ForeignKeys[0]: // room_members
			values[i] = new(sql.NullString)
		case roommember.ForeignKeys[1]: // user_rooms
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RoomMember fields.
func (rm *RoomMember) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case roommember.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rm.ID = int(value.Int64)
		case roommember.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				rm.CreatedAt = value.Time
			}
		case roommember.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field room_members", values[i])
			} else if value.Valid {
				rm.room_members = new(string)
				*rm.room_members = value.String
			}
		case roommember.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field user_rooms", values[i])
			} else if value.Valid {
				rm.user_rooms = new(string)
				*rm.user_rooms = value.String
			}
		default:
			rm.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the RoomMember.
// This includes values selected through modifiers, order, etc.
func (rm *RoomMember) Value(name string) (ent.Value, error) {
	return rm.selectValues.Get(name)
}

// QueryRoom queries the "room" edge of the RoomMember entity.
func (rm *RoomMember) QueryRoom() *RoomQuery {
	return NewRoomMemberClient(rm.config).QueryRoom(rm)
}

// QueryUser queries the "user" edge of the RoomMember entity.
func (rm *RoomMember) QueryUser() *UserQuery {
	return NewRoomMemberClient(rm.config).QueryUser(rm)
}

// Update returns a builder for updating this RoomMember.
// Note that you need to call RoomMember.Unwrap() before calling this method if this RoomMember
// was returned from a transaction, and the transaction was committed or rolled back.
func (rm *RoomMember) Update() *RoomMemberUpdateOne {
	return NewRoomMemberClient(rm.config).UpdateOne(rm)
}

// Unwrap unwraps the RoomMember entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rm *RoomMember) Unwrap() *RoomMember {
	_tx, ok := rm.config.driver.(*txDriver)
	if !ok {
		panic("ent: RoomMember is not a transactional entity")
	}
	rm.config.driver = _tx.drv
	return rm
}

// String implements the fmt.Stringer.
func (rm *RoomMember) String() string {
	var builder strings.Builder
	builder.WriteString("RoomMember(")
	builder.WriteString(fmt.Sprintf("id=%v, ", rm.ID))
	builder.WriteString("created_at=")
	builder.WriteString(rm.CreatedAt.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// RoomMembers is a parsable slice of RoomMember.
type RoomMembers []*RoomMember
