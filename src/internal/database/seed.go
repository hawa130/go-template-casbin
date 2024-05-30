package database

import (
	"context"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/permission"
	"github.com/hawa130/computility-cloud/ent/role"
	"github.com/hawa130/computility-cloud/internal/constant"
	"github.com/rs/xid"
)

func seedData(ctx context.Context) error {
	seeds := []func(context.Context) error{
		seedPermissions,
		seedRoles,
		seedUser,
	}

	for _, seed := range seeds {
		if err := seed(ctx); err != nil {
			return err
		}
	}

	return nil
}

func seedPermissions(ctx context.Context) error {
	if err := deleteNotExistsPermissions(ctx, constant.AllPermissions); err != nil {
		return err
	}
	for _, p := range constant.AllPermissions {
		if err := upsertPermission(ctx, p); err != nil {
			return err
		}
	}
	return nil
}

func seedRoles(ctx context.Context) error {
	allPermissions, err := client.Permission.Query().All(ctx)
	if err != nil {
		return err
	}

	userReadSummary, err := client.Permission.Query().Where(permission.Name("user:read:summary")).Only(ctx)
	if err != nil {
		return err
	}

	roles := []*ent.Role{
		{
			Name:        "admin",
			Description: "管理员",
			Edges:       ent.RoleEdges{Permissions: allPermissions},
		},
		{
			Name:        "user",
			Description: "用户",
			Edges:       ent.RoleEdges{Permissions: []*ent.Permission{userReadSummary}},
		},
	}

	if err := deleteNotExistsRoles(ctx, roles); err != nil {
		return err
	}

	for _, r := range roles {
		if err := upsertRole(ctx, r); err != nil {
			return err
		}
	}

	return nil
}

func seedUser(ctx context.Context) error {
	exists, err := client.User.Query().Exist(ctx)
	if err != nil {
		return err
	}
	if exists {
		return nil
	}

	adminRole, err := client.Role.Query().Where(role.NameEQ("admin")).Only(ctx)
	if err != nil {
		return err
	}

	if err := client.User.Create().
		SetPhone("12345678910").
		SetPassword("root").
		AddRoleIDs(adminRole.ID).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func upsertPermission(ctx context.Context, p *ent.Permission) error {
	existingPerm, err := client.Permission.Query().Where(permission.NameEQ(p.Name)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	// Create permission if not exists
	if existingPerm == nil || ent.IsNotFound(err) {
		if err := client.Permission.Create().
			SetName(p.Name).
			SetDescription(p.Description).
			Exec(ctx); err != nil {
			return err
		}
	}

	// Update permission if exists
	if err := client.Permission.UpdateOne(existingPerm).
		SetDescription(p.Description).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

// deleteNotExistsPermissions deletes permissions that are not in the given list
func deleteNotExistsPermissions(ctx context.Context, permissions []*ent.Permission) error {
	permissionNames := make([]string, len(permissions))
	for i, p := range permissions {
		permissionNames[i] = p.Name
	}

	if _, err := client.Permission.Delete().
		Where(permission.Not(permission.NameIn(permissionNames...))).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}

func upsertRole(ctx context.Context, r *ent.Role) error {
	existingRole, err := client.Role.Query().Where(role.NameEQ(r.Name)).Only(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return err
	}

	pIDs := make([]xid.ID, len(r.Edges.Permissions))
	for i, p := range r.Edges.Permissions {
		pIDs[i] = p.ID
	}

	// Create role if not exists
	if existingRole == nil || ent.IsNotFound(err) {
		if err := client.Role.Create().
			SetName(r.Name).
			SetDescription(r.Description).
			AddPermissionIDs(pIDs...).
			Exec(ctx); err != nil {
			return err
		}
		return nil
	}

	// Update role permissions if exists
	if err := client.Role.
		UpdateOne(existingRole).
		ClearPermissions().
		AddPermissionIDs(pIDs...).
		Exec(ctx); err != nil {
		return err
	}
	return nil
}

// deleteNotExistsRoles deletes roles that are not in the given list
func deleteNotExistsRoles(ctx context.Context, roles []*ent.Role) error {
	roleNames := make([]string, len(roles))
	for i, r := range roles {
		roleNames[i] = r.Name
	}

	if _, err := client.Role.Delete().
		Where(role.Not(role.NameIn(roleNames...))).
		Exec(ctx); err != nil {
		return err
	}

	return nil
}
