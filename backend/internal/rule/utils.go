package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent/privacy"
)

// WithAllowContext 允许上下文。使用此上下文时 Privacy 隐私层检查会被跳过。
func WithAllowContext(ctx context.Context) context.Context {
	return privacy.DecisionContext(ctx, privacy.Allow)
}
