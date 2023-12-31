// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/panforyou/seller-finding/go/ent/predicate"
	"github.com/panforyou/seller-finding/go/ent/website"
)

// WebsiteUpdate is the builder for updating Website entities.
type WebsiteUpdate struct {
	config
	hooks    []Hook
	mutation *WebsiteMutation
}

// Where appends a list predicates to the WebsiteUpdate builder.
func (wu *WebsiteUpdate) Where(ps ...predicate.Website) *WebsiteUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetUpdatedAt sets the "updated_at" field.
func (wu *WebsiteUpdate) SetUpdatedAt(t time.Time) *WebsiteUpdate {
	wu.mutation.SetUpdatedAt(t)
	return wu
}

// SetKey sets the "key" field.
func (wu *WebsiteUpdate) SetKey(s string) *WebsiteUpdate {
	wu.mutation.SetKey(s)
	return wu
}

// SetURLTemplate sets the "url_template" field.
func (wu *WebsiteUpdate) SetURLTemplate(s string) *WebsiteUpdate {
	wu.mutation.SetURLTemplate(s)
	return wu
}

// SetAllowDomains sets the "allow_domains" field.
func (wu *WebsiteUpdate) SetAllowDomains(s []string) *WebsiteUpdate {
	wu.mutation.SetAllowDomains(s)
	return wu
}

// ClearAllowDomains clears the value of the "allow_domains" field.
func (wu *WebsiteUpdate) ClearAllowDomains() *WebsiteUpdate {
	wu.mutation.ClearAllowDomains()
	return wu
}

// SetLatestVisitURL sets the "latest_visit_url" field.
func (wu *WebsiteUpdate) SetLatestVisitURL(s string) *WebsiteUpdate {
	wu.mutation.SetLatestVisitURL(s)
	return wu
}

// SetNillableLatestVisitURL sets the "latest_visit_url" field if the given value is not nil.
func (wu *WebsiteUpdate) SetNillableLatestVisitURL(s *string) *WebsiteUpdate {
	if s != nil {
		wu.SetLatestVisitURL(*s)
	}
	return wu
}

// ClearLatestVisitURL clears the value of the "latest_visit_url" field.
func (wu *WebsiteUpdate) ClearLatestVisitURL() *WebsiteUpdate {
	wu.mutation.ClearLatestVisitURL()
	return wu
}

// SetMaxPage sets the "max_page" field.
func (wu *WebsiteUpdate) SetMaxPage(i int) *WebsiteUpdate {
	wu.mutation.ResetMaxPage()
	wu.mutation.SetMaxPage(i)
	return wu
}

// SetNillableMaxPage sets the "max_page" field if the given value is not nil.
func (wu *WebsiteUpdate) SetNillableMaxPage(i *int) *WebsiteUpdate {
	if i != nil {
		wu.SetMaxPage(*i)
	}
	return wu
}

// AddMaxPage adds i to the "max_page" field.
func (wu *WebsiteUpdate) AddMaxPage(i int) *WebsiteUpdate {
	wu.mutation.AddMaxPage(i)
	return wu
}

