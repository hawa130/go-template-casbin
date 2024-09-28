// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/hawa130/serverx/ent/casbinrule"
	"github.com/hawa130/serverx/ent/publickey"
	"github.com/hawa130/serverx/ent/user"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (cr *CasbinRuleQuery) CollectFields(ctx context.Context, satisfies ...string) (*CasbinRuleQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return cr, nil
	}
	if err := cr.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return cr, nil
}

func (cr *CasbinRuleQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(casbinrule.Columns))
		selectedFields = []string{casbinrule.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "ptype":
			if _, ok := fieldSeen[casbinrule.FieldPtype]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldPtype)
				fieldSeen[casbinrule.FieldPtype] = struct{}{}
			}
		case "v0":
			if _, ok := fieldSeen[casbinrule.FieldV0]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldV0)
				fieldSeen[casbinrule.FieldV0] = struct{}{}
			}
		case "v1":
			if _, ok := fieldSeen[casbinrule.FieldV1]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldV1)
				fieldSeen[casbinrule.FieldV1] = struct{}{}
			}
		case "v2":
			if _, ok := fieldSeen[casbinrule.FieldV2]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldV2)
				fieldSeen[casbinrule.FieldV2] = struct{}{}
			}
		case "v3":
			if _, ok := fieldSeen[casbinrule.FieldV3]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldV3)
				fieldSeen[casbinrule.FieldV3] = struct{}{}
			}
		case "v4":
			if _, ok := fieldSeen[casbinrule.FieldV4]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldV4)
				fieldSeen[casbinrule.FieldV4] = struct{}{}
			}
		case "v5":
			if _, ok := fieldSeen[casbinrule.FieldV5]; !ok {
				selectedFields = append(selectedFields, casbinrule.FieldV5)
				fieldSeen[casbinrule.FieldV5] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		cr.Select(selectedFields...)
	}
	return nil
}

type casbinrulePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []CasbinRulePaginateOption
}

func newCasbinRulePaginateArgs(rv map[string]any) *casbinrulePaginateArgs {
	args := &casbinrulePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[whereField].(*CasbinRuleWhereInput); ok {
		args.opts = append(args.opts, WithCasbinRuleFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (pk *PublicKeyQuery) CollectFields(ctx context.Context, satisfies ...string) (*PublicKeyQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return pk, nil
	}
	if err := pk.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return pk, nil
}

func (pk *PublicKeyQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(publickey.Columns))
		selectedFields = []string{publickey.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "user":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: pk.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			pk.withUser = query
		case "createdAt":
			if _, ok := fieldSeen[publickey.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, publickey.FieldCreatedAt)
				fieldSeen[publickey.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[publickey.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, publickey.FieldUpdatedAt)
				fieldSeen[publickey.FieldUpdatedAt] = struct{}{}
			}
		case "key":
			if _, ok := fieldSeen[publickey.FieldKey]; !ok {
				selectedFields = append(selectedFields, publickey.FieldKey)
				fieldSeen[publickey.FieldKey] = struct{}{}
			}
		case "name":
			if _, ok := fieldSeen[publickey.FieldName]; !ok {
				selectedFields = append(selectedFields, publickey.FieldName)
				fieldSeen[publickey.FieldName] = struct{}{}
			}
		case "description":
			if _, ok := fieldSeen[publickey.FieldDescription]; !ok {
				selectedFields = append(selectedFields, publickey.FieldDescription)
				fieldSeen[publickey.FieldDescription] = struct{}{}
			}
		case "type":
			if _, ok := fieldSeen[publickey.FieldType]; !ok {
				selectedFields = append(selectedFields, publickey.FieldType)
				fieldSeen[publickey.FieldType] = struct{}{}
			}
		case "status":
			if _, ok := fieldSeen[publickey.FieldStatus]; !ok {
				selectedFields = append(selectedFields, publickey.FieldStatus)
				fieldSeen[publickey.FieldStatus] = struct{}{}
			}
		case "expiredAt":
			if _, ok := fieldSeen[publickey.FieldExpiredAt]; !ok {
				selectedFields = append(selectedFields, publickey.FieldExpiredAt)
				fieldSeen[publickey.FieldExpiredAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		pk.Select(selectedFields...)
	}
	return nil
}

type publickeyPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []PublicKeyPaginateOption
}

func newPublicKeyPaginateArgs(rv map[string]any) *publickeyPaginateArgs {
	args := &publickeyPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &PublicKeyOrder{Field: &PublicKeyOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithPublicKeyOrder(order))
			}
		case *PublicKeyOrder:
			if v != nil {
				args.opts = append(args.opts, WithPublicKeyOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*PublicKeyWhereInput); ok {
		args.opts = append(args.opts, WithPublicKeyFilter(v.Filter))
	}
	return args
}

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (u *UserQuery) CollectFields(ctx context.Context, satisfies ...string) (*UserQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return u, nil
	}
	if err := u.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return u, nil
}

func (u *UserQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(user.Columns))
		selectedFields = []string{user.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {

		case "children":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, false, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			u.WithNamedChildren(alias, func(wq *UserQuery) {
				*wq = *query
			})

		case "parent":
			var (
				alias = field.Alias
				path  = append(path, alias)
				query = (&UserClient{config: u.config}).Query()
			)
			if err := query.collectField(ctx, oneNode, opCtx, field, path, mayAddCondition(satisfies, userImplementors)...); err != nil {
				return err
			}
			u.withParent = query
		case "createdAt":
			if _, ok := fieldSeen[user.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldCreatedAt)
				fieldSeen[user.FieldCreatedAt] = struct{}{}
			}
		case "updatedAt":
			if _, ok := fieldSeen[user.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, user.FieldUpdatedAt)
				fieldSeen[user.FieldUpdatedAt] = struct{}{}
			}
		case "nickname":
			if _, ok := fieldSeen[user.FieldNickname]; !ok {
				selectedFields = append(selectedFields, user.FieldNickname)
				fieldSeen[user.FieldNickname] = struct{}{}
			}
		case "username":
			if _, ok := fieldSeen[user.FieldUsername]; !ok {
				selectedFields = append(selectedFields, user.FieldUsername)
				fieldSeen[user.FieldUsername] = struct{}{}
			}
		case "email":
			if _, ok := fieldSeen[user.FieldEmail]; !ok {
				selectedFields = append(selectedFields, user.FieldEmail)
				fieldSeen[user.FieldEmail] = struct{}{}
			}
		case "phone":
			if _, ok := fieldSeen[user.FieldPhone]; !ok {
				selectedFields = append(selectedFields, user.FieldPhone)
				fieldSeen[user.FieldPhone] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		u.Select(selectedFields...)
	}
	return nil
}

type userPaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []UserPaginateOption
}

func newUserPaginateArgs(rv map[string]any) *userPaginateArgs {
	args := &userPaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &UserOrder{Field: &UserOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithUserOrder(order))
			}
		case *UserOrder:
			if v != nil {
				args.opts = append(args.opts, WithUserOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*UserWhereInput); ok {
		args.opts = append(args.opts, WithUserFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}
