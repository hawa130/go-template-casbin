package schema

import (
	"context"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	gen "github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/hook"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/ent/schema/gqlutils"
	"github.com/hawa130/computility-cloud/ent/schema/mixinx"
	"github.com/hawa130/computility-cloud/ent/user"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/hookx"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("nickname").Optional(),
		field.String("username").Unique().Optional(),
		field.String("email").Unique().Optional(),
		field.String("phone").Unique(),
		field.String("password").
			Sensitive().
			Annotations(entgql.Skip(entgql.SkipMutationUpdateInput)),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("parent", User.Type).
			Unique().
			Immutable().
			From("children").
			Annotations(
				entgql.Skip(entgql.SkipMutationUpdateInput | entgql.SkipMutationCreateInput),
			),
	}
}

// Mixin of the User.
func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixinx.XId{},
		mixinx.Time{},
	}
}

// Annotations of the User.
func (User) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entgql.RelayConnection(),
		gqlutils.PermissionDirective(user.Table, perm.OpRead),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
		hookx.OnCreate.AddObjectGroup(user.Table),
		hookx.OnRemove.RemoveObjectGroupsAndPolicies(),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (gen.Value, error) {
					// Hash the password if mutation has password field
					password, exists := m.Password()
					if !exists {
						return next.Mutate(ctx, m)
					}
					hashed, err := auth.HashPassword(password)
					if err != nil {
						return nil, err
					}
					m.SetPassword(hashed)
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (gen.Value, error) {
					id, exists := m.ID()
					if !exists {
						return next.Mutate(ctx, m)
					}
					// 更新用户权限组，拥有自己的权限
					_, err := perm.GrantObjectPermissionX(id, id)
					if err != nil {
						return nil, err
					}
					// 如果有主用户则添加对应资源组
					parId, exists := m.ParentID()
					if !exists {
						return next.Mutate(ctx, m)
					}
					_, err = perm.AddObjectGroupX(id, parId)
					if err != nil {
						return nil, err
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (gen.Value, error) {
					id, exists := m.ID()
					if !exists {
						return next.Mutate(ctx, m)
					}
					// 更新子用户资源组，隶属于主用户
					ids := m.ChildrenIDs()
					for _, subId := range ids {
						_, err := perm.AddObjectGroupX(subId, id)
						if err != nil {
							return nil, err
						}
					}
					removedIds := m.RemovedChildrenIDs()
					for _, subId := range removedIds {
						_, err := perm.RemoveSubjectRoleX(subId, id)
						if err != nil {
							return nil, err
						}
						_, err = perm.RemoveObjectGroupX(subId, id)
						if err != nil {
							return nil, err
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate|ent.OpUpdate|ent.OpUpdateOne,
		),
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.UserFunc(func(ctx context.Context, m *gen.UserMutation) (gen.Value, error) {
					id, exists := m.ID()
					if !exists {
						return next.Mutate(ctx, m)
					}
					// 删除用户权限组
					_, err := perm.RemoveAllSubjectRolesX(id)
					if err != nil {
						return nil, err
					}
					// 删除用户所有权限
					_, err = perm.RevokeAllPermissionsX(id)
					if err != nil {
						return nil, err
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpDelete|ent.OpDeleteOne,
		),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowAuthorizedMutation(user.Table),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			rule.AllowPermission(user.Table, perm.OpRead),
			rule.LimitUserQueryFields(
				user.FieldID,
				user.FieldNickname,
				user.FieldUsername,
				user.FieldCreatedAt,
			),
			privacy.AlwaysAllowRule(),
		},
	}
}
