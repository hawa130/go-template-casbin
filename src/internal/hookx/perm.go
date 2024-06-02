package hookx

import (
	"context"
	"fmt"

	"github.com/casbin/ent-adapter/ent"
	gen "github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/hook"
	"github.com/hawa130/computility-cloud/internal/logger"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/rs/xid"
)

type EntMutation interface {
	gen.Mutation
	ID() (xid.ID, bool)
}

type HookFunc func(context.Context, EntMutation) (ent.Value, error)

func (f HookFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(EntMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect EntMutation", m)
}

// AddObjectGroup 创建对象时添加对应资源组
func AddObjectGroup(model string) ent.Hook {
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

// RemoveObjectGroup 删除对象时删除对应资源组
func RemoveObjectGroup() ent.Hook {
	return hook.On(
		func(next ent.Mutator) ent.Mutator {
			return HookFunc(func(ctx context.Context, m EntMutation) (gen.Value, error) {
				id, exists := m.ID()
				if !exists {
					logger.Logger().Warn("mutation id not exists")
					return next.Mutate(ctx, m)
				}
				_, err := perm.RemoveAllObjectGroupsX(id)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			})
		},
		ent.OpDelete|ent.OpDeleteOne,
	)
}
