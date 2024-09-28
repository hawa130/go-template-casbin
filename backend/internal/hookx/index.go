package hookx

import (
	"context"
	"fmt"

	"entgo.io/ent"
	gen "github.com/hawa130/serverx/ent"
	"github.com/rs/xid"
)

type EntMutation interface {
	gen.Mutation
	ID() (xid.ID, bool)
}

type HookFunc func(context.Context, EntMutation) (ent.Value, error)

func (f HookFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	if mv, ok := m.(EntMutation); ok {
		return f(ctx, mv)
	}
	return nil, fmt.Errorf("unexpected mutation type %T. expect EntMutation", m)
}
