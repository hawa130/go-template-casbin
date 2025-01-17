package utils

import (
	"github.com/hawa130/serverx/ent"
	"github.com/hawa130/serverx/internal/perm"
)

func ToPermOp(op ent.Op) string {
	switch op {
	case ent.OpCreate:
		return perm.OpCreate
	case ent.OpUpdate:
		return perm.OpUpdate
	case ent.OpUpdateOne:
		return perm.OpUpdate
	case ent.OpDelete:
		return perm.OpDelete
	case ent.OpDeleteOne:
		return perm.OpDelete
	}
	return perm.OpRead
}
