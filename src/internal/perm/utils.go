package perm

import "fmt"

// AddObjectGroup 将操作对象添加到指定组
func AddObjectGroup(obj, group string) (bool, error) {
	return enforcer.AddNamedGroupingPolicy("g2", obj, group)
}

// AddObjectGroupX 将操作对象添加到指定组
func AddObjectGroupX(obj, group fmt.Stringer) (bool, error) {
	return enforcer.AddNamedGroupingPolicy("g2", obj.String(), group.String())
}

// RemoveObjectGroup 将操作对象从指定组移除
func RemoveObjectGroup(obj, group string) (bool, error) {
	return enforcer.RemoveNamedGroupingPolicy("g2", obj, group)
}

// RemoveObjectGroupX 将操作对象从指定组移除
func RemoveObjectGroupX(obj, group fmt.Stringer) (bool, error) {
	return enforcer.RemoveNamedGroupingPolicy("g2", obj.String(), group.String())
}

// RemoveAllObjectGroups 移除操作对象的所有组
func RemoveAllObjectGroups(obj string) (bool, error) {
	return enforcer.RemoveFilteredNamedGroupingPolicy("g2", 0, obj)
}

// RemoveAllObjectGroupsX 移除操作对象的所有组
func RemoveAllObjectGroupsX(obj fmt.Stringer) (bool, error) {
	return enforcer.RemoveFilteredNamedGroupingPolicy("g2", 0, obj.String())
}

// AddSubjectRole 将用户添加到指定角色
func AddSubjectRole(sub, role string) (bool, error) {
	return enforcer.AddNamedGroupingPolicy("g", sub, role)
}

// AddSubjectRoleX 将用户添加到指定角色
func AddSubjectRoleX(sub, role fmt.Stringer) (bool, error) {
	return enforcer.AddNamedGroupingPolicy("g", sub.String(), role.String())
}

// RemoveSubjectRole 将用户从指定角色移除
func RemoveSubjectRole(sub, role string) (bool, error) {
	return enforcer.RemoveNamedGroupingPolicy("g", sub, role)
}

// RemoveSubjectRoleX 将用户从指定角色移除
func RemoveSubjectRoleX(sub, role fmt.Stringer) (bool, error) {
	return enforcer.RemoveNamedGroupingPolicy("g", sub.String(), role.String())
}

// RemoveAllSubjectRoles 移除用户的所有角色
func RemoveAllSubjectRoles(sub string) (bool, error) {
	return enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, sub)
}

// RemoveAllSubjectRolesX 移除用户的所有角色
func RemoveAllSubjectRolesX(sub fmt.Stringer) (bool, error) {
	return enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, sub.String())
}
