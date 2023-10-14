// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/function/errors"
	"github.com/ysugimoto/falco/interpreter/value"
)

const Xml_escape_Name = "xml_escape"

var Xml_escape_ArgumentTypes = []value.Type{value.StringType}

func Xml_escape_Validate(args []value.Value) error {
	if len(args) != 1 {
		return errors.ArgumentNotEnough(Xml_escape_Name, 1, args)
	}
	for i := range args {
		if args[i].Type() != Xml_escape_ArgumentTypes[i] {
			return errors.TypeMismatch(Xml_escape_Name, i+1, Xml_escape_ArgumentTypes[i], args[i].Type())
		}
	}
	return nil
}

// Fastly built-in function implementation of xml_escape
// Arguments may be:
// - STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/strings/xml-escape/
func Xml_escape(ctx *context.Context, args ...value.Value) (value.Value, error) {
	// Argument validations
	if err := Xml_escape_Validate(args); err != nil {
		return value.Null, err
	}

	input := value.Unwrap[*value.String](args[0]).Value

	var escaped []byte
	for _, b := range []byte(input) {
		switch b {
		case 0x26: // "&"
			escaped = append(escaped, []byte("&amp;")...)
		case 0x3C: // "<"
			escaped = append(escaped, []byte("&lt;")...)
		case 0x3E: // ">"
			escaped = append(escaped, []byte("&gt;")...)
		case 0x27: // "'"
			escaped = append(escaped, []byte("&apos;")...)
		case 0x22: // '"'
			escaped = append(escaped, []byte("&quot;")...)
		default:
			escaped = append(escaped, b)
		}
	}

	return &value.String{Value: string(escaped)}, nil
}