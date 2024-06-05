package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/ent/publickey"
	"github.com/hawa130/computility-cloud/ent/schema/gqlutils"
	"github.com/hawa130/computility-cloud/ent/schema/mixinx"
	"github.com/hawa130/computility-cloud/internal/hookx"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule"
)

type PublicKey struct {
	ent.Schema
}

func (PublicKey) Fields() []ent.Field {
	return []ent.Field{
		field.String("key").Unique(),
		field.String("name").Optional(),
		field.String("description").Optional(),
		field.String("type").Optional(),
		field.String("status").Optional(),
		field.Time("expired_at").Optional(),
	}
}

func (PublicKey) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("user", User.Type).Unique().Immutable().
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput | entgql.SkipMutationCreateInput),
			),
	}
}

func (PublicKey) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixinx.XId{},
		mixinx.Time{},
	}
}

func (PublicKey) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("expired_at"),
	}
}

func (PublicKey) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		gqlutils.PermissionDirective(publickey.Table, perm.OpRead),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

func (PublicKey) Hooks() []ent.Hook {
	return []ent.Hook{
		hookx.OnCreate.AddObjectGroup(publickey.Table),
		hookx.OnCreate.AddObjectOwner(),
		hookx.OnRemove.RemoveObjectGroupsAndPolicies(),
	}
}

func (PublicKey) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowAuthorizedMutation(publickey.Table),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
