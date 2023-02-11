// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Querystring_remove_Name = "querystring.remove"

var Querystring_remove_ArgumentTypes = []value.Type{value.StringType}

func Querystring_remove_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Querystring_remove_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Querystring_remove_ArgumentTypes[i] {
			return errors.TypeMismatch(Querystring_remove_Name, i+1, Querystring_remove_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of querystring.remove
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/query-string/querystring-remove/
func Querystring_remove(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Querystring_remove_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("querystring.remove")
}
