// Code generated by __generator__/interpreter.go at once

package builtin

import (
	"testing"

	"github.com/ysugimoto/falco/interpreter/context"
	"github.com/ysugimoto/falco/interpreter/value"
	// "github.com/ysugimoto/falco/interpreter/context"
	// "github.com/ysugimoto/falco/interpreter/value"
)

// Fastly built-in function testing implementation of accept.language_lookup
// Arguments may be:
// - STRING, STRING, STRING
// Reference: https://developer.fastly.com/reference/vcl/functions/content-negotiation/accept-language-lookup/
func Test_Accept_language_lookup(t *testing.T) {

	ret, err := Accept_charset_lookup(
		&context.Context{},
		&value.String{Value: "en:de:fr:nl"},
		&value.String{Value: "nl"},
		&value.String{Value: "ja, unknown"},
	)
	if err != nil {
		t.Errorf("Unexpected error: %s", err)
	}
	if ret.Type() != value.StringType {
		t.Errorf("Unexpected type returned, expect=%s, got=%s", value.StringType, ret.Type())
	}
	v := value.Unwrap[*value.String](ret)
	if v.Value != "nl" {
		t.Errorf("Unexpected value returned, expect=nl, got=%s", v.Value)
	}
}
