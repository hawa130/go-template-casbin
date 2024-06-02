package utils

import (
	"context"

	"github.com/hawa130/computility-cloud/ent"
	"github.com/hawa130/computility-cloud/ent/privacy"
	"github.com/hawa130/computility-cloud/internal/auth"
)

// GetUserOrSkip 获取上下文的用户，若无用户则跳过
func GetUserOrSkip(ctx context.Context) (*ent.User, error) {
	user, ok := auth.FromContext(ctx)
	if !ok || user == nil {
		return nil, privacy.Skipf("unauthenticated")
	}
	return user, nil
}

// GetUserOrDeny 获取上下文的用户，若无用户则拒绝
func GetUserOrDeny(ctx context.Context) (*ent.User, error) {
	user, ok := auth.FromContext(ctx)
	if !ok || user == nil {
		return nil, privacy.Denyf("unauthenticated")
	}
	return user, nil
}
