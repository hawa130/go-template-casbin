package utils

import "github.com/hawa130/computility-cloud/ent"

func OpToString(op ent.Op) string {
	switch op {
	case ent.OpCreate:
		return "create"
	case ent.OpUpdate:
		return "update"
	case ent.OpUpdateOne:
		return "update"
	case ent.OpDelete:
		return "delete"
	case ent.OpDeleteOne:
		return "delete"
	}
	return ""
}
