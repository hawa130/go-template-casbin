package mixinx

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"entgo.io/ent/schema/mixin"
)

type CreateTime struct {
	mixin.Schema
}

func (CreateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now).
			Immutable().
			Annotations(
				entgql.OrderField("CREATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput),
			),
	}
}

func (CreateTime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("created_at"),
	}
}

type UpdateTime struct {
	mixin.Schema
}

func (UpdateTime) Fields() []ent.Field {
	return []ent.Field{
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now).
			Annotations(
				entgql.OrderField("UPDATED_AT"),
				entgql.Skip(entgql.SkipMutationCreateInput|entgql.SkipMutationUpdateInput),
			),
	}
}

func (UpdateTime) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("updated_at"),
	}
}

type Time struct {
	mixin.Schema
}

func (Time) Fields() []ent.Field {
	return append(
		CreateTime{}.Fields(),
		UpdateTime{}.Fields()...,
	)
}

func (Time) Indexes() []ent.Index {
	return append(
		CreateTime{}.Indexes(),
		UpdateTime{}.Indexes()...,
	)
}
