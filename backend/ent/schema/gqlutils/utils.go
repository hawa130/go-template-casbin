package gqlutils

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"github.com/vektah/gqlparser/v2/ast"
)

func PermissionDirective(object, action string) schema.Annotation {
	return entgql.QueryField().Directives(entgql.NewDirective("permission",
		&ast.Argument{
			Name: "object",
			Value: &ast.Value{
				Raw:  object,
				Kind: ast.StringValue,
			},
		},
		&ast.Argument{
			Name: "action",
			Value: &ast.Value{
				Raw:  action,
				Kind: ast.StringValue,
			},
		},
	))
}
