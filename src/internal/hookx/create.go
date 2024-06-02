package hookx

import (
	"context"

	"entgo.io/ent"
	gen "github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/hook"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/logger"
	"github.com/hawa130/computility-cloud/internal/perm"
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

				_, err := auth.GrantObjectPermissionX(ctx, id)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			})
		},
		ent.OpCreate,
	)
}
