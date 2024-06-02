package auth

import (
	"context"
	"fmt"

	"github.com/hawa130/computility-cloud/graph/reqerr"
	"github.com/hawa130/computility-cloud/internal/perm"
)

func EnforceCtx(ctx context.Context, obj string, act string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.Enforce(user.ID.String(), obj, act)
}

func EnforceXCtx(ctx context.Context, obj fmt.Stringer, act string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.EnforceX(user.ID, obj, act)
}

func IsAdmin(ctx context.Context) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
}

func IsAdminReq(ctx context.Context) error {
	allow, err := IsAdmin(ctx)
	if err != nil {
		return err
	}
	if !allow {
		return reqerr.ErrForbidden
	}
	return nil
}

func GrantObjectPermission(ctx context.Context, obj string) error {
	user, ok := FromContext(ctx)
	if !ok {
		return nil
	}
	return perm.GrantObjectPermission(user.ID.String(), obj)
}

func GrantObjectPermissionX(ctx context.Context, obj fmt.Stringer) error {
	user, ok := FromContext(ctx)
	if !ok {
		return nil
	}
	return perm.GrantObjectPermissionX(user.ID, obj)
}
