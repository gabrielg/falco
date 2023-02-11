// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Digest_base64url_nopad_decode_Name = "digest.base64url_nopad_decode"

var Digest_base64url_nopad_decode_ArgumentTypes = []value.Type{value.StringType}

func Digest_base64url_nopad_decode_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Digest_base64url_nopad_decode_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Digest_base64url_nopad_decode_ArgumentTypes[i] {
			return errors.TypeMismatch(Digest_base64url_nopad_decode_Name, i+1, Digest_base64url_nopad_decode_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of digest.base64url_nopad_decode
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/cryptographic/digest-base64url-nopad-decode/
func Digest_base64url_nopad_decode(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Digest_base64url_nopad_decode_Validate(args); err != nil {
		return value.Null, err
	}

	// Need to be implemented
	return value.Null, errors.NotImplemented("digest.base64url_nopad_decode")
}
