// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"math"
	"math/rand"
	"time"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Randombool_Name = "randombool"

var Randombool_ArgumentTypes = []value.Type{value.IntegerType, value.IntegerType}

func Randombool_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Randombool_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Randombool_ArgumentTypes[i] {
			return errors.TypeMismatch(Randombool_Name, i+1, Randombool_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of randombool
// Arguments may be:
// - INTEGER, INTEGER
// Reference: https://developer.fastly.com/reference/vcl/functions/randomness/randombool/
func Randombool(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Randombool_Validate(args); err != nil {
		return value.Null, err
	}

	numerator := value.Unwrap[*value.Integer](args[0])
	denominator := value.Unwrap[*value.Integer](args[1])

	rand.Seed(time.Now().UnixNano())
	r := rand.Int63n(math.MaxInt64)

	return &value.Boolean{
		Value: r/math.MaxInt64 < numerator.Value/denominator.Value,
	}, nil
}