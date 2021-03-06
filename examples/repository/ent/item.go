// Code generated by entc, DO NOT EDIT.

package ent

import (
	"examples/repository/ent/item"
	"examples/repository/ent/order"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Item is the model entity for the Item schema.
type Item struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// OrderID holds the value of the "order_id" field.
	OrderID int `json:"order_id,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ItemQuery when eager-loading is set.
	Edges ItemEdges `json:"edges"`
}

// ItemEdges holds the relations/edges for other nodes in the graph.
type ItemEdges struct {
	// Order holds the value of the order edge.
	Order *Order `json:"order,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// OrderOrErr returns the Order value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e ItemEdges) OrderOrErr() (*Order, error) {
	if e.loadedTypes[0] {
		if e.Order == nil {
			// The edge order was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: order.Label}
		}
		return e.Order, nil
	}
	return nil, &NotLoadedError{edge: "order"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Item) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case item.FieldID, item.FieldOrderID, item.FieldPrice:
			values[i] = new(sql.NullInt64)
		case item.FieldTitle:
			values[i] = new(sql.NullString)
		case item.FieldCreateTime, item.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Item", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Item fields.
func (i *Item) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for j := range columns {
		switch columns[j] {
		case item.FieldID:
			value, ok := values[j].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			i.ID = int(value.Int64)
		case item.FieldCreateTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[j])
			} else if value.Valid {
				i.CreateTime = value.Time
			}
		case item.FieldUpdateTime:
			if value, ok := values[j].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[j])
			} else if value.Valid {
				i.UpdateTime = value.Time
			}
		case item.FieldOrderID:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order_id", values[j])
			} else if value.Valid {
				i.OrderID = int(value.Int64)
			}
		case item.FieldTitle:
			if value, ok := values[j].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[j])
			} else if value.Valid {
				i.Title = value.String
			}
		case item.FieldPrice:
			if value, ok := values[j].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field price", values[j])
			} else if value.Valid {
				i.Price = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryOrder queries the "order" edge of the Item entity.
func (i *Item) QueryOrder() *OrderQuery {
	return (&ItemClient{config: i.config}).QueryOrder(i)
}

// Update returns a builder for updating this Item.
// Note that you need to call Item.Unwrap() before calling this method if this Item
// was returned from a transaction, and the transaction was committed or rolled back.
func (i *Item) Update() *ItemUpdateOne {
	return (&ItemClient{config: i.config}).UpdateOne(i)
}

// Unwrap unwraps the Item entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (i *Item) Unwrap() *Item {
	tx, ok := i.config.driver.(*txDriver)
	if !ok {
		panic("ent: Item is not a transactional entity")
	}
	i.config.driver = tx.drv
	return i
}

// String implements the fmt.Stringer.
func (i *Item) String() string {
	var builder strings.Builder
	builder.WriteString("Item(")
	builder.WriteString(fmt.Sprintf("id=%v", i.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(i.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(i.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", order_id=")
	builder.WriteString(fmt.Sprintf("%v", i.OrderID))
	builder.WriteString(", title=")
	builder.WriteString(i.Title)
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", i.Price))
	builder.WriteByte(')')
	return builder.String()
}

// Items is a parsable slice of Item.
type Items []*Item

func (i Items) config(cfg config) {
	for _i := range i {
		i[_i].config = cfg
	}
}
