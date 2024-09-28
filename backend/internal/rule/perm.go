package rule

import (
	"context"
	"fmt"

	"github.com/hawa130/serverx/ent/privacy"
	"github.com/hawa130/serverx/internal/auth"
	"github.com/hawa130/serverx/internal/perm"
	"github.com/hawa130/serverx/internal/rule/utils"
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
