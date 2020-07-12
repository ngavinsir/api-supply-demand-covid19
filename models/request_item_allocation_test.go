// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testRequestItemAllocations(t *testing.T) {
	t.Parallel()

	query := RequestItemAllocations()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testRequestItemAllocationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
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

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestItemAllocationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := RequestItemAllocations().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestItemAllocationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RequestItemAllocationSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testRequestItemAllocationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := RequestItemAllocationExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if RequestItemAllocation exists: %s", err)
	}
	if !e {
		t.Errorf("Expected RequestItemAllocationExists to return true, but got false.")
	}
}

func testRequestItemAllocationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	requestItemAllocationFound, err := FindRequestItemAllocation(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if requestItemAllocationFound == nil {
		t.Error("want a record, got nil")
	}
}

func testRequestItemAllocationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = RequestItemAllocations().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testRequestItemAllocationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := RequestItemAllocations().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testRequestItemAllocationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	requestItemAllocationOne := &RequestItemAllocation{}
	requestItemAllocationTwo := &RequestItemAllocation{}
	if err = randomize.Struct(seed, requestItemAllocationOne, requestItemAllocationDBTypes, false, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}
	if err = randomize.Struct(seed, requestItemAllocationTwo, requestItemAllocationDBTypes, false, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = requestItemAllocationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = requestItemAllocationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RequestItemAllocations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testRequestItemAllocationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	requestItemAllocationOne := &RequestItemAllocation{}
	requestItemAllocationTwo := &RequestItemAllocation{}
	if err = randomize.Struct(seed, requestItemAllocationOne, requestItemAllocationDBTypes, false, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}
	if err = randomize.Struct(seed, requestItemAllocationTwo, requestItemAllocationDBTypes, false, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = requestItemAllocationOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = requestItemAllocationTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func requestItemAllocationBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func requestItemAllocationAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *RequestItemAllocation) error {
	*o = RequestItemAllocation{}
	return nil
}

func testRequestItemAllocationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &RequestItemAllocation{}
	o := &RequestItemAllocation{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation object: %s", err)
	}

	AddRequestItemAllocationHook(boil.BeforeInsertHook, requestItemAllocationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationBeforeInsertHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.AfterInsertHook, requestItemAllocationAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationAfterInsertHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.AfterSelectHook, requestItemAllocationAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationAfterSelectHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.BeforeUpdateHook, requestItemAllocationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationBeforeUpdateHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.AfterUpdateHook, requestItemAllocationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationAfterUpdateHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.BeforeDeleteHook, requestItemAllocationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationBeforeDeleteHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.AfterDeleteHook, requestItemAllocationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationAfterDeleteHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.BeforeUpsertHook, requestItemAllocationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationBeforeUpsertHooks = []RequestItemAllocationHook{}

	AddRequestItemAllocationHook(boil.AfterUpsertHook, requestItemAllocationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	requestItemAllocationAfterUpsertHooks = []RequestItemAllocationHook{}
}

func testRequestItemAllocationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRequestItemAllocationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(requestItemAllocationColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testRequestItemAllocationToOneRequestItemUsingRequestItem(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local RequestItemAllocation
	var foreign RequestItem

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, requestItemAllocationDBTypes, false, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, requestItemDBTypes, false, requestItemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItem struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.RequestItemID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.RequestItem().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := RequestItemAllocationSlice{&local}
	if err = local.L.LoadRequestItem(ctx, tx, false, (*[]*RequestItemAllocation)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RequestItem == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.RequestItem = nil
	if err = local.L.LoadRequestItem(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.RequestItem == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testRequestItemAllocationToOneSetOpRequestItemUsingRequestItem(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a RequestItemAllocation
	var b, c RequestItem

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, requestItemAllocationDBTypes, false, strmangle.SetComplement(requestItemAllocationPrimaryKeyColumns, requestItemAllocationColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, requestItemDBTypes, false, strmangle.SetComplement(requestItemPrimaryKeyColumns, requestItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, requestItemDBTypes, false, strmangle.SetComplement(requestItemPrimaryKeyColumns, requestItemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*RequestItem{&b, &c} {
		err = a.SetRequestItem(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.RequestItem != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.RequestItemAllocation != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.RequestItemID != x.ID {
			t.Error("foreign key was wrong value", a.RequestItemID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.RequestItemID))
		reflect.Indirect(reflect.ValueOf(&a.RequestItemID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.RequestItemID != x.ID {
			t.Error("foreign key was wrong value", a.RequestItemID, x.ID)
		}
	}
}

func testRequestItemAllocationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
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

func testRequestItemAllocationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := RequestItemAllocationSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testRequestItemAllocationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := RequestItemAllocations().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	requestItemAllocationDBTypes = map[string]string{`ID`: `text`, `RequestItemID`: `text`, `AllocationDate`: `timestamp with time zone`, `Description`: `text`}
	_                            = bytes.MinRead
)

func testRequestItemAllocationsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(requestItemAllocationPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(requestItemAllocationAllColumns) == len(requestItemAllocationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testRequestItemAllocationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(requestItemAllocationAllColumns) == len(requestItemAllocationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &RequestItemAllocation{}
	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, requestItemAllocationDBTypes, true, requestItemAllocationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(requestItemAllocationAllColumns, requestItemAllocationPrimaryKeyColumns) {
		fields = requestItemAllocationAllColumns
	} else {
		fields = strmangle.SetComplement(
			requestItemAllocationAllColumns,
			requestItemAllocationPrimaryKeyColumns,
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

	slice := RequestItemAllocationSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testRequestItemAllocationsUpsert(t *testing.T) {
	t.Parallel()

	if len(requestItemAllocationAllColumns) == len(requestItemAllocationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := RequestItemAllocation{}
	if err = randomize.Struct(seed, &o, requestItemAllocationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RequestItemAllocation: %s", err)
	}

	count, err := RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, requestItemAllocationDBTypes, false, requestItemAllocationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize RequestItemAllocation struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert RequestItemAllocation: %s", err)
	}

	count, err = RequestItemAllocations().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
