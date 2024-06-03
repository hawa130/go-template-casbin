// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"github.com/99designs/gqlgen/graphql"
)

func (pk *PublicKey) User(ctx context.Context) (*User, error) {
	result, err := pk.Edges.UserOrErr()
	if IsNotLoaded(err) {
		result, err = pk.QueryUser().Only(ctx)
	}
	return result, MaskNotFound(err)
}

func (u *User) Children(ctx context.Context) (result []*User, err error) {
	if fc := graphql.GetFieldContext(ctx); fc != nil && fc.Field.Alias != "" {
		result, err = u.NamedChildren(graphql.GetFieldContext(ctx).Field.Alias)
	} else {
		result, err = u.Edges.ChildrenOrErr()
	}
	if IsNotLoaded(err) {
		result, err = u.QueryChildren().All(ctx)
	}
	return result, err
}

func (u *User) Parent(ctx context.Context) (*User, error) {
	result, err := u.Edges.ParentOrErr()
	if IsNotLoaded(err) {
		result, err = u.QueryParent().Only(ctx)
	}
	return result, MaskNotFound(err)
}
