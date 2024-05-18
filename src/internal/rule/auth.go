package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/auth"
)

func DenyUnauthorized() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, ok := auth.FromContext(ctx)
		if !ok || user == nil {
			return privacy.Denyf("unauthenticated")
		}
		return privacy.Allow
	})
}
