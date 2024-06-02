package perm

func CheckAdmin(args ...interface{}) (interface{}, error) {
	name, ok := args[0].(string)
	if !ok {
		return false, nil
	}
	return enforcer.HasRoleForUser(name, "root")
}
