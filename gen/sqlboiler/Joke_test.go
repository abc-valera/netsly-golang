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

func testJokes(t *testing.T) {
	t.Parallel()

	query := Jokes()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testJokesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
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

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJokesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Jokes().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJokesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JokeSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJokesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := JokeExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Joke exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JokeExists to return true, but got false.")
	}
}

func testJokesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	jokeFound, err := FindJoke(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if jokeFound == nil {
		t.Error("want a record, got nil")
	}
}

func testJokesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Jokes().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testJokesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Jokes().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJokesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jokeOne := &Joke{}
	jokeTwo := &Joke{}
	if err = randomize.Struct(seed, jokeOne, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}
	if err = randomize.Struct(seed, jokeTwo, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jokeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jokeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jokes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJokesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jokeOne := &Joke{}
	jokeTwo := &Joke{}
	if err = randomize.Struct(seed, jokeOne, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}
	if err = randomize.Struct(seed, jokeTwo, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = jokeOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = jokeTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func jokeBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func jokeAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Joke) error {
	*o = Joke{}
	return nil
}

func testJokesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Joke{}
	o := &Joke{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jokeDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Joke object: %s", err)
	}

	AddJokeHook(boil.BeforeInsertHook, jokeBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jokeBeforeInsertHooks = []JokeHook{}

	AddJokeHook(boil.AfterInsertHook, jokeAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jokeAfterInsertHooks = []JokeHook{}

	AddJokeHook(boil.AfterSelectHook, jokeAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jokeAfterSelectHooks = []JokeHook{}

	AddJokeHook(boil.BeforeUpdateHook, jokeBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jokeBeforeUpdateHooks = []JokeHook{}

	AddJokeHook(boil.AfterUpdateHook, jokeAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jokeAfterUpdateHooks = []JokeHook{}

	AddJokeHook(boil.BeforeDeleteHook, jokeBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jokeBeforeDeleteHooks = []JokeHook{}

	AddJokeHook(boil.AfterDeleteHook, jokeAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jokeAfterDeleteHooks = []JokeHook{}

	AddJokeHook(boil.BeforeUpsertHook, jokeBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jokeBeforeUpsertHooks = []JokeHook{}

	AddJokeHook(boil.AfterUpsertHook, jokeAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jokeAfterUpsertHooks = []JokeHook{}
}

func testJokesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJokesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(jokeColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJokeToManyJokeComments(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Joke
	var b, c Comment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, commentDBTypes, false, commentColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.JokeID = a.ID
	c.JokeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.JokeComments().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.JokeID == b.JokeID {
			bFound = true
		}
		if v.JokeID == c.JokeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := JokeSlice{&a}
	if err = a.L.LoadJokeComments(ctx, tx, false, (*[]*Joke)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JokeComments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.JokeComments = nil
	if err = a.L.LoadJokeComments(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JokeComments); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testJokeToManyJokeLikes(t *testing.T) {
	var err error
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Joke
	var b, c Like

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = randomize.Struct(seed, &b, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, likeDBTypes, false, likeColumnsWithDefault...); err != nil {
		t.Fatal(err)
	}

	b.JokeID = a.ID
	c.JokeID = a.ID

	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := a.JokeLikes().All(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	bFound, cFound := false, false
	for _, v := range check {
		if v.JokeID == b.JokeID {
			bFound = true
		}
		if v.JokeID == c.JokeID {
			cFound = true
		}
	}

	if !bFound {
		t.Error("expected to find b")
	}
	if !cFound {
		t.Error("expected to find c")
	}

	slice := JokeSlice{&a}
	if err = a.L.LoadJokeLikes(ctx, tx, false, (*[]*Joke)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JokeLikes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	a.R.JokeLikes = nil
	if err = a.L.LoadJokeLikes(ctx, tx, true, &a, nil); err != nil {
		t.Fatal(err)
	}
	if got := len(a.R.JokeLikes); got != 2 {
		t.Error("number of eager loaded records wrong, got:", got)
	}

	if t.Failed() {
		t.Logf("%#v", check)
	}
}

func testJokeToManyAddOpJokeComments(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Joke
	var b, c, d, e Comment

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jokeDBTypes, false, strmangle.SetComplement(jokePrimaryKeyColumns, jokeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Comment{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, commentDBTypes, false, strmangle.SetComplement(commentPrimaryKeyColumns, commentColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Comment{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddJokeComments(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.JokeID {
			t.Error("foreign key was wrong value", a.ID, first.JokeID)
		}
		if a.ID != second.JokeID {
			t.Error("foreign key was wrong value", a.ID, second.JokeID)
		}

		if first.R.Joke != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Joke != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.JokeComments[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.JokeComments[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.JokeComments().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testJokeToManyAddOpJokeLikes(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Joke
	var b, c, d, e Like

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jokeDBTypes, false, strmangle.SetComplement(jokePrimaryKeyColumns, jokeColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	foreigners := []*Like{&b, &c, &d, &e}
	for _, x := range foreigners {
		if err = randomize.Struct(seed, x, likeDBTypes, false, strmangle.SetComplement(likePrimaryKeyColumns, likeColumnsWithoutDefault)...); err != nil {
			t.Fatal(err)
		}
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = c.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	foreignersSplitByInsertion := [][]*Like{
		{&b, &c},
		{&d, &e},
	}

	for i, x := range foreignersSplitByInsertion {
		err = a.AddJokeLikes(ctx, tx, i != 0, x...)
		if err != nil {
			t.Fatal(err)
		}

		first := x[0]
		second := x[1]

		if a.ID != first.JokeID {
			t.Error("foreign key was wrong value", a.ID, first.JokeID)
		}
		if a.ID != second.JokeID {
			t.Error("foreign key was wrong value", a.ID, second.JokeID)
		}

		if first.R.Joke != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}
		if second.R.Joke != &a {
			t.Error("relationship was not added properly to the foreign slice")
		}

		if a.R.JokeLikes[i*2] != first {
			t.Error("relationship struct slice not set to correct value")
		}
		if a.R.JokeLikes[i*2+1] != second {
			t.Error("relationship struct slice not set to correct value")
		}

		count, err := a.JokeLikes().Count(ctx, tx)
		if err != nil {
			t.Fatal(err)
		}
		if want := int64((i + 1) * 2); count != want {
			t.Error("want", want, "got", count)
		}
	}
}
func testJokeToOneUserUsingUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Joke
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, jokeDBTypes, false, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
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

	slice := JokeSlice{&local}
	if err = local.L.LoadUser(ctx, tx, false, (*[]*Joke)(&slice), nil); err != nil {
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

func testJokeToOneSetOpUserUsingUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Joke
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, jokeDBTypes, false, strmangle.SetComplement(jokePrimaryKeyColumns, jokeColumnsWithoutDefault)...); err != nil {
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

		if x.R.UserJokes[0] != &a {
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

func testJokesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
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

func testJokesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := JokeSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testJokesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Jokes().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jokeDBTypes = map[string]string{`ID`: `uuid`, `Title`: `character varying`, `Text`: `character varying`, `Explanation`: `character varying`, `CreatedAt`: `timestamp without time zone`, `UserID`: `uuid`}
	_           = bytes.MinRead
)

func testJokesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(jokeAllColumns) == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testJokesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jokeAllColumns) == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Joke{}
	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokeColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, jokeDBTypes, true, jokePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jokeAllColumns, jokePrimaryKeyColumns) {
		fields = jokeAllColumns
	} else {
		fields = strmangle.SetComplement(
			jokeAllColumns,
			jokePrimaryKeyColumns,
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

	slice := JokeSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testJokesUpsert(t *testing.T) {
	t.Parallel()

	if len(jokeAllColumns) == len(jokePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Joke{}
	if err = randomize.Struct(seed, &o, jokeDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Joke: %s", err)
	}

	count, err := Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, jokeDBTypes, false, jokePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Joke struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Joke: %s", err)
	}

	count, err = Jokes().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
