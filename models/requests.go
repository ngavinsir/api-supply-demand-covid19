// Code generated by SQLBoiler 3.6.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// Request is an object representing the database table.
type Request struct {
	ID                  string    `boil:"id" json:"id" toml:"id" yaml:"id"`
	Date                time.Time `boil:"date" json:"date" toml:"date" yaml:"date"`
	IsFulfilled         bool      `boil:"is_fulfilled" json:"is_fulfilled" toml:"is_fulfilled" yaml:"is_fulfilled"`
	DonationApplicantID string    `boil:"donation_applicant_id" json:"donation_applicant_id" toml:"donation_applicant_id" yaml:"donation_applicant_id"`

	R *requestR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L requestL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var RequestColumns = struct {
	ID                  string
	Date                string
	IsFulfilled         string
	DonationApplicantID string
}{
	ID:                  "id",
	Date:                "date",
	IsFulfilled:         "is_fulfilled",
	DonationApplicantID: "donation_applicant_id",
}

// Generated where

var RequestWhere = struct {
	ID                  whereHelperstring
	Date                whereHelpertime_Time
	IsFulfilled         whereHelperbool
	DonationApplicantID whereHelperstring
}{
	ID:                  whereHelperstring{field: "\"requests\".\"id\""},
	Date:                whereHelpertime_Time{field: "\"requests\".\"date\""},
	IsFulfilled:         whereHelperbool{field: "\"requests\".\"is_fulfilled\""},
	DonationApplicantID: whereHelperstring{field: "\"requests\".\"donation_applicant_id\""},
}

// RequestRels is where relationship names are stored.
var RequestRels = struct {
	DonationApplicant string
	Allocations       string
	RequestItems      string
}{
	DonationApplicant: "DonationApplicant",
	Allocations:       "Allocations",
	RequestItems:      "RequestItems",
}

// requestR is where relationships are stored.
type requestR struct {
	DonationApplicant *User
	Allocations       AllocationSlice
	RequestItems      RequestItemSlice
}

// NewStruct creates a new relationship struct
func (*requestR) NewStruct() *requestR {
	return &requestR{}
}

// requestL is where Load methods for each relationship are stored.
type requestL struct{}

var (
	requestAllColumns            = []string{"id", "date", "is_fulfilled", "donation_applicant_id"}
	requestColumnsWithoutDefault = []string{"id", "date", "donation_applicant_id"}
	requestColumnsWithDefault    = []string{"is_fulfilled"}
	requestPrimaryKeyColumns     = []string{"id"}
)

type (
	// RequestSlice is an alias for a slice of pointers to Request.
	// This should generally be used opposed to []Request.
	RequestSlice []*Request
	// RequestHook is the signature for custom Request hook methods
	RequestHook func(context.Context, boil.ContextExecutor, *Request) error

	requestQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	requestType                 = reflect.TypeOf(&Request{})
	requestMapping              = queries.MakeStructMapping(requestType)
	requestPrimaryKeyMapping, _ = queries.BindMapping(requestType, requestMapping, requestPrimaryKeyColumns)
	requestInsertCacheMut       sync.RWMutex
	requestInsertCache          = make(map[string]insertCache)
	requestUpdateCacheMut       sync.RWMutex
	requestUpdateCache          = make(map[string]updateCache)
	requestUpsertCacheMut       sync.RWMutex
	requestUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var requestBeforeInsertHooks []RequestHook
var requestBeforeUpdateHooks []RequestHook
var requestBeforeDeleteHooks []RequestHook
var requestBeforeUpsertHooks []RequestHook

var requestAfterInsertHooks []RequestHook
var requestAfterSelectHooks []RequestHook
var requestAfterUpdateHooks []RequestHook
var requestAfterDeleteHooks []RequestHook
var requestAfterUpsertHooks []RequestHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Request) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Request) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Request) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Request) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Request) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Request) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Request) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Request) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Request) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range requestAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddRequestHook registers your hook function for all future operations.
func AddRequestHook(hookPoint boil.HookPoint, requestHook RequestHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		requestBeforeInsertHooks = append(requestBeforeInsertHooks, requestHook)
	case boil.BeforeUpdateHook:
		requestBeforeUpdateHooks = append(requestBeforeUpdateHooks, requestHook)
	case boil.BeforeDeleteHook:
		requestBeforeDeleteHooks = append(requestBeforeDeleteHooks, requestHook)
	case boil.BeforeUpsertHook:
		requestBeforeUpsertHooks = append(requestBeforeUpsertHooks, requestHook)
	case boil.AfterInsertHook:
		requestAfterInsertHooks = append(requestAfterInsertHooks, requestHook)
	case boil.AfterSelectHook:
		requestAfterSelectHooks = append(requestAfterSelectHooks, requestHook)
	case boil.AfterUpdateHook:
		requestAfterUpdateHooks = append(requestAfterUpdateHooks, requestHook)
	case boil.AfterDeleteHook:
		requestAfterDeleteHooks = append(requestAfterDeleteHooks, requestHook)
	case boil.AfterUpsertHook:
		requestAfterUpsertHooks = append(requestAfterUpsertHooks, requestHook)
	}
}

