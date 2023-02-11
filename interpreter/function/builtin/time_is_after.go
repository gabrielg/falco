// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Time_is_after_Name = "time.is_after"

var Time_is_after_ArgumentTypes = []value.Type{value.TimeType, value.TimeType}

func Time_is_after_Validate(args []value.Value) error {
	if len(args) != 2 {
		return errors.ArgumentNotEnough(Time_is_after_Name, 2, args)
	}
	for i := range args {
		if args[i].Type() != Time_is_after_ArgumentTypes[i] {
			return errors.TypeMismatch(Time_is_after_Name, i+1, Time_is_after_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of time.is_after
// Arguments may be:
// - TIME, TIME
// Reference: https://developer.fastly.com/reference/vcl/functions/date-and-time/time-is-after/
func Time_is_after(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Time_is_after_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("time.is_after")
}
