// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/function/shared"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Querystring_sort_Name = "querystring.sort"

var Querystring_sort_ArgumentTypes = []value.Type{value.StringType}

func Querystring_sort_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Querystring_sort_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Querystring_sort_ArgumentTypes[i] {
			return errors.TypeMismatch(Querystring_sort_Name, i+1, Querystring_sort_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of querystring.sort
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-sort/
func Querystring_sort(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Querystring_sort_Validate(args); err != nil {
		return value.Null, err
	}

	u := value.Unwrap[*value.String](args[0])
	query, err := shared.ParseQuery(u.Value)
	if err != nil {
		return value.Null, errors.New(
			Querystring_sort_Name, "Failed to parse urquery: %s, error: %s", u.Value, err.Error(),
		)
	}

	query.Sort(shared.SortAsc)
	return &value.String{Value: query.String()}, nil
}