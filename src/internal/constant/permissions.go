package constant

import (
	"slices"

	"github.com/hawa130/computility-cloud/ent"
)

// TODO: 限制普通用户查询的字段；数据库同步代码中的角色与权限

// 不要修改常量中的 Name 字段，因为这里的修改会与数据库同步
// Name 是作为权限的唯一标识，修改后会导致原来的权限被删除
// 若要修改，可以关闭服务，在数据库里替换为新名称，再修改这里的常量，最后启动服务
//
// 顺序和描述可以随便改

// 权限命名采用树状结构，例如 user:read:summary 是 user:read 的子权限
// 拥有 user:mutate 权限的用户，自动拥有其子权限
// 即 user:mutate:create, user:mutate:update, user:mutate:delete 权限

// UserPermissions 用户模型的权限
var UserPermissions = []*ent.Permission{
	{Name: "user", Description: "用户"},
	{Name: "user:read", Description: "查询用户详情"},
	{Name: "user:read:summary", Description: "查询用户"},
	{Name: "user:list", Description: "列出用户详情"},
	{Name: "user:list:summary", Description: "列出用户"},
	{Name: "user:mutate", Description: "增改删用户"},
	{Name: "user:mutate:create", Description: "创建用户"},
	{Name: "user:mutate:update", Description: "更新用户"},
	{Name: "user:mutate:delete", Description: "删除用户"},
}

// RolePermissions 角色模型的权限
var RolePermissions = []*ent.Permission{
	{Name: "role", Description: "角色"},
	{Name: "permission", Description: "权限"},
	{Name: "role:list", Description: "查询角色"},
	{Name: "permission:list", Description: "查询权限"},
	{Name: "role:mutate", Description: "增改删角色"},
	{Name: "permission:mutate", Description: "增改删权限"},
}

var AllPermissions = slices.Concat(UserPermissions, RolePermissions)
