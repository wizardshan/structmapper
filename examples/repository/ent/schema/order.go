package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/mixin"
)

type Order struct {
	ent.Schema
}

func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("user_id").Optional(),
		field.Int("shop_id").Optional(),
		field.Int("status"),
		field.String("consignee"),
		field.String("mobile"),
		field.String("province"),
		field.String("city"),
	}
}

func (Order) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixin.Time{},
	}
}

func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("orders").Field("user_id").
			Unique(),
		edge.From("shop", Shop.Type).
			Ref("orders").Field("shop_id").
			Unique(),
		edge.To("items", Item.Type),
	}
}
