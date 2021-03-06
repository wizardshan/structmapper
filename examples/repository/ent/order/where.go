// Code generated by entc, DO NOT EDIT.

package order

import (
	"examples/repository/ent/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// ShopID applies equality check predicate on the "shop_id" field. It's identical to ShopIDEQ.
func ShopID(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShopID), v))
	})
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// Consignee applies equality check predicate on the "consignee" field. It's identical to ConsigneeEQ.
func Consignee(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConsignee), v))
	})
}

// Mobile applies equality check predicate on the "mobile" field. It's identical to MobileEQ.
func Mobile(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobile), v))
	})
}

// Province applies equality check predicate on the "province" field. It's identical to ProvinceEQ.
func Province(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvince), v))
	})
}

// City applies equality check predicate on the "city" field. It's identical to CityEQ.
func City(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCreateTime), v))
	})
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCreateTime), v...))
	})
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCreateTime), v))
	})
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCreateTime), v))
	})
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUpdateTime), v...))
	})
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldUpdateTime), v))
	})
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldUpdateTime), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// UserIDIsNil applies the IsNil predicate on the "user_id" field.
func UserIDIsNil() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldUserID)))
	})
}

// UserIDNotNil applies the NotNil predicate on the "user_id" field.
func UserIDNotNil() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldUserID)))
	})
}

// ShopIDEQ applies the EQ predicate on the "shop_id" field.
func ShopIDEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldShopID), v))
	})
}

// ShopIDNEQ applies the NEQ predicate on the "shop_id" field.
func ShopIDNEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldShopID), v))
	})
}

// ShopIDIn applies the In predicate on the "shop_id" field.
func ShopIDIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldShopID), v...))
	})
}

// ShopIDNotIn applies the NotIn predicate on the "shop_id" field.
func ShopIDNotIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldShopID), v...))
	})
}

// ShopIDIsNil applies the IsNil predicate on the "shop_id" field.
func ShopIDIsNil() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldShopID)))
	})
}

// ShopIDNotNil applies the NotNil predicate on the "shop_id" field.
func ShopIDNotNil() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldShopID)))
	})
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldStatus), v))
	})
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldStatus), v))
	})
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldStatus), v...))
	})
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...int) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldStatus), v...))
	})
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldStatus), v))
	})
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldStatus), v))
	})
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldStatus), v))
	})
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v int) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldStatus), v))
	})
}

// ConsigneeEQ applies the EQ predicate on the "consignee" field.
func ConsigneeEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldConsignee), v))
	})
}

// ConsigneeNEQ applies the NEQ predicate on the "consignee" field.
func ConsigneeNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldConsignee), v))
	})
}

// ConsigneeIn applies the In predicate on the "consignee" field.
func ConsigneeIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldConsignee), v...))
	})
}

// ConsigneeNotIn applies the NotIn predicate on the "consignee" field.
func ConsigneeNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldConsignee), v...))
	})
}

// ConsigneeGT applies the GT predicate on the "consignee" field.
func ConsigneeGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldConsignee), v))
	})
}

// ConsigneeGTE applies the GTE predicate on the "consignee" field.
func ConsigneeGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldConsignee), v))
	})
}

// ConsigneeLT applies the LT predicate on the "consignee" field.
func ConsigneeLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldConsignee), v))
	})
}

// ConsigneeLTE applies the LTE predicate on the "consignee" field.
func ConsigneeLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldConsignee), v))
	})
}

// ConsigneeContains applies the Contains predicate on the "consignee" field.
func ConsigneeContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldConsignee), v))
	})
}

// ConsigneeHasPrefix applies the HasPrefix predicate on the "consignee" field.
func ConsigneeHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldConsignee), v))
	})
}

// ConsigneeHasSuffix applies the HasSuffix predicate on the "consignee" field.
func ConsigneeHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldConsignee), v))
	})
}

// ConsigneeEqualFold applies the EqualFold predicate on the "consignee" field.
func ConsigneeEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldConsignee), v))
	})
}

// ConsigneeContainsFold applies the ContainsFold predicate on the "consignee" field.
func ConsigneeContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldConsignee), v))
	})
}

// MobileEQ applies the EQ predicate on the "mobile" field.
func MobileEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMobile), v))
	})
}

// MobileNEQ applies the NEQ predicate on the "mobile" field.
func MobileNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMobile), v))
	})
}

// MobileIn applies the In predicate on the "mobile" field.
func MobileIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMobile), v...))
	})
}

// MobileNotIn applies the NotIn predicate on the "mobile" field.
func MobileNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMobile), v...))
	})
}

