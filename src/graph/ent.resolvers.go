package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.47

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/rs/xid"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id xid.ID) (ent.Noder, error) {
	err := auth.IsAdminReq(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.Noder(ctx, id)
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []xid.ID) ([]ent.Noder, error) {
	err := auth.IsAdminReq(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.Noders(ctx, ids)
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[xid.ID], first *int, before *entgql.Cursor[xid.ID], last *int, orderBy *ent.UserOrder, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	err := auth.IsAdminReq(ctx)
	if err != nil {
		return nil, err
	}
	return r.client.User.Query().
		Paginate(ctx, after, first, before, last,
			ent.WithUserOrder(orderBy),
			ent.WithUserFilter(where.Filter),
		)
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
