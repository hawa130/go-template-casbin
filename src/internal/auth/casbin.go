package auth

import (
	"context"
	"fmt"

	"github.com/hawa130/computility-cloud/graph/reqerr"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/rs/xid"
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

func EnforceReq(ctx context.Context, obj string, act string) error {
	allow, err := EnforceCtx(ctx, obj, act)
	if err != nil {
		return err
	}
	if !allow {
		return reqerr.ErrForbidden
	}
	return nil
}

func EnforceXReq(ctx context.Context, obj fmt.Stringer, act string) error {
	allow, err := EnforceXCtx(ctx, obj, act)
	if err != nil {
		return err
	}
	if !allow {
		return reqerr.ErrForbidden
	}
	return nil
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

func GrantObjectPermission(ctx context.Context, obj string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.GrantObjectPermission(user.ID.String(), obj)
}

func GrantObjectPermissionX(ctx context.Context, obj fmt.Stringer) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.GrantObjectPermissionX(user.ID, obj)
}

// SelfOrAuthenticated 会判断是否有传入的 ID，如果没有则从上下文中获取用户 ID，如果有则判断是否有权限
func SelfOrAuthenticated(ctx context.Context, id *xid.ID, act string) (*xid.ID, error) {
	if id == nil {
		u, ok := FromContext(ctx)
		if !ok {
			return id, reqerr.ErrForbidden
		}
		id = &u.ID
	} else {
		err := EnforceXReq(ctx, id, act)
		if err != nil {
			return id, err
		}
	}
	return id, nil
}
