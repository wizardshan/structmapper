// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"examples/repository/ent/item"
	"examples/repository/ent/order"
	"examples/repository/ent/predicate"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ItemUpdate is the builder for updating Item entities.
type ItemUpdate struct {
	config
	hooks    []Hook
	mutation *ItemMutation
}

// Where appends a list predicates to the ItemUpdate builder.
func (iu *ItemUpdate) Where(ps ...predicate.Item) *ItemUpdate {
	iu.mutation.Where(ps...)
	return iu
}

// SetUpdateTime sets the "update_time" field.
func (iu *ItemUpdate) SetUpdateTime(t time.Time) *ItemUpdate {
	iu.mutation.SetUpdateTime(t)
	return iu
}

// SetOrderID sets the "order_id" field.
func (iu *ItemUpdate) SetOrderID(i int) *ItemUpdate {
	iu.mutation.SetOrderID(i)
	return iu
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (iu *ItemUpdate) SetNillableOrderID(i *int) *ItemUpdate {
	if i != nil {
		iu.SetOrderID(*i)
	}
	return iu
}

// ClearOrderID clears the value of the "order_id" field.
func (iu *ItemUpdate) ClearOrderID() *ItemUpdate {
	iu.mutation.ClearOrderID()
	return iu
}

// SetTitle sets the "title" field.
func (iu *ItemUpdate) SetTitle(s string) *ItemUpdate {
	iu.mutation.SetTitle(s)
	return iu
}

// SetPrice sets the "price" field.
func (iu *ItemUpdate) SetPrice(i int) *ItemUpdate {
	iu.mutation.ResetPrice()
	iu.mutation.SetPrice(i)
	return iu
}

// AddPrice adds i to the "price" field.
func (iu *ItemUpdate) AddPrice(i int) *ItemUpdate {
	iu.mutation.AddPrice(i)
	return iu
}

// SetOrder sets the "order" edge to the Order entity.
func (iu *ItemUpdate) SetOrder(o *Order) *ItemUpdate {
	return iu.SetOrderID(o.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iu *ItemUpdate) Mutation() *ItemMutation {
	return iu.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (iu *ItemUpdate) ClearOrder() *ItemUpdate {
	iu.mutation.ClearOrder()
	return iu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (iu *ItemUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	iu.defaults()
	if len(iu.hooks) == 0 {
		affected, err = iu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iu.mutation = mutation
			affected, err = iu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(iu.hooks) - 1; i >= 0; i-- {
			if iu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (iu *ItemUpdate) SaveX(ctx context.Context) int {
	affected, err := iu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (iu *ItemUpdate) Exec(ctx context.Context) error {
	_, err := iu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iu *ItemUpdate) ExecX(ctx context.Context) {
	if err := iu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iu *ItemUpdate) defaults() {
	if _, ok := iu.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iu.mutation.SetUpdateTime(v)
	}
}

func (iu *ItemUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	if ps := iu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldTitle,
		})
	}
	if value, ok := iu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: item.FieldPrice,
		})
	}
	if value, ok := iu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: item.FieldPrice,
		})
	}
	if iu.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.OrderTable,
			Columns: []string{item.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iu.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.OrderTable,
			Columns: []string{item.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, iu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ItemUpdateOne is the builder for updating a single Item entity.
type ItemUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ItemMutation
}

// SetUpdateTime sets the "update_time" field.
func (iuo *ItemUpdateOne) SetUpdateTime(t time.Time) *ItemUpdateOne {
	iuo.mutation.SetUpdateTime(t)
	return iuo
}

// SetOrderID sets the "order_id" field.
func (iuo *ItemUpdateOne) SetOrderID(i int) *ItemUpdateOne {
	iuo.mutation.SetOrderID(i)
	return iuo
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (iuo *ItemUpdateOne) SetNillableOrderID(i *int) *ItemUpdateOne {
	if i != nil {
		iuo.SetOrderID(*i)
	}
	return iuo
}

// ClearOrderID clears the value of the "order_id" field.
func (iuo *ItemUpdateOne) ClearOrderID() *ItemUpdateOne {
	iuo.mutation.ClearOrderID()
	return iuo
}

// SetTitle sets the "title" field.
func (iuo *ItemUpdateOne) SetTitle(s string) *ItemUpdateOne {
	iuo.mutation.SetTitle(s)
	return iuo
}

// SetPrice sets the "price" field.
func (iuo *ItemUpdateOne) SetPrice(i int) *ItemUpdateOne {
	iuo.mutation.ResetPrice()
	iuo.mutation.SetPrice(i)
	return iuo
}

// AddPrice adds i to the "price" field.
func (iuo *ItemUpdateOne) AddPrice(i int) *ItemUpdateOne {
	iuo.mutation.AddPrice(i)
	return iuo
}

// SetOrder sets the "order" edge to the Order entity.
func (iuo *ItemUpdateOne) SetOrder(o *Order) *ItemUpdateOne {
	return iuo.SetOrderID(o.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (iuo *ItemUpdateOne) Mutation() *ItemMutation {
	return iuo.mutation
}

// ClearOrder clears the "order" edge to the Order entity.
func (iuo *ItemUpdateOne) ClearOrder() *ItemUpdateOne {
	iuo.mutation.ClearOrder()
	return iuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (iuo *ItemUpdateOne) Select(field string, fields ...string) *ItemUpdateOne {
	iuo.fields = append([]string{field}, fields...)
	return iuo
}

// Save executes the query and returns the updated Item entity.
func (iuo *ItemUpdateOne) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	iuo.defaults()
	if len(iuo.hooks) == 0 {
		node, err = iuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			iuo.mutation = mutation
			node, err = iuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(iuo.hooks) - 1; i >= 0; i-- {
			if iuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = iuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, iuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (iuo *ItemUpdateOne) SaveX(ctx context.Context) *Item {
	node, err := iuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (iuo *ItemUpdateOne) Exec(ctx context.Context) error {
	_, err := iuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (iuo *ItemUpdateOne) ExecX(ctx context.Context) {
	if err := iuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (iuo *ItemUpdateOne) defaults() {
	if _, ok := iuo.mutation.UpdateTime(); !ok {
		v := item.UpdateDefaultUpdateTime()
		iuo.mutation.SetUpdateTime(v)
	}
}

func (iuo *ItemUpdateOne) sqlSave(ctx context.Context) (_node *Item, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   item.Table,
			Columns: item.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		},
	}
	id, ok := iuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Item.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := iuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, item.FieldID)
		for _, f := range fields {
			if !item.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != item.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := iuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := iuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
	}
	if value, ok := iuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldTitle,
		})
	}
	if value, ok := iuo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: item.FieldPrice,
		})
	}
	if value, ok := iuo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: item.FieldPrice,
		})
	}
	if iuo.mutation.OrderCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.OrderTable,
			Columns: []string{item.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := iuo.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   item.OrderTable,
			Columns: []string{item.OrderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Item{config: iuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, iuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{item.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
