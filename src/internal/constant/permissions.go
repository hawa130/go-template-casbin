package constant

import (
	"slices"

	"github.com/hawa130/computility-cloud/ent"
)

var UserPermissions = []*ent.Permission{
	{Name: "user:read:summary", Description: "查询用户"},
	{Name: "user:read:detail", Description: "查询用户详情"},
	{Name: "user:list", Description: "列出用户"},
	{Name: "user:mutate", Description: "增改删用户"},
}

var RolePermissions = []*ent.Permission{
	{Name: "role:list", Description: "查询角色"},
	{Name: "permission:list", Description: "查询权限"},
}

var AllPermissions = slices.Concat(UserPermissions, RolePermissions)
