package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/permission"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/ent/role"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/database"
)

func AllowAdmin() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, ok := auth.FromContext(ctx)
		if !ok || user == nil {
			return privacy.Skipf("unauthenticated")
		}

		exists, _ := user.QueryRoles().Where(role.NameIn("admin", "root")).Exist(database.AllowContext)
		if exists {
			return privacy.Allow
		}

		return privacy.Skip
	})
}

func AllowHasPermission(name ...string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, ok := auth.FromContext(ctx)
		if !ok || user == nil {
			return privacy.Skipf("unauthenticated")
		}
		if hasPermissionsIn(user, name...) {
			return privacy.Allow
		}
		return privacy.Skip
	})
}

func DenyNotHasPermission(name ...string) privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, ok := auth.FromContext(ctx)
		if !ok || user == nil {
			return privacy.Skipf("unauthenticated")
		}
		if hasPermissionsIn(user, name...) {
			return privacy.Skip
		}
		return privacy.Deny
	})
}

func hasPermissionsIn(user *ent.User, name ...string) bool {
	exists, err := user.
		QueryRoles().
		Where(role.HasPermissionsWith(permission.NameIn(name...))).
		Exist(database.AllowContext)

	if err != nil {
		return false
	}
	return exists
}
