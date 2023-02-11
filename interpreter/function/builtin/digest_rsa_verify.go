// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Digest_rsa_verify_Name = "digest.rsa_verify"

var Digest_rsa_verify_ArgumentTypes = []value.Type{value.IdentType, value.StringType, value.StringType, value.StringType, value.IdentType}

func Digest_rsa_verify_Validate(args []value.Value) error {
	if len(args) < 4 || len(args) > 5 {
		return errors.ArgumentNotInRange(Digest_rsa_verify_Name, 4, 5, args)
	}
	for i := range args {
		if args[i].Type() != Digest_rsa_verify_ArgumentTypes[i] {
			return errors.TypeMismatch(Digest_rsa_verify_Name, i+1, Digest_rsa_verify_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of digest.rsa_verify
// Arguments may be:
// - ID, STRING, STRING, STRING, ID
// - ID, STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-rsa-verify/
func Digest_rsa_verify(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Digest_rsa_verify_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("digest.rsa_verify")
}
