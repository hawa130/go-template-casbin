package rule

import (
	"context"
	"fmt"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/logger"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule/utils"
	"github.com/rs/xid"
)

// AllowAdmin 允许管理员访问
func AllowAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return err
		}
		allow, err := perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
		if err != nil {
			return fmt.Errorf("unexpected error %v", err)
		}
		if allow {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// DenyNonAdmin 拒绝非管理员访问
func DenyNonAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, err := utils.GetUserOrDeny(ctx)
		if err != nil {
			return err
		}
		allow, err := perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
		if err != nil {
			return fmt.Errorf("unexpected error %v", err)
		}
		if allow {
			return privacy.Skip
		}
		return privacy.Denyf("forbidden")
	})
}

// AllowPermission 允许拥有权限的用户访问
func AllowPermission(model, act string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		allow, err := auth.EnforceCtx(ctx, model, act)
		if err != nil {
			return fmt.Errorf("unexpected error %v", err)
		}
		if allow {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// DenyNonPermission 拒绝未授权的用户访问
func DenyNonPermission(model, act string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		allow, err := auth.EnforceCtx(ctx, model, act)
		if err != nil {
			return fmt.Errorf("unexpected error %v", err)
		}
		if allow {
			return privacy.Skip
		}
		return privacy.Denyf("forbidden")
	})
}

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
