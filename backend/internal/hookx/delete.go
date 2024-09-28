package hookx

import (
	"context"

	"entgo.io/ent"
	gen "github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/ent/hook"
	"github.com/hawa130/serverx/internal/logger"
	"github.com/hawa130/serverx/internal/perm"
)

var OnRemove = onRemoveType{}

type onRemoveType struct{}

// RemoveObjectGroupsAndPolicies 删除对象时删除对应资源组和权限
func (onRemoveType) RemoveObjectGroupsAndPolicies() ent.Hook {
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

				_, err = perm.RemoveAllObjectPoliciesX(id)
				if err != nil {
					return nil, err
				}
				return next.Mutate(ctx, m)
			})
		},
		ent.OpDelete|ent.OpDeleteOne,
	)
}
