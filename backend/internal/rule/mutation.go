package rule

import (
	"context"
	"fmt"

	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/ent/privacy"
	"github.com/hawa130/serverx/internal/logger"
	"github.com/hawa130/serverx/internal/perm"
	"github.com/hawa130/serverx/internal/rule/utils"
	"github.com/rs/xid"
)

// AllowAuthorizedMutation 允许拥有权限的用户修改
func AllowAuthorizedMutation(model string) privacy.MutationRule {
	type EntMutation interface {
		ID() (xid.ID, bool)
	}
	return privacy.MutationRuleFunc(func(ctx context.Context, m ent.Mutation) error {
		u, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return err
		}

		mut, ok := m.(EntMutation)
		if !ok {
			return fmt.Errorf("unexpected mutation type %T", m)
		}
		id, exists := mut.ID()
		if !exists {
			logger.Logger().Warn("mutation id not exists")
			return privacy.Skip
		}

		var allow bool
		if m.Op() == ent.OpCreate {
			allow, err = perm.Enforce(u.ID.String(), model, perm.OpCreate)
		} else {
			allow, err = perm.EnforceX(u.ID, id, utils.ToPermOp(m.Op()))
		}
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}
		if allow {
			return privacy.Allow
		}

		return privacy.Skip
	})
}
