package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.60

import (
	"context"

	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/ent/user"
	"github.com/hawa130/serverx/graph/model"
	"github.com/hawa130/serverx/graph/reqerr"
	"github.com/hawa130/serverx/internal/auth"
	"github.com/hawa130/serverx/internal/perm"
	"github.com/hawa130/serverx/internal/rule"
	"github.com/rs/xid"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return ent.FromContext(ctx).User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id *xid.ID, input ent.UpdateUserInput) (*ent.User, error) {
	id, err := auth.SelfOrAuthenticated(ctx, id, perm.OpUpdate)
	if err != nil {
		return nil, err
	}
	return ent.FromContext(ctx).User.UpdateOneID(*id).SetInput(input).Save(rule.WithAllowContext(ctx))
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id *xid.ID) (bool, error) {
	id, err := auth.SelfOrAuthenticated(ctx, id, perm.OpDelete)
	if err != nil {
		return false, err
	}
	if err := ent.FromContext(ctx).User.DeleteOneID(*id).Exec(rule.WithAllowContext(ctx)); err != nil {
		return false, err
	}
	return true, nil
}

// CreateChildren is the resolver for the createChildren field.
func (r *mutationResolver) CreateChildren(ctx context.Context, id *xid.ID, children []*ent.CreateUserInput) (*ent.User, error) {
	id, err := auth.SelfOrAuthenticated(ctx, id, perm.OpUpdate)
	if err != nil {
		return nil, err
	}

	c := ent.FromContext(ctx)
	builders := make([]*ent.UserCreate, len(children))
	for i, d := range children {
		builders[i] = c.User.Create().SetInput(*d)
	}
	subs, err := c.User.CreateBulk(builders...).Save(rule.WithAllowContext(ctx))
	if err != nil {
		return nil, err
	}

	ids := make([]xid.ID, len(subs))
	for i, sub := range subs {
		ids[i] = sub.ID
	}
	return c.User.UpdateOneID(*id).AddChildIDs(ids...).Save(rule.WithAllowContext(ctx))
}

// RemoveChildren is the resolver for the removeChildren field.
func (r *mutationResolver) RemoveChildren(ctx context.Context, id *xid.ID, child xid.ID) (*ent.User, error) {
	id, err := auth.SelfOrAuthenticated(ctx, id, perm.OpUpdate)
	if err != nil {
		return nil, err
	}

	c := ent.FromContext(ctx)
	err = c.User.DeleteOneID(child).Exec(rule.WithAllowContext(ctx))
	if err != nil {
		return nil, err
	}
	return c.User.UpdateOneID(*id).RemoveChildIDs(child).Save(rule.WithAllowContext(ctx))
}

// UpdatePassword is the resolver for the updatePassword field.
func (r *mutationResolver) UpdatePassword(ctx context.Context, id *xid.ID, input model.UpdatePasswordInput) (*ent.User, error) {
	id, err := auth.SelfOrAuthenticated(ctx, id, perm.OpUpdate)
	if err != nil {
		return nil, err
	}

	c := ent.FromContext(ctx)
	u, err := c.User.Get(ctx, *id)
	if err != nil {
		return nil, err
	}
	ok, err := auth.ComparePasswordAndHash(input.OldPassword, u.Password)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, reqerr.ErrPasswordNotMatch
	}

	return c.User.UpdateOneID(*id).SetPassword(input.NewPassword).Save(ctx)
}

// User 获取指定用户信息。当 ID 为空时获取自己的用户信息。
func (r *queryResolver) User(ctx context.Context, id *xid.ID) (*ent.User, error) {
	u, exists := auth.FromContext(ctx)
	if u == nil && id == nil {
		return nil, reqerr.ErrBadRequest
	}
	if id == nil {
		id = &u.ID
	}

	builder := r.client.User.Query().Where(user.IDEQ(*id))

	if exists {
		// 判断是否有对应读取权限
		allow, err := perm.EnforceX(u.ID, id, perm.OpRead)
		if err != nil {
			return nil, err
		}
		if allow {
			// 如果拥有权限则忽略字段过滤规则
			return builder.Only(rule.WithAllowContext(ctx))
		}
	}

	return builder.Only(ctx)
}

// Children is the resolver for the children field.
func (r *queryResolver) Children(ctx context.Context, id *xid.ID) ([]*ent.User, error) {
	id, err := auth.SelfOrAuthenticated(ctx, id, perm.OpRead)
	if err != nil {
		return nil, err
	}

	return r.client.User.Query().
		Where(user.HasParentWith(user.IDEQ(*id))).
		All(rule.WithAllowContext(ctx))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
