// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"examples/repository/ent/item"
	"examples/repository/ent/order"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ItemCreate is the builder for creating a Item entity.
type ItemCreate struct {
	config
	mutation *ItemMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (ic *ItemCreate) SetCreateTime(t time.Time) *ItemCreate {
	ic.mutation.SetCreateTime(t)
	return ic
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ic *ItemCreate) SetNillableCreateTime(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetCreateTime(*t)
	}
	return ic
}

// SetUpdateTime sets the "update_time" field.
func (ic *ItemCreate) SetUpdateTime(t time.Time) *ItemCreate {
	ic.mutation.SetUpdateTime(t)
	return ic
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ic *ItemCreate) SetNillableUpdateTime(t *time.Time) *ItemCreate {
	if t != nil {
		ic.SetUpdateTime(*t)
	}
	return ic
}

// SetOrderID sets the "order_id" field.
func (ic *ItemCreate) SetOrderID(i int) *ItemCreate {
	ic.mutation.SetOrderID(i)
	return ic
}

// SetNillableOrderID sets the "order_id" field if the given value is not nil.
func (ic *ItemCreate) SetNillableOrderID(i *int) *ItemCreate {
	if i != nil {
		ic.SetOrderID(*i)
	}
	return ic
}

// SetTitle sets the "title" field.
func (ic *ItemCreate) SetTitle(s string) *ItemCreate {
	ic.mutation.SetTitle(s)
	return ic
}

// SetPrice sets the "price" field.
func (ic *ItemCreate) SetPrice(i int) *ItemCreate {
	ic.mutation.SetPrice(i)
	return ic
}

// SetOrder sets the "order" edge to the Order entity.
func (ic *ItemCreate) SetOrder(o *Order) *ItemCreate {
	return ic.SetOrderID(o.ID)
}

// Mutation returns the ItemMutation object of the builder.
func (ic *ItemCreate) Mutation() *ItemMutation {
	return ic.mutation
}

// Save creates the Item in the database.
func (ic *ItemCreate) Save(ctx context.Context) (*Item, error) {
	var (
		err  error
		node *Item
	)
	ic.defaults()
	if len(ic.hooks) == 0 {
		if err = ic.check(); err != nil {
			return nil, err
		}
		node, err = ic.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ItemMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ic.check(); err != nil {
				return nil, err
			}
			ic.mutation = mutation
			if node, err = ic.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ic.hooks) - 1; i >= 0; i-- {
			if ic.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ic.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ic.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ic *ItemCreate) SaveX(ctx context.Context) *Item {
	v, err := ic.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ic *ItemCreate) Exec(ctx context.Context) error {
	_, err := ic.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ic *ItemCreate) ExecX(ctx context.Context) {
	if err := ic.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ic *ItemCreate) defaults() {
	if _, ok := ic.mutation.CreateTime(); !ok {
		v := item.DefaultCreateTime()
		ic.mutation.SetCreateTime(v)
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		v := item.DefaultUpdateTime()
		ic.mutation.SetUpdateTime(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ic *ItemCreate) check() error {
	if _, ok := ic.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Item.create_time"`)}
	}
	if _, ok := ic.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Item.update_time"`)}
	}
	if _, ok := ic.mutation.Title(); !ok {
		return &ValidationError{Name: "title", err: errors.New(`ent: missing required field "Item.title"`)}
	}
	if _, ok := ic.mutation.Price(); !ok {
		return &ValidationError{Name: "price", err: errors.New(`ent: missing required field "Item.price"`)}
	}
	return nil
}

func (ic *ItemCreate) sqlSave(ctx context.Context) (*Item, error) {
	_node, _spec := ic.createSpec()
	if err := sqlgraph.CreateNode(ctx, ic.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (ic *ItemCreate) createSpec() (*Item, *sqlgraph.CreateSpec) {
	var (
		_node = &Item{config: ic.config}
		_spec = &sqlgraph.CreateSpec{
			Table: item.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: item.FieldID,
			},
		}
	)
	_spec.OnConflict = ic.conflict
	if value, ok := ic.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ic.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: item.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ic.mutation.Title(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: item.FieldTitle,
		})
		_node.Title = value
	}
	if value, ok := ic.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: item.FieldPrice,
		})
		_node.Price = value
	}
	if nodes := ic.mutation.OrderIDs(); len(nodes) > 0 {
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
		_node.OrderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Item.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (ic *ItemCreate) OnConflict(opts ...sql.ConflictOption) *ItemUpsertOne {
	ic.conflict = opts
	return &ItemUpsertOne{
		create: ic,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ic *ItemCreate) OnConflictColumns(columns ...string) *ItemUpsertOne {
	ic.conflict = append(ic.conflict, sql.ConflictColumns(columns...))
	return &ItemUpsertOne{
		create: ic,
	}
}

type (
	// ItemUpsertOne is the builder for "upsert"-ing
	//  one Item node.
	ItemUpsertOne struct {
		create *ItemCreate
	}

	// ItemUpsert is the "OnConflict" setter.
	ItemUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *ItemUpsert) SetCreateTime(v time.Time) *ItemUpsert {
	u.Set(item.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ItemUpsert) UpdateCreateTime() *ItemUpsert {
	u.SetExcluded(item.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ItemUpsert) SetUpdateTime(v time.Time) *ItemUpsert {
	u.Set(item.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ItemUpsert) UpdateUpdateTime() *ItemUpsert {
	u.SetExcluded(item.FieldUpdateTime)
	return u
}

// SetOrderID sets the "order_id" field.
func (u *ItemUpsert) SetOrderID(v int) *ItemUpsert {
	u.Set(item.FieldOrderID, v)
	return u
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *ItemUpsert) UpdateOrderID() *ItemUpsert {
	u.SetExcluded(item.FieldOrderID)
	return u
}

// ClearOrderID clears the value of the "order_id" field.
func (u *ItemUpsert) ClearOrderID() *ItemUpsert {
	u.SetNull(item.FieldOrderID)
	return u
}

// SetTitle sets the "title" field.
func (u *ItemUpsert) SetTitle(v string) *ItemUpsert {
	u.Set(item.FieldTitle, v)
	return u
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ItemUpsert) UpdateTitle() *ItemUpsert {
	u.SetExcluded(item.FieldTitle)
	return u
}

// SetPrice sets the "price" field.
func (u *ItemUpsert) SetPrice(v int) *ItemUpsert {
	u.Set(item.FieldPrice, v)
	return u
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ItemUpsert) UpdatePrice() *ItemUpsert {
	u.SetExcluded(item.FieldPrice)
	return u
}

// AddPrice adds v to the "price" field.
func (u *ItemUpsert) AddPrice(v int) *ItemUpsert {
	u.Add(item.FieldPrice, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ItemUpsertOne) UpdateNewValues() *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(item.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Item.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ItemUpsertOne) Ignore() *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemUpsertOne) DoNothing() *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemCreate.OnConflict
// documentation for more info.
func (u *ItemUpsertOne) Update(set func(*ItemUpsert)) *ItemUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ItemUpsertOne) SetCreateTime(v time.Time) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateCreateTime() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ItemUpsertOne) SetUpdateTime(v time.Time) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateUpdateTime() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetOrderID sets the "order_id" field.
func (u *ItemUpsertOne) SetOrderID(v int) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateOrderID() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *ItemUpsertOne) ClearOrderID() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.ClearOrderID()
	})
}

