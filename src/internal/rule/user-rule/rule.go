package userrule

import (
	"context"

	"entgo.io/ent/entql"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/ent/user"
	"github.com/hawa130/computility-cloud/internal/auth"
	"github.com/hawa130/computility-cloud/internal/perm"
	"github.com/hawa130/computility-cloud/internal/rule"
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

// LimitQueryFields 限制查询字段
func LimitQueryFields() privacy.QueryRule {
	return privacy.UserQueryRuleFunc(func(ctx context.Context, q *ent.UserQuery) error {
		// 允许查询所有字段
		if rule.IsQueryAllFields(ctx) {
			return privacy.Skip
		}
		// 允许经过授权的用户查询
		if allow, err := auth.EnforceCtx(ctx, user.Table, perm.OpRead); err == nil && allow {
			return privacy.Skip
		}
		// 限制查询字段
		q.Select(
			user.FieldID,
			user.FieldNickname,
			user.FieldUsername,
			user.FieldCreatedAt,
		)
		return privacy.Skip
	})
}
