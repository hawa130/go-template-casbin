package rule

import (
	"context"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/auth"
)

// AllowMutateSelf allows a user to mutate their own user entity.
//
// Effectively, this rule will reset the roles field from the mutation.
func AllowMutateSelf() privacy.MutationRule {
	return privacy.UserMutationRuleFunc(func(ctx context.Context, m *ent.UserMutation) error {
		if !m.Op().Is(ent.OpUpdateOne) {
			return privacy.Skip
		}

		id, exists := m.ID()
		if !exists {
			return privacy.Skipf("missing user id in mutation")
		}

		user, ok := auth.FromContext(ctx)
		if !ok || user == nil {
			return privacy.Skipf("unauthenticated")
		}
		if user.ID == id {
			// Roles should not be modified by normal users
			m.ResetRoles()
			return privacy.Allow
		}

		return privacy.Skip
	})
}
