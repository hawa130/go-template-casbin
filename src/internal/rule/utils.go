package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent/privacy"
)

func WithAllowContext(ctx context.Context) context.Context {
	return privacy.DecisionContext(ctx, privacy.Allow)
}

type queryAllFields struct{}

func WithQueryAllFields(ctx context.Context) context.Context {
	return context.WithValue(ctx, queryAllFields{}, true)
}

func IsQueryAllFields(ctx context.Context) bool {
	if allow, exists := ctx.Value(queryAllFields{}).(bool); exists && allow {
		return true
	}
	return false
}
