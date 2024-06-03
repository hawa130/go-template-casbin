package perm

import "fmt"

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

// RemoveAllSubjectRolesX 移除用户的所有角色
func RemoveAllSubjectRolesX(sub fmt.Stringer) (bool, error) {
	return enforcer.RemoveFilteredNamedGroupingPolicy("g", 0, sub.String())
}

// NewModelAdminRole 创建模型管理员角色
func NewModelAdminRole(model string) (bool, error) {
	role := fmt.Sprintf("%s:admin", model)
	return enforcer.AddPoliciesEx([][]string{
		{role, model, OpCreate},
		{role, model, OpRead},
		{role, model, OpUpdate},
		{role, model, OpDelete},
	})
}

// AddModelAdminRole 将用户添加到模型管理员角色
func AddModelAdminRole(sub, model string) (bool, error) {
	role := fmt.Sprintf("%s:admin", model)

	// 检查模型管理员角色是否存在
	exists, err := enforcer.GetFilteredPolicy(0, role)
	if err != nil {
		return false, err
	}

	// 如果不存在模型管理员角色，则创建
	if len(exists) < 4 {
		_, err = NewModelAdminRole(model)
		if err != nil {
			return false, err
		}
	}

	return enforcer.AddNamedGroupingPolicy("g", sub, role)
}

// RemoveModelAdminRole 将用户从模型管理员角色移除
func RemoveModelAdminRole(sub, model string) (bool, error) {
	role := fmt.Sprintf("%s:admin", model)
	return enforcer.RemoveNamedGroupingPolicy("g", sub, role)
}
