package perm

import (
	"fmt"

	"github.com/samber/lo"
)

const (
	OpCreate = "create"
	OpRead   = "read"
	OpUpdate = "update"
	OpDelete = "delete"
)

// GrantObjectPermission 获得对象 RUD 权限
func GrantObjectPermission(sub, obj string, act ...string) error {
	_, err := enforcer.AddPolicies(append(
		[][]string{
			{sub, obj, OpRead},
			{sub, obj, OpUpdate},
			{sub, obj, OpDelete},
		},
		lo.Map(act, func(a string, _ int) []string {
			return []string{sub, obj, a}
		})...,
	))
	if err != nil {
		return err
	}

	return nil
}

// GrantObjectPermissionX 获得对象 RUD 权限
func GrantObjectPermissionX(sub, obj fmt.Stringer, act ...string) error {
	return GrantObjectPermission(sub.String(), obj.String(), act...)
}

// RevokeAllPermissionsX 移除用户的所有权限
func RevokeAllPermissionsX(sub fmt.Stringer) (bool, error) {
	return enforcer.RemoveFilteredPolicy(0, sub.String())
}