// Mutation returns the WebsiteMutation object of the builder.
func (wu *WebsiteUpdate) Mutation() *WebsiteMutation {
	return wu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WebsiteUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	wu.defaults()
	if len(wu.hooks) == 0 {
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WebsiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WebsiteUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WebsiteUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WebsiteUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wu *WebsiteUpdate) defaults() {
	if _, ok := wu.mutation.UpdatedAt(); !ok {
		v := website.UpdateDefaultUpdatedAt()
		wu.mutation.SetUpdatedAt(v)
	}
}

func (wu *WebsiteUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   website.Table,
			Columns: website.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: website.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: website.FieldUpdatedAt,
		})
	}
	if value, ok := wu.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: website.FieldKey,
		})
	}
	if value, ok := wu.mutation.URLTemplate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: website.FieldURLTemplate,
		})
	}
	if value, ok := wu.mutation.AllowDomains(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: website.FieldAllowDomains,
		})
	}
	if wu.mutation.AllowDomainsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: website.FieldAllowDomains,
		})
	}
	if value, ok := wu.mutation.LatestVisitURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: website.FieldLatestVisitURL,
		})
	}
	if wu.mutation.LatestVisitURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: website.FieldLatestVisitURL,
		})
	}
	if value, ok := wu.mutation.MaxPage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: website.FieldMaxPage,
		})
	}
	if value, ok := wu.mutation.AddedMaxPage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: website.FieldMaxPage,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{website.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// WebsiteUpdateOne is the builder for updating a single Website entity.
type WebsiteUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WebsiteMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (wuo *WebsiteUpdateOne) SetUpdatedAt(t time.Time) *WebsiteUpdateOne {
	wuo.mutation.SetUpdatedAt(t)
	return wuo
}

// SetKey sets the "key" field.
func (wuo *WebsiteUpdateOne) SetKey(s string) *WebsiteUpdateOne {
	wuo.mutation.SetKey(s)
	return wuo
}

// SetURLTemplate sets the "url_template" field.
func (wuo *WebsiteUpdateOne) SetURLTemplate(s string) *WebsiteUpdateOne {
	wuo.mutation.SetURLTemplate(s)
	return wuo
}

// SetAllowDomains sets the "allow_domains" field.
func (wuo *WebsiteUpdateOne) SetAllowDomains(s []string) *WebsiteUpdateOne {
	wuo.mutation.SetAllowDomains(s)
	return wuo
}

// ClearAllowDomains clears the value of the "allow_domains" field.
func (wuo *WebsiteUpdateOne) ClearAllowDomains() *WebsiteUpdateOne {
	wuo.mutation.ClearAllowDomains()
	return wuo
}

// SetLatestVisitURL sets the "latest_visit_url" field.
func (wuo *WebsiteUpdateOne) SetLatestVisitURL(s string) *WebsiteUpdateOne {
	wuo.mutation.SetLatestVisitURL(s)
	return wuo
}

// SetNillableLatestVisitURL sets the "latest_visit_url" field if the given value is not nil.
func (wuo *WebsiteUpdateOne) SetNillableLatestVisitURL(s *string) *WebsiteUpdateOne {
	if s != nil {
		wuo.SetLatestVisitURL(*s)
	}
	return wuo
}

// ClearLatestVisitURL clears the value of the "latest_visit_url" field.
func (wuo *WebsiteUpdateOne) ClearLatestVisitURL() *WebsiteUpdateOne {
	wuo.mutation.ClearLatestVisitURL()
	return wuo
}

// SetMaxPage sets the "max_page" field.
func (wuo *WebsiteUpdateOne) SetMaxPage(i int) *WebsiteUpdateOne {
	wuo.mutation.ResetMaxPage()
	wuo.mutation.SetMaxPage(i)
	return wuo
}

// SetNillableMaxPage sets the "max_page" field if the given value is not nil.
func (wuo *WebsiteUpdateOne) SetNillableMaxPage(i *int) *WebsiteUpdateOne {
	if i != nil {
		wuo.SetMaxPage(*i)
	}
	return wuo
}

// AddMaxPage adds i to the "max_page" field.
func (wuo *WebsiteUpdateOne) AddMaxPage(i int) *WebsiteUpdateOne {
	wuo.mutation.AddMaxPage(i)
	return wuo
}

// Mutation returns the WebsiteMutation object of the builder.
func (wuo *WebsiteUpdateOne) Mutation() *WebsiteMutation {
	return wuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WebsiteUpdateOne) Select(field string, fields ...string) *WebsiteUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Website entity.
func (wuo *WebsiteUpdateOne) Save(ctx context.Context) (*Website, error) {
	var (
		err  error
		node *Website
	)
	wuo.defaults()
	if len(wuo.hooks) == 0 {
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WebsiteMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WebsiteUpdateOne) SaveX(ctx context.Context) *Website {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WebsiteUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WebsiteUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (wuo *WebsiteUpdateOne) defaults() {
	if _, ok := wuo.mutation.UpdatedAt(); !ok {
		v := website.UpdateDefaultUpdatedAt()
		wuo.mutation.SetUpdatedAt(v)
	}
}

func (wuo *WebsiteUpdateOne) sqlSave(ctx context.Context) (_node *Website, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   website.Table,
			Columns: website.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: website.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Website.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, website.FieldID)
		for _, f := range fields {
			if !website.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != website.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.UpdatedAt(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: website.FieldUpdatedAt,
		})
	}
	if value, ok := wuo.mutation.Key(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: website.FieldKey,
		})
	}
	if value, ok := wuo.mutation.URLTemplate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: website.FieldURLTemplate,
		})
	}
	if value, ok := wuo.mutation.AllowDomains(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: website.FieldAllowDomains,
		})
	}
	if wuo.mutation.AllowDomainsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: website.FieldAllowDomains,
		})
	}
	if value, ok := wuo.mutation.LatestVisitURL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: website.FieldLatestVisitURL,
		})
	}
	if wuo.mutation.LatestVisitURLCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: website.FieldLatestVisitURL,
		})
	}
	if value, ok := wuo.mutation.MaxPage(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: website.FieldMaxPage,
		})
	}
	if value, ok := wuo.mutation.AddedMaxPage(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: website.FieldMaxPage,
		})
	}
	_node = &Website{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{website.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
