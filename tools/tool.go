package tools

import (
	"reflect"

	"github.com/99designs/gqlgen/graphql"
	"github.com/mitchellh/mapstructure"
	"github.com/sirupsen/logrus"
	"github.com/vektah/gqlparser/v2/ast"
	"golang.org/x/net/context"
)

func ArgumentToValue(arg *ast.Argument, v interface{}) bool {
	r := reflect.TypeOf(v).Elem()

	if arg.Value.Definition.Name != r.Name() {
		return false
	}

	m, err := arg.Value.Value(map[string]interface{}{})

	if err != nil {
		return false
	}

	if mapstructure.Decode(m, v) != nil {
		return false
	}

	return true
}

func GetValuesFromSelections(ctx context.Context, template interface{}) map[string]interface{} {
	r := reflect.TypeOf(template).Elem()
	gctx := graphql.GetFieldContext(ctx)

	values := make(map[string]interface{})

	for _, selection := range gctx.Field.Selections {
		field := selection.(*ast.Field)
		for _, argument := range field.Arguments {
			v := reflect.New(r).Interface()
			if ArgumentToValue(argument, v) {
				values[field.Alias] = v
			}
		}
	}
	return values
}

func FailOnError(err error, msg string) {
	if err != nil {
		logrus.Panicf("%s: %s", msg, err)
	}
}
