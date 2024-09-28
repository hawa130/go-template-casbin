package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.54

import (
	"context"

	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/internal/rule"
	"github.com/rs/xid"
)

// ResetPassword is the resolver for the resetPassword field.
func (r *mutationResolver) ResetPassword(ctx context.Context, id xid.ID, password string) (*ent.User, error) {
	return ent.FromContext(ctx).User.
		UpdateOneID(id).
		SetPassword(password).
		Save(rule.WithAllowContext(ctx))
}
