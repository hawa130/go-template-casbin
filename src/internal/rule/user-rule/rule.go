package userrule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/ent/user"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule/utils"
)

// FilterUserQuery 过滤用户查询
func FilterUserQuery() privacy.QueryRule {
	type UserFilter interface {
		Where(p entql.P)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		u, err := utils.GetUserOrDeny(ctx)
		if err != nil {
			return err
		}

		filter, ok := f.(UserFilter)
		if !ok {
			return privacy.Skipf("unexpected filter type %T", f)
		}

		// 只允许查询自己和自己的子用户（最好可以改成权限实现）
		filter.Where(entql.Or(
			entql.FieldEQ(user.FieldID, u.ID.String()),
			entql.HasEdgeWith(user.EdgeParent, entql.FieldEQ(user.FieldID, u.ID.String())),
		))

		return privacy.Skip
	})
}

// AllowAuthorizedMutation 允许拥有权限的用户修改
func AllowAuthorizedMutation() privacy.MutationRule {
	return privacy.UserMutationRuleFunc(func(ctx context.Context, m *ent.UserMutation) error {
		u, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return err
		}
		id, exists := m.ID()
		if !exists {
			return privacy.Skip
		}

		var allow bool
		if m.Op() == ent.OpCreate {
			allow, err = perm.Enforce(u.ID.String(), user.Table, "create")
		} else {
			allow, err = perm.EnforceX(u.ID, id, utils.ToPermOp(m.Op()))
		}
		if err != nil {
			return privacy.Skipf("unexpected error %v", err)
		}
		if allow {
			return privacy.Allow
		}

		return privacy.Skip
	})
}