// SetTitle sets the "title" field.
func (u *ItemUpsertOne) SetTitle(v string) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdateTitle() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateTitle()
	})
}

// SetPrice sets the "price" field.
func (u *ItemUpsertOne) SetPrice(v int) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *ItemUpsertOne) AddPrice(v int) *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ItemUpsertOne) UpdatePrice() *ItemUpsertOne {
	return u.Update(func(s *ItemUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *ItemUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ItemUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ItemUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ItemCreateBulk is the builder for creating many Item entities in bulk.
type ItemCreateBulk struct {
	config
	builders []*ItemCreate
	conflict []sql.ConflictOption
}

// Save creates the Item entities in the database.
func (icb *ItemCreateBulk) Save(ctx context.Context) ([]*Item, error) {
	specs := make([]*sqlgraph.CreateSpec, len(icb.builders))
	nodes := make([]*Item, len(icb.builders))
	mutators := make([]Mutator, len(icb.builders))
	for i := range icb.builders {
		func(i int, root context.Context) {
			builder := icb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ItemMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, icb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = icb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, icb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, icb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (icb *ItemCreateBulk) SaveX(ctx context.Context) []*Item {
	v, err := icb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (icb *ItemCreateBulk) Exec(ctx context.Context) error {
	_, err := icb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (icb *ItemCreateBulk) ExecX(ctx context.Context) {
	if err := icb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Item.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ItemUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (icb *ItemCreateBulk) OnConflict(opts ...sql.ConflictOption) *ItemUpsertBulk {
	icb.conflict = opts
	return &ItemUpsertBulk{
		create: icb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (icb *ItemCreateBulk) OnConflictColumns(columns ...string) *ItemUpsertBulk {
	icb.conflict = append(icb.conflict, sql.ConflictColumns(columns...))
	return &ItemUpsertBulk{
		create: icb,
	}
}

// ItemUpsertBulk is the builder for "upsert"-ing
// a bulk of Item nodes.
type ItemUpsertBulk struct {
	create *ItemCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ItemUpsertBulk) UpdateNewValues() *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(item.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Item.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ItemUpsertBulk) Ignore() *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ItemUpsertBulk) DoNothing() *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ItemCreateBulk.OnConflict
// documentation for more info.
func (u *ItemUpsertBulk) Update(set func(*ItemUpsert)) *ItemUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ItemUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ItemUpsertBulk) SetCreateTime(v time.Time) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateCreateTime() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ItemUpsertBulk) SetUpdateTime(v time.Time) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateUpdateTime() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetOrderID sets the "order_id" field.
func (u *ItemUpsertBulk) SetOrderID(v int) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetOrderID(v)
	})
}

// UpdateOrderID sets the "order_id" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateOrderID() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateOrderID()
	})
}

// ClearOrderID clears the value of the "order_id" field.
func (u *ItemUpsertBulk) ClearOrderID() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.ClearOrderID()
	})
}

// SetTitle sets the "title" field.
func (u *ItemUpsertBulk) SetTitle(v string) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetTitle(v)
	})
}

// UpdateTitle sets the "title" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdateTitle() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdateTitle()
	})
}

// SetPrice sets the "price" field.
func (u *ItemUpsertBulk) SetPrice(v int) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.SetPrice(v)
	})
}

// AddPrice adds v to the "price" field.
func (u *ItemUpsertBulk) AddPrice(v int) *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.AddPrice(v)
	})
}

// UpdatePrice sets the "price" field to the value that was provided on create.
func (u *ItemUpsertBulk) UpdatePrice() *ItemUpsertBulk {
	return u.Update(func(s *ItemUpsert) {
		s.UpdatePrice()
	})
}

// Exec executes the query.
func (u *ItemUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ItemCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ItemCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ItemUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
