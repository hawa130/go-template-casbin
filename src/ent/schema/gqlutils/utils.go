package gqlutils

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"github.com/vektah/gqlparser/v2/ast"
)

func PermissionDirective(model string) schema.Annotation {
	return entgql.QueryField().Directives(entgql.NewDirective("queryPermission", &ast.Argument{
		Name: "model",
		Value: &ast.Value{
			Raw:  model,
			Kind: ast.StringValue,
		},
	}))
}