// One returns a single request record from the query.
func (q requestQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Request, error) {
	o := &Request{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for requests")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Request records from the query.
func (q requestQuery) All(ctx context.Context, exec boil.ContextExecutor) (RequestSlice, error) {
	var o []*Request

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Request slice")
	}

	if len(requestAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Request records in the query.
func (q requestQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count requests rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q requestQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if requests exists")
	}

	return count > 0, nil
}

// DonationApplicant pointed to by the foreign key.
func (o *Request) DonationApplicant(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.DonationApplicantID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// Allocations retrieves all the allocation's Allocations with an executor.
func (o *Request) Allocations(mods ...qm.QueryMod) allocationQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"allocations\".\"request_id\"=?", o.ID),
	)

	query := Allocations(queryMods...)
	queries.SetFrom(query.Query, "\"allocations\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"allocations\".*"})
	}

	return query
}

// RequestItems retrieves all the request_item's RequestItems with an executor.
func (o *Request) RequestItems(mods ...qm.QueryMod) requestItemQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("\"request_items\".\"request_id\"=?", o.ID),
	)

	query := RequestItems(queryMods...)
	queries.SetFrom(query.Query, "\"request_items\"")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"\"request_items\".*"})
	}

	return query
}

// LoadDonationApplicant allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (requestL) LoadDonationApplicant(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRequest interface{}, mods queries.Applicator) error {
	var slice []*Request
	var object *Request

	if singular {
		object = maybeRequest.(*Request)
	} else {
		slice = *maybeRequest.(*[]*Request)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &requestR{}
		}
		args = append(args, object.DonationApplicantID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &requestR{}
			}

			for _, a := range args {
				if a == obj.DonationApplicantID {
					continue Outer
				}
			}

			args = append(args, obj.DonationApplicantID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`users`), qm.WhereIn(`users.id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(requestAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.DonationApplicant = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.DonationApplicantRequests = append(foreign.R.DonationApplicantRequests, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.DonationApplicantID == foreign.ID {
				local.R.DonationApplicant = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.DonationApplicantRequests = append(foreign.R.DonationApplicantRequests, local)
				break
			}
		}
	}

	return nil
}

// LoadAllocations allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (requestL) LoadAllocations(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRequest interface{}, mods queries.Applicator) error {
	var slice []*Request
	var object *Request

	if singular {
		object = maybeRequest.(*Request)
	} else {
		slice = *maybeRequest.(*[]*Request)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &requestR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &requestR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`allocations`), qm.WhereIn(`allocations.request_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load allocations")
	}

	var resultSlice []*Allocation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice allocations")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on allocations")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for allocations")
	}

	if len(allocationAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.Allocations = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &allocationR{}
			}
			foreign.R.Request = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RequestID {
				local.R.Allocations = append(local.R.Allocations, foreign)
				if foreign.R == nil {
					foreign.R = &allocationR{}
				}
				foreign.R.Request = local
				break
			}
		}
	}

	return nil
}

