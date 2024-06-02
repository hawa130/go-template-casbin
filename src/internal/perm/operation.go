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

// RevokeObjectPermission 撤销对象 RUD 权限
func RevokeObjectPermission(sub, obj string, act ...string) error {
	_, err := enforcer.RemovePolicies(append(
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

// RevokeObjectPermissionX 撤销对象 RUD 权限
func RevokeObjectPermissionX(sub, obj fmt.Stringer, act ...string) error {
	return RevokeObjectPermission(sub.String(), obj.String(), act...)
}
