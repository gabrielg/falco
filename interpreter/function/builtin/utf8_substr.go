// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Utf8_substr_Name = "utf8.substr"

var Utf8_substr_ArgumentTypes = []value.Type{value.StringType, value.IntegerType, value.IntegerType}

func Utf8_substr_Validate(args []value.Value) error {
	if len(args) < 2 || len(args) > 3 {
		return errors.ArgumentNotInRange(Utf8_substr_Name, 2, 3, args)
	}
	for i := range args {
		if args[i].Type() != Utf8_substr_ArgumentTypes[i] {
			return errors.TypeMismatch(Utf8_substr_Name, i+1, Utf8_substr_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of utf8.substr
// Arguments may be:
// - STRING, INTEGER, INTEGER
// - STRING, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/utf8-substr/
func Utf8_substr(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Utf8_substr_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("utf8.substr")
}
