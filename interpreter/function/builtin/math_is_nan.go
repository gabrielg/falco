// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Math_is_nan_Name = "math.is_nan"

var Math_is_nan_ArgumentTypes = []value.Type{value.FloatType}

func Math_is_nan_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Math_is_nan_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Math_is_nan_ArgumentTypes[i] {
			return errors.TypeMismatch(Math_is_nan_Name, i+1, Math_is_nan_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of math.is_nan
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/floating-point-classifications/math-is-nan/
func Math_is_nan(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Math_is_nan_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("math.is_nan")
}
