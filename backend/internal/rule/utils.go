package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent/privacy"
)

func WithAllowContext(ctx context.Context) context.Context {
	return privacy.DecisionContext(ctx, privacy.Allow)
}