// MobileGT applies the GT predicate on the "mobile" field.
func MobileGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMobile), v))
	})
}

// MobileGTE applies the GTE predicate on the "mobile" field.
func MobileGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMobile), v))
	})
}

// MobileLT applies the LT predicate on the "mobile" field.
func MobileLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMobile), v))
	})
}

// MobileLTE applies the LTE predicate on the "mobile" field.
func MobileLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMobile), v))
	})
}

// MobileContains applies the Contains predicate on the "mobile" field.
func MobileContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMobile), v))
	})
}

// MobileHasPrefix applies the HasPrefix predicate on the "mobile" field.
func MobileHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMobile), v))
	})
}

// MobileHasSuffix applies the HasSuffix predicate on the "mobile" field.
func MobileHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMobile), v))
	})
}

// MobileEqualFold applies the EqualFold predicate on the "mobile" field.
func MobileEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMobile), v))
	})
}

// MobileContainsFold applies the ContainsFold predicate on the "mobile" field.
func MobileContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMobile), v))
	})
}

// ProvinceEQ applies the EQ predicate on the "province" field.
func ProvinceEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProvince), v))
	})
}

// ProvinceNEQ applies the NEQ predicate on the "province" field.
func ProvinceNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProvince), v))
	})
}

// ProvinceIn applies the In predicate on the "province" field.
func ProvinceIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProvince), v...))
	})
}

// ProvinceNotIn applies the NotIn predicate on the "province" field.
func ProvinceNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProvince), v...))
	})
}

// ProvinceGT applies the GT predicate on the "province" field.
func ProvinceGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProvince), v))
	})
}

// ProvinceGTE applies the GTE predicate on the "province" field.
func ProvinceGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProvince), v))
	})
}

// ProvinceLT applies the LT predicate on the "province" field.
func ProvinceLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProvince), v))
	})
}

// ProvinceLTE applies the LTE predicate on the "province" field.
func ProvinceLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProvince), v))
	})
}

// ProvinceContains applies the Contains predicate on the "province" field.
func ProvinceContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProvince), v))
	})
}

// ProvinceHasPrefix applies the HasPrefix predicate on the "province" field.
func ProvinceHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProvince), v))
	})
}

// ProvinceHasSuffix applies the HasSuffix predicate on the "province" field.
func ProvinceHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProvince), v))
	})
}

// ProvinceEqualFold applies the EqualFold predicate on the "province" field.
func ProvinceEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProvince), v))
	})
}

// ProvinceContainsFold applies the ContainsFold predicate on the "province" field.
func ProvinceContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProvince), v))
	})
}

// CityEQ applies the EQ predicate on the "city" field.
func CityEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCity), v))
	})
}

// CityNEQ applies the NEQ predicate on the "city" field.
func CityNEQ(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCity), v))
	})
}

// CityIn applies the In predicate on the "city" field.
func CityIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCity), v...))
	})
}

// CityNotIn applies the NotIn predicate on the "city" field.
func CityNotIn(vs ...string) predicate.Order {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Order(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCity), v...))
	})
}

// CityGT applies the GT predicate on the "city" field.
func CityGT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCity), v))
	})
}

// CityGTE applies the GTE predicate on the "city" field.
func CityGTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCity), v))
	})
}

// CityLT applies the LT predicate on the "city" field.
func CityLT(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCity), v))
	})
}

// CityLTE applies the LTE predicate on the "city" field.
func CityLTE(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCity), v))
	})
}

// CityContains applies the Contains predicate on the "city" field.
func CityContains(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCity), v))
	})
}

// CityHasPrefix applies the HasPrefix predicate on the "city" field.
func CityHasPrefix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCity), v))
	})
}

// CityHasSuffix applies the HasSuffix predicate on the "city" field.
func CityHasSuffix(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCity), v))
	})
}

// CityEqualFold applies the EqualFold predicate on the "city" field.
func CityEqualFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCity), v))
	})
}

// CityContainsFold applies the ContainsFold predicate on the "city" field.
func CityContainsFold(v string) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCity), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasShop applies the HasEdge predicate on the "shop" edge.
func HasShop() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasShopWith applies the HasEdge predicate on the "shop" edge with a given conditions (other predicates).
func HasShopWith(preds ...predicate.Shop) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ShopInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ShopTable, ShopColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasItems applies the HasEdge predicate on the "items" edge.
func HasItems() predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasItemsWith applies the HasEdge predicate on the "items" edge with a given conditions (other predicates).
func HasItemsWith(preds ...predicate.Item) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(ItemsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, ItemsTable, ItemsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Order) predicate.Order {
	return predicate.Order(func(s *sql.Selector) {
		p(s.Not())
	})
}
