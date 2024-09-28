package rule

import (
	"context"

	"github.com/hawa130/serverx/ent/privacy"
	"github.com/hawa130/serverx/internal/auth"
)

// DenyAuthenticated 拒绝未登录用户访问
func DenyAuthenticated() privacy.QueryMutationRule {
	return privacy.ContextQueryMutationRule(func(ctx context.Context) error {
		user, ok := auth.FromContext(ctx)
		if !ok || user == nil {
			return privacy.Denyf("unauthenticated")
		}
		return privacy.Allow
	})
}
