// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Std_itoa_Name = "std.itoa"

var Std_itoa_ArgumentTypes = []value.Type{value.IntegerType, value.IntegerType}

func Std_itoa_Validate(args []value.Value) error {
	if len(args) < 1 || len(args) > 2 {
		return errors.ArgumentNotInRange(Std_itoa_Name, 1, 2, args)
	}
	for i := range args {
		if args[i].Type() != Std_itoa_ArgumentTypes[i] {
			return errors.TypeMismatch(Std_itoa_Name, i+1, Std_itoa_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of std.itoa
// Arguments may be:
// - INTEGER, INTEGER
// - INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/std-itoa/
func Std_itoa(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Std_itoa_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("std.itoa")
}
