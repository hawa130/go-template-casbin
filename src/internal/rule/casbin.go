package rule

import (
	"context"
	"fmt"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/logger"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule/utils"
	"github.com/rs/xid"
)

// AllowAuthorized 允许拥有权限的用户访问
func AllowAuthorized(obj xid.ID, act string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}

		res, err := perm.EnforceX(user.ID, obj, act)
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}

		if res {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

// DenyUnauthorized 拒绝未授权的用户访问
func DenyUnauthorized(obj xid.ID, act string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, err := utils.GetUserOrDeny(ctx)
		if err != nil {
			return privacy.Denyf("unauthenticated")
		}
		res, err := perm.EnforceX(user.ID, obj, act)
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}
		if res {
			return privacy.Skip
		}
		return privacy.Denyf("forbidden")
	})
}

// AllowAdmin 允许管理员访问
func AllowAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}
		allow, err := perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
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
		user, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}
		allow, err := perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}
		if allow {
			return privacy.Skip
		}
		return privacy.Denyf("forbidden")
	})
}

type EntMutation interface {
	ent.Mutation
	ID() (xid.ID, bool)
}

type MutationRuleFunc func(context.Context, EntMutation) error

func (f MutationRuleFunc) EvalMutation(ctx context.Context, m ent.Mutation) error {
	if mut, ok := m.(EntMutation); ok {
		return f(ctx, mut)
	}
	return fmt.Errorf("unexpected mutation type %T. expect EntMutation", m)
}

// AllowAuthorizedMutation 允许拥有权限的用户修改
func AllowAuthorizedMutation(model string) privacy.MutationRule {
	return MutationRuleFunc(func(ctx context.Context, m EntMutation) error {
		u, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}

		id, exists := m.ID()
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
