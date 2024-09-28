package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/internal/auth"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

// Resolver is the resolver root.
type Resolver struct{ client *ent.Client }

// NewSchema creates a graphql executable schema.
func NewSchema(client *ent.Client) graphql.ExecutableSchema {
	c := Config{Resolvers: &Resolver{client}}

	c.Directives.Admin = func(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
		if err := auth.AdminRequired(ctx); err != nil {
			return nil, err
		}
		return next(ctx)
	}

	c.Directives.Permission = func(ctx context.Context, obj interface{}, next graphql.Resolver, object string, action string) (res interface{}, err error) {
		if err := auth.EnforceReq(ctx, object, action); err != nil {
			return nil, err
		}
		return next(ctx)
	}

	return NewExecutableSchema(c)
}