// LoadRequestItems allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (requestL) LoadRequestItems(ctx context.Context, e boil.ContextExecutor, singular bool, maybeRequest interface{}, mods queries.Applicator) error {
	var slice []*Request
	var object *Request

	if singular {
		object = maybeRequest.(*Request)
	} else {
		slice = *maybeRequest.(*[]*Request)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &requestR{}
		}
		args = append(args, object.ID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &requestR{}
			}

			for _, a := range args {
				if a == obj.ID {
					continue Outer
				}
			}

			args = append(args, obj.ID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`request_items`), qm.WhereIn(`request_items.request_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load request_items")
	}

	var resultSlice []*RequestItem
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice request_items")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on request_items")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for request_items")
	}

	if len(requestItemAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.RequestItems = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &requestItemR{}
			}
			foreign.R.Request = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ID == foreign.RequestID {
				local.R.RequestItems = append(local.R.RequestItems, foreign)
				if foreign.R == nil {
					foreign.R = &requestItemR{}
				}
				foreign.R.Request = local
				break
			}
		}
	}

	return nil
}

// SetDonationApplicant of the request to the related item.
// Sets o.R.DonationApplicant to related.
// Adds o to related.R.DonationApplicantRequests.
func (o *Request) SetDonationApplicant(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"requests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"donation_applicant_id"}),
		strmangle.WhereClause("\"", "\"", 2, requestPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.DonationApplicantID = related.ID
	if o.R == nil {
		o.R = &requestR{
			DonationApplicant: related,
		}
	} else {
		o.R.DonationApplicant = related
	}

	if related.R == nil {
		related.R = &userR{
			DonationApplicantRequests: RequestSlice{o},
		}
	} else {
		related.R.DonationApplicantRequests = append(related.R.DonationApplicantRequests, o)
	}

	return nil
}

// AddAllocations adds the given related objects to the existing relationships
// of the request, optionally inserting them as new records.
// Appends related to o.R.Allocations.
// Sets related.R.Request appropriately.
func (o *Request) AddAllocations(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Allocation) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RequestID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"allocations\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"request_id"}),
				strmangle.WhereClause("\"", "\"", 2, allocationPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RequestID = o.ID
		}
	}

	if o.R == nil {
		o.R = &requestR{
			Allocations: related,
		}
	} else {
		o.R.Allocations = append(o.R.Allocations, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &allocationR{
				Request: o,
			}
		} else {
			rel.R.Request = o
		}
	}
	return nil
}

// AddRequestItems adds the given related objects to the existing relationships
// of the request, optionally inserting them as new records.
// Appends related to o.R.RequestItems.
// Sets related.R.Request appropriately.
func (o *Request) AddRequestItems(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RequestItem) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.RequestID = o.ID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE \"request_items\" SET %s WHERE %s",
				strmangle.SetParamNames("\"", "\"", 1, []string{"request_id"}),
				strmangle.WhereClause("\"", "\"", 2, requestItemPrimaryKeyColumns),
			)
			values := []interface{}{o.ID, rel.ID}

			if boil.IsDebug(ctx) {
				writer := boil.DebugWriterFrom(ctx)
				fmt.Fprintln(writer, updateQuery)
				fmt.Fprintln(writer, values)
			}
			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.RequestID = o.ID
		}
	}

	if o.R == nil {
		o.R = &requestR{
			RequestItems: related,
		}
	} else {
		o.R.RequestItems = append(o.R.RequestItems, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &requestItemR{
				Request: o,
			}
		} else {
			rel.R.Request = o
		}
	}
	return nil
}

// Requests retrieves all the records using an executor.
func Requests(mods ...qm.QueryMod) requestQuery {
	mods = append(mods, qm.From("\"requests\""))
	return requestQuery{NewQuery(mods...)}
}

// FindRequest retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindRequest(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*Request, error) {
	requestObj := &Request{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"requests\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, requestObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from requests")
	}

	return requestObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Request) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no requests provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(requestColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	requestInsertCacheMut.RLock()
	cache, cached := requestInsertCache[key]
	requestInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			requestAllColumns,
			requestColumnsWithDefault,
			requestColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(requestType, requestMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(requestType, requestMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"requests\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"requests\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into requests")
	}

	if !cached {
		requestInsertCacheMut.Lock()
		requestInsertCache[key] = cache
		requestInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Request.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Request) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	requestUpdateCacheMut.RLock()
	cache, cached := requestUpdateCache[key]
	requestUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			requestAllColumns,
			requestPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update requests, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"requests\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, requestPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(requestType, requestMapping, append(wl, requestPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update requests row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for requests")
	}

	if !cached {
		requestUpdateCacheMut.Lock()
		requestUpdateCache[key] = cache
		requestUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q requestQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for requests")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o RequestSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"requests\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, requestPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in request slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all request")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Request) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no requests provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(requestColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	requestUpsertCacheMut.RLock()
	cache, cached := requestUpsertCache[key]
	requestUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			requestAllColumns,
			requestColumnsWithDefault,
			requestColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			requestAllColumns,
			requestPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert requests, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(requestPrimaryKeyColumns))
			copy(conflict, requestPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"requests\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(requestType, requestMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(requestType, requestMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert requests")
	}

	if !cached {
		requestUpsertCacheMut.Lock()
		requestUpsertCache[key] = cache
		requestUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Request record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Request) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Request provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), requestPrimaryKeyMapping)
	sql := "DELETE FROM \"requests\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for requests")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q requestQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no requestQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from requests")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for requests")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o RequestSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	if len(requestBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"requests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, requestPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from request slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for requests")
	}

	if len(requestAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Request) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindRequest(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *RequestSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := RequestSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), requestPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"requests\".* FROM \"requests\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, requestPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in RequestSlice")
	}

	*o = slice

	return nil
}

// RequestExists checks if the Request row exists.
func RequestExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"requests\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if requests exists")
	}

	return exists, nil
}
