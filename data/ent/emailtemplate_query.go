// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/emailtemplate"
	"github.com/ngdangkietswe/swe-notification-service/data/ent/predicate"
)

// EmailTemplateQuery is the builder for querying EmailTemplate entities.
type EmailTemplateQuery struct {
	config
	ctx        *QueryContext
	order      []emailtemplate.OrderOption
	inters     []Interceptor
	predicates []predicate.EmailTemplate
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the EmailTemplateQuery builder.
func (etq *EmailTemplateQuery) Where(ps ...predicate.EmailTemplate) *EmailTemplateQuery {
	etq.predicates = append(etq.predicates, ps...)
	return etq
}

// Limit the number of records to be returned by this query.
func (etq *EmailTemplateQuery) Limit(limit int) *EmailTemplateQuery {
	etq.ctx.Limit = &limit
	return etq
}

// Offset to start from.
func (etq *EmailTemplateQuery) Offset(offset int) *EmailTemplateQuery {
	etq.ctx.Offset = &offset
	return etq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (etq *EmailTemplateQuery) Unique(unique bool) *EmailTemplateQuery {
	etq.ctx.Unique = &unique
	return etq
}

// Order specifies how the records should be ordered.
func (etq *EmailTemplateQuery) Order(o ...emailtemplate.OrderOption) *EmailTemplateQuery {
	etq.order = append(etq.order, o...)
	return etq
}

// First returns the first EmailTemplate entity from the query.
// Returns a *NotFoundError when no EmailTemplate was found.
func (etq *EmailTemplateQuery) First(ctx context.Context) (*EmailTemplate, error) {
	nodes, err := etq.Limit(1).All(setContextOp(ctx, etq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{emailtemplate.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (etq *EmailTemplateQuery) FirstX(ctx context.Context) *EmailTemplate {
	node, err := etq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first EmailTemplate ID from the query.
// Returns a *NotFoundError when no EmailTemplate ID was found.
func (etq *EmailTemplateQuery) FirstID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = etq.Limit(1).IDs(setContextOp(ctx, etq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{emailtemplate.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (etq *EmailTemplateQuery) FirstIDX(ctx context.Context) uuid.UUID {
	id, err := etq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single EmailTemplate entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one EmailTemplate entity is found.
// Returns a *NotFoundError when no EmailTemplate entities are found.
func (etq *EmailTemplateQuery) Only(ctx context.Context) (*EmailTemplate, error) {
	nodes, err := etq.Limit(2).All(setContextOp(ctx, etq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{emailtemplate.Label}
	default:
		return nil, &NotSingularError{emailtemplate.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (etq *EmailTemplateQuery) OnlyX(ctx context.Context) *EmailTemplate {
	node, err := etq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only EmailTemplate ID in the query.
// Returns a *NotSingularError when more than one EmailTemplate ID is found.
// Returns a *NotFoundError when no entities are found.
func (etq *EmailTemplateQuery) OnlyID(ctx context.Context) (id uuid.UUID, err error) {
	var ids []uuid.UUID
	if ids, err = etq.Limit(2).IDs(setContextOp(ctx, etq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{emailtemplate.Label}
	default:
		err = &NotSingularError{emailtemplate.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (etq *EmailTemplateQuery) OnlyIDX(ctx context.Context) uuid.UUID {
	id, err := etq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of EmailTemplates.
func (etq *EmailTemplateQuery) All(ctx context.Context) ([]*EmailTemplate, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryAll)
	if err := etq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*EmailTemplate, *EmailTemplateQuery]()
	return withInterceptors[[]*EmailTemplate](ctx, etq, qr, etq.inters)
}

// AllX is like All, but panics if an error occurs.
func (etq *EmailTemplateQuery) AllX(ctx context.Context) []*EmailTemplate {
	nodes, err := etq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of EmailTemplate IDs.
func (etq *EmailTemplateQuery) IDs(ctx context.Context) (ids []uuid.UUID, err error) {
	if etq.ctx.Unique == nil && etq.path != nil {
		etq.Unique(true)
	}
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryIDs)
	if err = etq.Select(emailtemplate.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (etq *EmailTemplateQuery) IDsX(ctx context.Context) []uuid.UUID {
	ids, err := etq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (etq *EmailTemplateQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryCount)
	if err := etq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, etq, querierCount[*EmailTemplateQuery](), etq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (etq *EmailTemplateQuery) CountX(ctx context.Context) int {
	count, err := etq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (etq *EmailTemplateQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, etq.ctx, ent.OpQueryExist)
	switch _, err := etq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (etq *EmailTemplateQuery) ExistX(ctx context.Context) bool {
	exist, err := etq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the EmailTemplateQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (etq *EmailTemplateQuery) Clone() *EmailTemplateQuery {
	if etq == nil {
		return nil
	}
	return &EmailTemplateQuery{
		config:     etq.config,
		ctx:        etq.ctx.Clone(),
		order:      append([]emailtemplate.OrderOption{}, etq.order...),
		inters:     append([]Interceptor{}, etq.inters...),
		predicates: append([]predicate.EmailTemplate{}, etq.predicates...),
		// clone intermediate query.
		sql:  etq.sql.Clone(),
		path: etq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		TemplateKey string `json:"template_key,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.EmailTemplate.Query().
//		GroupBy(emailtemplate.FieldTemplateKey).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (etq *EmailTemplateQuery) GroupBy(field string, fields ...string) *EmailTemplateGroupBy {
	etq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &EmailTemplateGroupBy{build: etq}
	grbuild.flds = &etq.ctx.Fields
	grbuild.label = emailtemplate.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		TemplateKey string `json:"template_key,omitempty"`
//	}
//
//	client.EmailTemplate.Query().
//		Select(emailtemplate.FieldTemplateKey).
//		Scan(ctx, &v)
func (etq *EmailTemplateQuery) Select(fields ...string) *EmailTemplateSelect {
	etq.ctx.Fields = append(etq.ctx.Fields, fields...)
	sbuild := &EmailTemplateSelect{EmailTemplateQuery: etq}
	sbuild.label = emailtemplate.Label
	sbuild.flds, sbuild.scan = &etq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a EmailTemplateSelect configured with the given aggregations.
func (etq *EmailTemplateQuery) Aggregate(fns ...AggregateFunc) *EmailTemplateSelect {
	return etq.Select().Aggregate(fns...)
}

func (etq *EmailTemplateQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range etq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, etq); err != nil {
				return err
			}
		}
	}
	for _, f := range etq.ctx.Fields {
		if !emailtemplate.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if etq.path != nil {
		prev, err := etq.path(ctx)
		if err != nil {
			return err
		}
		etq.sql = prev
	}
	return nil
}

func (etq *EmailTemplateQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*EmailTemplate, error) {
	var (
		nodes = []*EmailTemplate{}
		_spec = etq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*EmailTemplate).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &EmailTemplate{config: etq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, etq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (etq *EmailTemplateQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := etq.querySpec()
	_spec.Node.Columns = etq.ctx.Fields
	if len(etq.ctx.Fields) > 0 {
		_spec.Unique = etq.ctx.Unique != nil && *etq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, etq.driver, _spec)
}

func (etq *EmailTemplateQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(emailtemplate.Table, emailtemplate.Columns, sqlgraph.NewFieldSpec(emailtemplate.FieldID, field.TypeUUID))
	_spec.From = etq.sql
	if unique := etq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if etq.path != nil {
		_spec.Unique = true
	}
	if fields := etq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, emailtemplate.FieldID)
		for i := range fields {
			if fields[i] != emailtemplate.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := etq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := etq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := etq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := etq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (etq *EmailTemplateQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(etq.driver.Dialect())
	t1 := builder.Table(emailtemplate.Table)
	columns := etq.ctx.Fields
	if len(columns) == 0 {
		columns = emailtemplate.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if etq.sql != nil {
		selector = etq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if etq.ctx.Unique != nil && *etq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range etq.predicates {
		p(selector)
	}
	for _, p := range etq.order {
		p(selector)
	}
	if offset := etq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := etq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// EmailTemplateGroupBy is the group-by builder for EmailTemplate entities.
type EmailTemplateGroupBy struct {
	selector
	build *EmailTemplateQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (etgb *EmailTemplateGroupBy) Aggregate(fns ...AggregateFunc) *EmailTemplateGroupBy {
	etgb.fns = append(etgb.fns, fns...)
	return etgb
}

// Scan applies the selector query and scans the result into the given value.
func (etgb *EmailTemplateGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, etgb.build.ctx, ent.OpQueryGroupBy)
	if err := etgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailTemplateQuery, *EmailTemplateGroupBy](ctx, etgb.build, etgb, etgb.build.inters, v)
}

func (etgb *EmailTemplateGroupBy) sqlScan(ctx context.Context, root *EmailTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(etgb.fns))
	for _, fn := range etgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*etgb.flds)+len(etgb.fns))
		for _, f := range *etgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*etgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := etgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// EmailTemplateSelect is the builder for selecting fields of EmailTemplate entities.
type EmailTemplateSelect struct {
	*EmailTemplateQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ets *EmailTemplateSelect) Aggregate(fns ...AggregateFunc) *EmailTemplateSelect {
	ets.fns = append(ets.fns, fns...)
	return ets
}

// Scan applies the selector query and scans the result into the given value.
func (ets *EmailTemplateSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ets.ctx, ent.OpQuerySelect)
	if err := ets.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*EmailTemplateQuery, *EmailTemplateSelect](ctx, ets.EmailTemplateQuery, ets, ets.inters, v)
}

func (ets *EmailTemplateSelect) sqlScan(ctx context.Context, root *EmailTemplateQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ets.fns))
	for _, fn := range ets.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ets.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ets.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
