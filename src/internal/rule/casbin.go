package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule/utils"
	"github.com/rs/xid"
)

// AllowAuthorized 允许拥有权限的用户访问
func AllowAuthorized(obj xid.ID, act string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return err
		}

		res, err := perm.EnforceX(user.ID, obj, act)
		if err != nil {
			return err
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
			return err
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
			return err
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
			return err
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
