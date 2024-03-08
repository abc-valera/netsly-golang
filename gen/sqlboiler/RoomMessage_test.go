// Code generated by SQLBoiler 4.16.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package sqlboiler

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRoomMessages(t *testing.T) {
	t.Parallel()

	query := RoomMessages()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRoomMessagesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoomMessagesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RoomMessages().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoomMessagesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoomMessageSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRoomMessagesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RoomMessageExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RoomMessage exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RoomMessageExists to return true, but got false.")
	}
}

func testRoomMessagesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	roomMessageFound, err := FindRoomMessage(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if roomMessageFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRoomMessagesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RoomMessages().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRoomMessagesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RoomMessages().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRoomMessagesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	roomMessageOne := &RoomMessage{}
	roomMessageTwo := &RoomMessage{}
	if err = randomize.Struct(seed, roomMessageOne, roomMessageDBTypes, false, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, roomMessageTwo, roomMessageDBTypes, false, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roomMessageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roomMessageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoomMessages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRoomMessagesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	roomMessageOne := &RoomMessage{}
	roomMessageTwo := &RoomMessage{}
	if err = randomize.Struct(seed, roomMessageOne, roomMessageDBTypes, false, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}
	if err = randomize.Struct(seed, roomMessageTwo, roomMessageDBTypes, false, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = roomMessageOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = roomMessageTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func roomMessageBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func roomMessageAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RoomMessage) error {
	*o = RoomMessage{}
	return nil
}

func testRoomMessagesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RoomMessage{}
	o := &RoomMessage{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, roomMessageDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RoomMessage object: %s", err)
	}

	AddRoomMessageHook(boil.BeforeInsertHook, roomMessageBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	roomMessageBeforeInsertHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.AfterInsertHook, roomMessageAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	roomMessageAfterInsertHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.AfterSelectHook, roomMessageAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	roomMessageAfterSelectHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.BeforeUpdateHook, roomMessageBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	roomMessageBeforeUpdateHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.AfterUpdateHook, roomMessageAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	roomMessageAfterUpdateHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.BeforeDeleteHook, roomMessageBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	roomMessageBeforeDeleteHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.AfterDeleteHook, roomMessageAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	roomMessageAfterDeleteHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.BeforeUpsertHook, roomMessageBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	roomMessageBeforeUpsertHooks = []RoomMessageHook{}

	AddRoomMessageHook(boil.AfterUpsertHook, roomMessageAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	roomMessageAfterUpsertHooks = []RoomMessageHook{}
}

func testRoomMessagesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoomMessagesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(roomMessageColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRoomMessageToOneRoomUsingRoom(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoomMessage
	var foreign Room

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roomMessageDBTypes, false, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, roomDBTypes, false, roomColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Room struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RoomID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Room().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddRoomHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Room) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := RoomMessageSlice{&local}
	if err = local.L.LoadRoom(ctx, tx, false, (*[]*RoomMessage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Room == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Room = nil
	if err = local.L.LoadRoom(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Room == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testRoomMessageToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RoomMessage
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, roomMessageDBTypes, false, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UserID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.User().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := RoomMessageSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*RoomMessage)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.User = nil
	if err = local.L.LoadUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.User == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testRoomMessageToOneSetOpRoomUsingRoom(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoomMessage
	var b, c Room

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roomMessageDBTypes, false, strmangle.SetComplement(roomMessagePrimaryKeyColumns, roomMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, roomDBTypes, false, strmangle.SetComplement(roomPrimaryKeyColumns, roomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, roomDBTypes, false, strmangle.SetComplement(roomPrimaryKeyColumns, roomColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Room{&b, &c} {
		err = a.SetRoom(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Room != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RoomRoomMessages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RoomID != x.ID {
			t.Error("foreign key was wrong value", a.RoomID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RoomID))
		reflect.Indirect(reflect.ValueOf(&a.RoomID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RoomID != x.ID {
			t.Error("foreign key was wrong value", a.RoomID, x.ID)
		}
	}
}
func testRoomMessageToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RoomMessage
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, roomMessageDBTypes, false, strmangle.SetComplement(roomMessagePrimaryKeyColumns, roomMessageColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.User != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.UserRoomMessages[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UserID))
		reflect.Indirect(reflect.ValueOf(&a.UserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UserID != x.ID {
			t.Error("foreign key was wrong value", a.UserID, x.ID)
		}
	}
}

func testRoomMessagesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoomMessagesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RoomMessageSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRoomMessagesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RoomMessages().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	roomMessageDBTypes = map[string]string{`ID`: `uuid`, `Text`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UserID`: `uuid`, `RoomID`: `uuid`}
	_                  = bytes.MinRead
)

func testRoomMessagesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(roomMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(roomMessageAllColumns) == len(roomMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRoomMessagesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(roomMessageAllColumns) == len(roomMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RoomMessage{}
	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessageColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, roomMessageDBTypes, true, roomMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(roomMessageAllColumns, roomMessagePrimaryKeyColumns) {
		fields = roomMessageAllColumns
	} else {
		fields = strmangle.SetComplement(
			roomMessageAllColumns,
			roomMessagePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := RoomMessageSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRoomMessagesUpsert(t *testing.T) {
	t.Parallel()

	if len(roomMessageAllColumns) == len(roomMessagePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RoomMessage{}
	if err = randomize.Struct(seed, &o, roomMessageDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoomMessage: %s", err)
	}

	count, err := RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, roomMessageDBTypes, false, roomMessagePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RoomMessage struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RoomMessage: %s", err)
	}

	count, err = RoomMessages().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}