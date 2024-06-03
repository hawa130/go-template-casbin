package graph

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/internal/auth"
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
		if err := auth.IsAdminReq(ctx); err != nil {
			return nil, err
		}
		return next(ctx)
	}

	return NewExecutableSchema(c)
}
