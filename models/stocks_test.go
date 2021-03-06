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

func testStocks(t *testing.T) {
	t.Parallel()

	query := Stocks()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testStocksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
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

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Stocks().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StockSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testStocksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := StockExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Stock exists: %s", err)
	}
	if !e {
		t.Errorf("Expected StockExists to return true, but got false.")
	}
}

func testStocksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	stockFound, err := FindStock(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if stockFound == nil {
		t.Error("want a record, got nil")
	}
}

func testStocksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Stocks().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testStocksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Stocks().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testStocksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	stockOne := &Stock{}
	stockTwo := &Stock{}
	if err = randomize.Struct(seed, stockOne, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockTwo, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testStocksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	stockOne := &Stock{}
	stockTwo := &Stock{}
	if err = randomize.Struct(seed, stockOne, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err = randomize.Struct(seed, stockTwo, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = stockOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = stockTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func stockBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func stockAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Stock) error {
	*o = Stock{}
	return nil
}

func testStocksHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Stock{}
	o := &Stock{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, stockDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Stock object: %s", err)
	}

	AddStockHook(boil.BeforeInsertHook, stockBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	stockBeforeInsertHooks = []StockHook{}

	AddStockHook(boil.AfterInsertHook, stockAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	stockAfterInsertHooks = []StockHook{}

	AddStockHook(boil.AfterSelectHook, stockAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	stockAfterSelectHooks = []StockHook{}

	AddStockHook(boil.BeforeUpdateHook, stockBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	stockBeforeUpdateHooks = []StockHook{}

	AddStockHook(boil.AfterUpdateHook, stockAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	stockAfterUpdateHooks = []StockHook{}

	AddStockHook(boil.BeforeDeleteHook, stockBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	stockBeforeDeleteHooks = []StockHook{}

	AddStockHook(boil.AfterDeleteHook, stockAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	stockAfterDeleteHooks = []StockHook{}

	AddStockHook(boil.BeforeUpsertHook, stockBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	stockBeforeUpsertHooks = []StockHook{}

	AddStockHook(boil.AfterUpsertHook, stockAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	stockAfterUpsertHooks = []StockHook{}
}

func testStocksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStocksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(stockColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testStockToOneItemUsingItem(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Stock
	var foreign Item

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, itemDBTypes, false, itemColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Item struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.ItemID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Item().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadItem(ctx, tx, false, (*[]*Stock)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Item = nil
	if err = local.L.LoadItem(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Item == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockToOneUnitUsingUnit(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Stock
	var foreign Unit

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, stockDBTypes, false, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, unitDBTypes, false, unitColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Unit struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.UnitID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Unit().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	slice := StockSlice{&local}
	if err = local.L.LoadUnit(ctx, tx, false, (*[]*Stock)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Unit == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Unit = nil
	if err = local.L.LoadUnit(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Unit == nil {
		t.Error("struct should have been eager loaded")
	}
}

func testStockToOneSetOpItemUsingItem(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Stock
	var b, c Item

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, itemDBTypes, false, strmangle.SetComplement(itemPrimaryKeyColumns, itemColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Item{&b, &c} {
		err = a.SetItem(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Item != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.ItemID != x.ID {
			t.Error("foreign key was wrong value", a.ItemID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.ItemID))
		reflect.Indirect(reflect.ValueOf(&a.ItemID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.ItemID != x.ID {
			t.Error("foreign key was wrong value", a.ItemID, x.ID)
		}
	}
}
func testStockToOneSetOpUnitUsingUnit(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Stock
	var b, c Unit

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, stockDBTypes, false, strmangle.SetComplement(stockPrimaryKeyColumns, stockColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, unitDBTypes, false, strmangle.SetComplement(unitPrimaryKeyColumns, unitColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, unitDBTypes, false, strmangle.SetComplement(unitPrimaryKeyColumns, unitColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Unit{&b, &c} {
		err = a.SetUnit(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Unit != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Stocks[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.UnitID != x.ID {
			t.Error("foreign key was wrong value", a.UnitID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.UnitID))
		reflect.Indirect(reflect.ValueOf(&a.UnitID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.UnitID != x.ID {
			t.Error("foreign key was wrong value", a.UnitID, x.ID)
		}
	}
}

func testStocksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
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

func testStocksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := StockSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testStocksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Stocks().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	stockDBTypes = map[string]string{`ID`: `text`, `ItemID`: `text`, `UnitID`: `text`, `Quantity`: `numeric`}
	_            = bytes.MinRead
)

func testStocksUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(stockAllColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stockDBTypes, true, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testStocksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(stockAllColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Stock{}
	if err = randomize.Struct(seed, o, stockDBTypes, true, stockColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, stockDBTypes, true, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(stockAllColumns, stockPrimaryKeyColumns) {
		fields = stockAllColumns
	} else {
		fields = strmangle.SetComplement(
			stockAllColumns,
			stockPrimaryKeyColumns,
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

	slice := StockSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testStocksUpsert(t *testing.T) {
	t.Parallel()

	if len(stockAllColumns) == len(stockPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Stock{}
	if err = randomize.Struct(seed, &o, stockDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Stock: %s", err)
	}

	count, err := Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, stockDBTypes, false, stockPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Stock struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Stock: %s", err)
	}

	count, err = Stocks().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
