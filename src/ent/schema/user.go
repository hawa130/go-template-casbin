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
	"github.com/hawa130/computility-cloud/ent/role"
	"github.com/hawa130/computility-cloud/ent/schema/mixinx"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/database"
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
		field.String("password").Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("roles", Role.Type),
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
		entgql.QueryField(),
		entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
	}
}

// Hooks of the User.
func (User) Hooks() []ent.Hook {
	return []ent.Hook{
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
					// If the user has no roles, add the user role
					if len(m.RolesIDs()) == 0 {
						userRole, err := database.Client().Role.Query().Where(role.NameEQ("user")).Only(ctx)
						if err != nil && !gen.IsNotFound(err) {
							return nil, err
						}
						if userRole != nil {
							m.AddRoleIDs(userRole.ID)
						}
					}
					return next.Mutate(ctx, m)
				})
			},
			ent.OpCreate,
		),
	}
}

// Policy defines the privacy policy of the User.
func (User) Policy() ent.Policy {
	return privacy.Policy{
		Mutation: privacy.MutationPolicy{
			rule.AllowHasPermission("user:mutate"),
			rule.AllowMutateSelf(),
			privacy.AlwaysDenyRule(),
		},
		Query: privacy.QueryPolicy{
			privacy.AlwaysAllowRule(),
		},
	}
}
