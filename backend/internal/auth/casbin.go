package auth

import (
	"context"
	"fmt"

	"github.com/hawa130/serverx/graph/reqerr"
	"github.com/hawa130/serverx/internal/perm"
	"github.com/rs/xid"
)

// EnforceCtx 判断当前用户是否有对某个对象的某个操作的权限
func EnforceCtx(ctx context.Context, obj string, act string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.Enforce(user.ID.String(), obj, act)
}

// EnforceXCtx 判断当前用户是否有对某个对象的某个操作的权限
func EnforceXCtx(ctx context.Context, obj fmt.Stringer, act string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.EnforceX(user.ID, obj, act)
}

// EnforceReq 判断当前用户是否有对某个对象的某个操作的权限，如果没有则返回 forbidden 错误
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

// EnforceXReq 判断当前用户是否有对某个对象的某个操作的权限，如果没有则返回 forbidden 错误
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

// IsAdmin 判断当前用户是否是管理员
func IsAdmin(ctx context.Context) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.Enforcer().HasRoleForUser(user.ID.String(), "root")
}

// AdminRequired 判断当前用户是否是管理员，如果不是则返回 forbidden 错误
func AdminRequired(ctx context.Context) error {
	allow, err := IsAdmin(ctx)
	if err != nil {
		return err
	}
	if !allow {
		return reqerr.ErrForbidden
	}
	return nil
}

// GrantObjectPermission 为当前用户授予对某个对象的权限
func GrantObjectPermission(ctx context.Context, obj string) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.GrantObjectPermission(user.ID.String(), obj)
}

// GrantObjectPermissionX 为当前用户授予对某个对象的权限
func GrantObjectPermissionX(ctx context.Context, obj fmt.Stringer) (bool, error) {
	user, ok := FromContext(ctx)
	if !ok {
		return false, nil
	}
	return perm.GrantObjectPermissionX(user.ID, obj)
}

// SelfOrAuthenticated 会判断是否有传入的用户 ID，如果没有则从上下文中获取用户 ID。
// 如果指定了 ID 则判断是否有对该用户的权限。
//
// 适用于修改用户自己信息或管理员修改用户信息的场景。
//
// 返回最终的用户 ID，如果没有权限则返回 forbidden 错误。
func SelfOrAuthenticated(ctx context.Context, uid *xid.ID, act string) (*xid.ID, error) {
	if uid == nil {
		var err error
		uid, err = SelfOrSpecified(ctx, uid)
		if err != nil {
			return uid, err
		}
	}
	err := EnforceXReq(ctx, uid, act)
	if err != nil {
		return uid, err
	}
	return uid, nil
}

// SelfOrSpecified 会判断是否有传入的用户 ID，如果没有则从上下文中获取用户 ID，如果有则直接返回
func SelfOrSpecified(ctx context.Context, uid *xid.ID) (*xid.ID, error) {
	if uid == nil {
		u, ok := FromContext(ctx)
		if !ok {
			return nil, reqerr.ErrForbidden
		}
		uid = &u.ID
	}
	return uid, nil
}
