// Code generated by entc, DO NOT EDIT.

package ent

import (
	"examples/repository/ent/order"
	"examples/repository/ent/shop"
	"examples/repository/ent/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime time.Time `json:"update_time,omitempty"`
	// UserID holds the value of the "user_id" field.
	UserID int `json:"user_id,omitempty"`
	// ShopID holds the value of the "shop_id" field.
	ShopID int `json:"shop_id,omitempty"`
	// Status holds the value of the "status" field.
	Status int `json:"status,omitempty"`
	// Consignee holds the value of the "consignee" field.
	Consignee string `json:"consignee,omitempty"`
	// Mobile holds the value of the "mobile" field.
	Mobile string `json:"mobile,omitempty"`
	// Province holds the value of the "province" field.
	Province string `json:"province,omitempty"`
	// City holds the value of the "city" field.
	City string `json:"city,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges OrderEdges `json:"edges"`
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Shop holds the value of the shop edge.
	Shop *Shop `json:"shop,omitempty"`
	// Items holds the value of the items edge.
	Items []*Item `json:"items,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// ShopOrErr returns the Shop value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) ShopOrErr() (*Shop, error) {
	if e.loadedTypes[1] {
		if e.Shop == nil {
			// The edge shop was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: shop.Label}
		}
		return e.Shop, nil
	}
	return nil, &NotLoadedError{edge: "shop"}
}

// ItemsOrErr returns the Items value or an error if the edge
// was not loaded in eager-loading.
func (e OrderEdges) ItemsOrErr() ([]*Item, error) {
	if e.loadedTypes[2] {
		return e.Items, nil
	}
	return nil, &NotLoadedError{edge: "items"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case order.FieldID, order.FieldUserID, order.FieldShopID, order.FieldStatus:
			values[i] = new(sql.NullInt64)
		case order.FieldConsignee, order.FieldMobile, order.FieldProvince, order.FieldCity:
			values[i] = new(sql.NullString)
		case order.FieldCreateTime, order.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Order", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case order.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			o.ID = int(value.Int64)
		case order.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				o.CreateTime = value.Time
			}
		case order.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				o.UpdateTime = value.Time
			}
		case order.FieldUserID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field user_id", values[i])
			} else if value.Valid {
				o.UserID = int(value.Int64)
			}
		case order.FieldShopID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field shop_id", values[i])
			} else if value.Valid {
				o.ShopID = int(value.Int64)
			}
		case order.FieldStatus:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				o.Status = int(value.Int64)
			}
		case order.FieldConsignee:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field consignee", values[i])
			} else if value.Valid {
				o.Consignee = value.String
			}
		case order.FieldMobile:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mobile", values[i])
			} else if value.Valid {
				o.Mobile = value.String
			}
		case order.FieldProvince:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field province", values[i])
			} else if value.Valid {
				o.Province = value.String
			}
		case order.FieldCity:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field city", values[i])
			} else if value.Valid {
				o.City = value.String
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Order entity.
func (o *Order) QueryUser() *UserQuery {
	return (&OrderClient{config: o.config}).QueryUser(o)
}

// QueryShop queries the "shop" edge of the Order entity.
func (o *Order) QueryShop() *ShopQuery {
	return (&OrderClient{config: o.config}).QueryShop(o)
}

// QueryItems queries the "items" edge of the Order entity.
func (o *Order) QueryItems() *ItemQuery {
	return (&OrderClient{config: o.config}).QueryItems(o)
}

// Update returns a builder for updating this Order.
// Note that you need to call Order.Unwrap() before calling this method if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the Order entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", create_time=")
	builder.WriteString(o.CreateTime.Format(time.ANSIC))
	builder.WriteString(", update_time=")
	builder.WriteString(o.UpdateTime.Format(time.ANSIC))
	builder.WriteString(", user_id=")
	builder.WriteString(fmt.Sprintf("%v", o.UserID))
	builder.WriteString(", shop_id=")
	builder.WriteString(fmt.Sprintf("%v", o.ShopID))
	builder.WriteString(", status=")
	builder.WriteString(fmt.Sprintf("%v", o.Status))
	builder.WriteString(", consignee=")
	builder.WriteString(o.Consignee)
	builder.WriteString(", mobile=")
	builder.WriteString(o.Mobile)
	builder.WriteString(", province=")
	builder.WriteString(o.Province)
	builder.WriteString(", city=")
	builder.WriteString(o.City)
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
