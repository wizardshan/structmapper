// Code generated by entc, DO NOT EDIT.

package order

import (
	"time"
)

const (
	// Label holds the string label denoting the order type in the database.
	Label = "order"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the create_time field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the update_time field in the database.
	FieldUpdateTime = "update_time"
	// FieldUserID holds the string denoting the user_id field in the database.
	FieldUserID = "user_id"
	// FieldShopID holds the string denoting the shop_id field in the database.
	FieldShopID = "shop_id"
	// FieldStatus holds the string denoting the status field in the database.
	FieldStatus = "status"
	// FieldConsignee holds the string denoting the consignee field in the database.
	FieldConsignee = "consignee"
	// FieldMobile holds the string denoting the mobile field in the database.
	FieldMobile = "mobile"
	// FieldProvince holds the string denoting the province field in the database.
	FieldProvince = "province"
	// FieldCity holds the string denoting the city field in the database.
	FieldCity = "city"
	// EdgeUser holds the string denoting the user edge name in mutations.
	EdgeUser = "user"
	// EdgeShop holds the string denoting the shop edge name in mutations.
	EdgeShop = "shop"
	// EdgeItems holds the string denoting the items edge name in mutations.
	EdgeItems = "items"
	// Table holds the table name of the order in the database.
	Table = "orders"
	// UserTable is the table that holds the user relation/edge.
	UserTable = "orders"
	// UserInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UserInverseTable = "users"
	// UserColumn is the table column denoting the user relation/edge.
	UserColumn = "user_id"
	// ShopTable is the table that holds the shop relation/edge.
	ShopTable = "orders"
	// ShopInverseTable is the table name for the Shop entity.
	// It exists in this package in order to avoid circular dependency with the "shop" package.
	ShopInverseTable = "shops"
	// ShopColumn is the table column denoting the shop relation/edge.
	ShopColumn = "shop_id"
	// ItemsTable is the table that holds the items relation/edge.
	ItemsTable = "items"
	// ItemsInverseTable is the table name for the Item entity.
	// It exists in this package in order to avoid circular dependency with the "item" package.
	ItemsInverseTable = "items"
	// ItemsColumn is the table column denoting the items relation/edge.
	ItemsColumn = "order_id"
)

// Columns holds all SQL columns for order fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldUserID,
	FieldShopID,
	FieldStatus,
	FieldConsignee,
	FieldMobile,
	FieldProvince,
	FieldCity,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateTime holds the default value on creation for the "create_time" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "update_time" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "update_time" field.
	UpdateDefaultUpdateTime func() time.Time
)