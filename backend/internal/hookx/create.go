package hookx

import (
	"context"

	"entgo.io/ent"
	gen "github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/hook"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/logger"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/rs/xid"
)

var OnCreate = onCreateType{}

type onCreateType struct{}

// AddObjectGroup 创建对象时添加对应资源组
func (onCreateType) AddObjectGroup(model string) ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return HookFunc(func(ctx context.Context, m EntMutation) (gen.Value, error) {
				id, exists := m.ID()
				if !exists {
					logger.Logger().Warn("mutation id not exists")
					return next.Mutate(ctx, m)
				}
				_, err := perm.AddObjectGroup(id.String(), model)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			})
		},
		ent.OpCreate,
	)
}

type customOwnerKey struct{}

// WithCustomOwner 返回一个新的上下文，其中包含自定义的创建者
func WithCustomOwner(parent context.Context, uid xid.ID) context.Context {
	return context.WithValue(parent, customOwnerKey{}, uid)
}

// AddObjectOwner 创建对象时添加创建者
func (onCreateType) AddObjectOwner() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return HookFunc(func(ctx context.Context, m EntMutation) (gen.Value, error) {
				id, exists := m.ID()
				if !exists {
					logger.Logger().Warn("mutation id not exists")
					return next.Mutate(ctx, m)
				}

				uid, ok := ctx.Value(customOwnerKey{}).(xid.ID)
				if !ok {
					owner, ok := auth.FromContext(ctx)
					if !ok {
						return next.Mutate(ctx, m)
					}
					uid = owner.ID
				}

				_, err := perm.GrantObjectPermissionX(uid, id)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			})
		},
		ent.OpCreate,
	)
}
