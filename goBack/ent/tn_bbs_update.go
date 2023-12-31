// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"goBack/ent/predicate"
	"goBack/ent/tn_bbs"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TNBBSUpdate is the builder for updating TN_BBS entities.
type TNBBSUpdate struct {
	config
	hooks    []Hook
	mutation *TNBBSMutation
}

// Where appends a list predicates to the TNBBSUpdate builder.
func (tu *TNBBSUpdate) Where(ps ...predicate.TN_BBS) *TNBBSUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUserID sets the "user_id" field.
func (tu *TNBBSUpdate) SetUserID(s string) *TNBBSUpdate {
	tu.mutation.SetUserID(s)
	return tu
}

// SetUserName sets the "user_name" field.
func (tu *TNBBSUpdate) SetUserName(s string) *TNBBSUpdate {
	tu.mutation.SetUserName(s)
	return tu
}

// SetTitle sets the "title" field.
func (tu *TNBBSUpdate) SetTitle(s string) *TNBBSUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetContent sets the "content" field.
func (tu *TNBBSUpdate) SetContent(s string) *TNBBSUpdate {
	tu.mutation.SetContent(s)
	return tu
}

// SetDelectYn sets the "delect_yn" field.
func (tu *TNBBSUpdate) SetDelectYn(s string) *TNBBSUpdate {
	tu.mutation.SetDelectYn(s)
	return tu
}

// SetRegDate sets the "reg_date" field.
func (tu *TNBBSUpdate) SetRegDate(s string) *TNBBSUpdate {
	tu.mutation.SetRegDate(s)
	return tu
}

// SetUdtDate sets the "udt_date" field.
func (tu *TNBBSUpdate) SetUdtDate(s string) *TNBBSUpdate {
	tu.mutation.SetUdtDate(s)
	return tu
}

// Mutation returns the TNBBSMutation object of the builder.
func (tu *TNBBSUpdate) Mutation() *TNBBSMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TNBBSUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TNBBSUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TNBBSUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TNBBSUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tu *TNBBSUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(tn_bbs.Table, tn_bbs.Columns, sqlgraph.NewFieldSpec(tn_bbs.FieldID, field.TypeInt))
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UserID(); ok {
		_spec.SetField(tn_bbs.FieldUserID, field.TypeString, value)
	}
	if value, ok := tu.mutation.UserName(); ok {
		_spec.SetField(tn_bbs.FieldUserName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.SetField(tn_bbs.FieldTitle, field.TypeString, value)
	}
	if value, ok := tu.mutation.Content(); ok {
		_spec.SetField(tn_bbs.FieldContent, field.TypeString, value)
	}
	if value, ok := tu.mutation.DelectYn(); ok {
		_spec.SetField(tn_bbs.FieldDelectYn, field.TypeString, value)
	}
	if value, ok := tu.mutation.RegDate(); ok {
		_spec.SetField(tn_bbs.FieldRegDate, field.TypeString, value)
	}
	if value, ok := tu.mutation.UdtDate(); ok {
		_spec.SetField(tn_bbs.FieldUdtDate, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tn_bbs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TNBBSUpdateOne is the builder for updating a single TN_BBS entity.
type TNBBSUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TNBBSMutation
}

// SetUserID sets the "user_id" field.
func (tuo *TNBBSUpdateOne) SetUserID(s string) *TNBBSUpdateOne {
	tuo.mutation.SetUserID(s)
	return tuo
}

// SetUserName sets the "user_name" field.
func (tuo *TNBBSUpdateOne) SetUserName(s string) *TNBBSUpdateOne {
	tuo.mutation.SetUserName(s)
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TNBBSUpdateOne) SetTitle(s string) *TNBBSUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetContent sets the "content" field.
func (tuo *TNBBSUpdateOne) SetContent(s string) *TNBBSUpdateOne {
	tuo.mutation.SetContent(s)
	return tuo
}

// SetDelectYn sets the "delect_yn" field.
func (tuo *TNBBSUpdateOne) SetDelectYn(s string) *TNBBSUpdateOne {
	tuo.mutation.SetDelectYn(s)
	return tuo
}

// SetRegDate sets the "reg_date" field.
func (tuo *TNBBSUpdateOne) SetRegDate(s string) *TNBBSUpdateOne {
	tuo.mutation.SetRegDate(s)
	return tuo
}

// SetUdtDate sets the "udt_date" field.
func (tuo *TNBBSUpdateOne) SetUdtDate(s string) *TNBBSUpdateOne {
	tuo.mutation.SetUdtDate(s)
	return tuo
}

// Mutation returns the TNBBSMutation object of the builder.
func (tuo *TNBBSUpdateOne) Mutation() *TNBBSMutation {
	return tuo.mutation
}

// Where appends a list predicates to the TNBBSUpdate builder.
func (tuo *TNBBSUpdateOne) Where(ps ...predicate.TN_BBS) *TNBBSUpdateOne {
	tuo.mutation.Where(ps...)
	return tuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TNBBSUpdateOne) Select(field string, fields ...string) *TNBBSUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated TN_BBS entity.
func (tuo *TNBBSUpdateOne) Save(ctx context.Context) (*TN_BBS, error) {
	return withHooks(ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TNBBSUpdateOne) SaveX(ctx context.Context) *TN_BBS {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TNBBSUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TNBBSUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuo *TNBBSUpdateOne) sqlSave(ctx context.Context) (_node *TN_BBS, err error) {
	_spec := sqlgraph.NewUpdateSpec(tn_bbs.Table, tn_bbs.Columns, sqlgraph.NewFieldSpec(tn_bbs.FieldID, field.TypeInt))
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TN_BBS.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tn_bbs.FieldID)
		for _, f := range fields {
			if !tn_bbs.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tn_bbs.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UserID(); ok {
		_spec.SetField(tn_bbs.FieldUserID, field.TypeString, value)
	}
	if value, ok := tuo.mutation.UserName(); ok {
		_spec.SetField(tn_bbs.FieldUserName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.SetField(tn_bbs.FieldTitle, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Content(); ok {
		_spec.SetField(tn_bbs.FieldContent, field.TypeString, value)
	}
	if value, ok := tuo.mutation.DelectYn(); ok {
		_spec.SetField(tn_bbs.FieldDelectYn, field.TypeString, value)
	}
	if value, ok := tuo.mutation.RegDate(); ok {
		_spec.SetField(tn_bbs.FieldRegDate, field.TypeString, value)
	}
	if value, ok := tuo.mutation.UdtDate(); ok {
		_spec.SetField(tn_bbs.FieldUdtDate, field.TypeString, value)
	}
	_node = &TN_BBS{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tn_bbs.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
