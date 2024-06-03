package rule

import (
	"context"
	"fmt"

	"entgo.io/ent/entql"
	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/rule/utils"
)

// LimitQueryFields 限制查询字段
func LimitQueryFields(fields ...string) privacy.QueryRule {
	type EntQuery interface {
		Select(fields ...string) any
	}
	return privacy.QueryRuleFunc(func(ctx context.Context, q ent.Query) error {
		// 允许查询所有字段
		if IsQueryAllFields(ctx) {
			return privacy.Skip
		}
		// 限制查询字段
		query, ok := q.(EntQuery)
		if !ok {
			return fmt.Errorf("unexpected query type %T", q)
		}
		query.Select(fields...)
		return privacy.Skip
	})
}

// FilterQuery 过滤查询
func FilterQuery(fn func(u *ent.User) entql.P) privacy.QueryRule {
	type UserFilter interface {
		Where(p entql.P)
	}
	return privacy.FilterFunc(func(ctx context.Context, f privacy.Filter) error {
		u, err := utils.GetUserOrSkip(ctx)
		if err != nil {
			return err
		}

		filter, ok := f.(UserFilter)
		if !ok {
			return fmt.Errorf("unexpected filter type %T", f)
		}

		filter.Where(fn(u))

		return privacy.Skip
	})
}
