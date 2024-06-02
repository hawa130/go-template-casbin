package userrule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/ent/user"
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
			return privacy.Skipf("unexpected error %v", err)
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
