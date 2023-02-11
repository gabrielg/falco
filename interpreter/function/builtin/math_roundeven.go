// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Math_roundeven_Name = "math.roundeven"

var Math_roundeven_ArgumentTypes = []value.Type{value.FloatType}

func Math_roundeven_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Math_roundeven_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Math_roundeven_ArgumentTypes[i] {
			return errors.TypeMismatch(Math_roundeven_Name, i+1, Math_roundeven_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of math.roundeven
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-roundeven/
func Math_roundeven(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Math_roundeven_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("math.roundeven")
}
