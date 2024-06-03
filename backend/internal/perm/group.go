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
