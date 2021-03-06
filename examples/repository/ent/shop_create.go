// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"examples/repository/ent/order"
	"examples/repository/ent/shop"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ShopCreate is the builder for creating a Shop entity.
type ShopCreate struct {
	config
	mutation *ShopMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetCreateTime sets the "create_time" field.
func (sc *ShopCreate) SetCreateTime(t time.Time) *ShopCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *ShopCreate) SetNillableCreateTime(t *time.Time) *ShopCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *ShopCreate) SetUpdateTime(t time.Time) *ShopCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *ShopCreate) SetNillableUpdateTime(t *time.Time) *ShopCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *ShopCreate) SetName(s string) *ShopCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sc *ShopCreate) SetNillableName(s *string) *ShopCreate {
	if s != nil {
		sc.SetName(*s)
	}
	return sc
}

// AddOrderIDs adds the "orders" edge to the Order entity by IDs.
func (sc *ShopCreate) AddOrderIDs(ids ...int) *ShopCreate {
	sc.mutation.AddOrderIDs(ids...)
	return sc
}

// AddOrders adds the "orders" edges to the Order entity.
func (sc *ShopCreate) AddOrders(o ...*Order) *ShopCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return sc.AddOrderIDs(ids...)
}

// Mutation returns the ShopMutation object of the builder.
func (sc *ShopCreate) Mutation() *ShopMutation {
	return sc.mutation
}

// Save creates the Shop in the database.
func (sc *ShopCreate) Save(ctx context.Context) (*Shop, error) {
	var (
		err  error
		node *Shop
	)
	sc.defaults()
	if len(sc.hooks) == 0 {
		if err = sc.check(); err != nil {
			return nil, err
		}
		node, err = sc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ShopMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = sc.check(); err != nil {
				return nil, err
			}
			sc.mutation = mutation
			if node, err = sc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(sc.hooks) - 1; i >= 0; i-- {
			if sc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = sc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, sc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (sc *ShopCreate) SaveX(ctx context.Context) *Shop {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *ShopCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *ShopCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *ShopCreate) defaults() {
	if _, ok := sc.mutation.CreateTime(); !ok {
		v := shop.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		v := shop.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.Name(); !ok {
		v := shop.DefaultName
		sc.mutation.SetName(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *ShopCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`ent: missing required field "Shop.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`ent: missing required field "Shop.update_time"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Shop.name"`)}
	}
	return nil
}

func (sc *ShopCreate) sqlSave(ctx context.Context) (*Shop, error) {
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (sc *ShopCreate) createSpec() (*Shop, *sqlgraph.CreateSpec) {
	var (
		_node = &Shop{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: shop.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: shop.FieldID,
			},
		}
	)
	_spec.OnConflict = sc.conflict
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: shop.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: shop.FieldName,
		})
		_node.Name = value
	}
	if nodes := sc.mutation.OrdersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   shop.OrdersTable,
			Columns: []string{shop.OrdersColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Shop.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShopUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (sc *ShopCreate) OnConflict(opts ...sql.ConflictOption) *ShopUpsertOne {
	sc.conflict = opts
	return &ShopUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (sc *ShopCreate) OnConflictColumns(columns ...string) *ShopUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &ShopUpsertOne{
		create: sc,
	}
}

type (
	// ShopUpsertOne is the builder for "upsert"-ing
	//  one Shop node.
	ShopUpsertOne struct {
		create *ShopCreate
	}

	// ShopUpsert is the "OnConflict" setter.
	ShopUpsert struct {
		*sql.UpdateSet
	}
)

// SetCreateTime sets the "create_time" field.
func (u *ShopUpsert) SetCreateTime(v time.Time) *ShopUpsert {
	u.Set(shop.FieldCreateTime, v)
	return u
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ShopUpsert) UpdateCreateTime() *ShopUpsert {
	u.SetExcluded(shop.FieldCreateTime)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ShopUpsert) SetUpdateTime(v time.Time) *ShopUpsert {
	u.Set(shop.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ShopUpsert) UpdateUpdateTime() *ShopUpsert {
	u.SetExcluded(shop.FieldUpdateTime)
	return u
}

// SetName sets the "name" field.
func (u *ShopUpsert) SetName(v string) *ShopUpsert {
	u.Set(shop.FieldName, v)
	return u
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ShopUpsert) UpdateName() *ShopUpsert {
	u.SetExcluded(shop.FieldName)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ShopUpsertOne) UpdateNewValues() *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(shop.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.Shop.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *ShopUpsertOne) Ignore() *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShopUpsertOne) DoNothing() *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShopCreate.OnConflict
// documentation for more info.
func (u *ShopUpsertOne) Update(set func(*ShopUpsert)) *ShopUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShopUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ShopUpsertOne) SetCreateTime(v time.Time) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateCreateTime() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ShopUpsertOne) SetUpdateTime(v time.Time) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateUpdateTime() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *ShopUpsertOne) SetName(v string) *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ShopUpsertOne) UpdateName() *ShopUpsertOne {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ShopUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShopCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShopUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ShopUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ShopUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ShopCreateBulk is the builder for creating many Shop entities in bulk.
type ShopCreateBulk struct {
	config
	builders []*ShopCreate
	conflict []sql.ConflictOption
}

// Save creates the Shop entities in the database.
func (scb *ShopCreateBulk) Save(ctx context.Context) ([]*Shop, error) {
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Shop, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ShopMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *ShopCreateBulk) SaveX(ctx context.Context) []*Shop {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *ShopCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *ShopCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Shop.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ShopUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
//
func (scb *ShopCreateBulk) OnConflict(opts ...sql.ConflictOption) *ShopUpsertBulk {
	scb.conflict = opts
	return &ShopUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (scb *ShopCreateBulk) OnConflictColumns(columns ...string) *ShopUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &ShopUpsertBulk{
		create: scb,
	}
}

// ShopUpsertBulk is the builder for "upsert"-ing
// a bulk of Shop nodes.
type ShopUpsertBulk struct {
	create *ShopCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
//
func (u *ShopUpsertBulk) UpdateNewValues() *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(shop.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Shop.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *ShopUpsertBulk) Ignore() *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ShopUpsertBulk) DoNothing() *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ShopCreateBulk.OnConflict
// documentation for more info.
func (u *ShopUpsertBulk) Update(set func(*ShopUpsert)) *ShopUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ShopUpsert{UpdateSet: update})
	}))
	return u
}

// SetCreateTime sets the "create_time" field.
func (u *ShopUpsertBulk) SetCreateTime(v time.Time) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetCreateTime(v)
	})
}

// UpdateCreateTime sets the "create_time" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateCreateTime() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateCreateTime()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ShopUpsertBulk) SetUpdateTime(v time.Time) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateUpdateTime() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetName sets the "name" field.
func (u *ShopUpsertBulk) SetName(v string) *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.SetName(v)
	})
}

// UpdateName sets the "name" field to the value that was provided on create.
func (u *ShopUpsertBulk) UpdateName() *ShopUpsertBulk {
	return u.Update(func(s *ShopUpsert) {
		s.UpdateName()
	})
}

// Exec executes the query.
func (u *ShopUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the ShopCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for ShopCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ShopUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
