// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Math_round_Name = "math.round"

var Math_round_ArgumentTypes = []value.Type{value.FloatType}

func Math_round_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Math_round_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Math_round_ArgumentTypes[i] {
			return errors.TypeMismatch(Math_round_Name, i+1, Math_round_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of math.round
// Arguments may be:
// - FLOAT
// Reference: https://developer.fastly.com/reference/vcl/functions/math-rounding/math-round/
func Math_round(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Math_round_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("math.round")
}
