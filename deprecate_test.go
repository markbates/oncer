package oncer

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Deprecate_NoMessage(t *testing.T) {
	r := require.New(t)

	bb := &bytes.Buffer{}
	deprecationWriter = bb

	Deprecate(4, "Test_Deprecate_NoMessage", "")
	act := bb.String()
	r.True(strings.HasPrefix(act, "[DEPRECATED] Test_Deprecate_NoMessage has been deprecated."))
	r.Contains(act, "deprecate_test.go:")
}

func Test_Deprecate_Message(t *testing.T) {
	r := require.New(t)

	bb := &bytes.Buffer{}
	deprecationWriter = bb

	Deprecate(4, "Test_Deprecate_Message", "Use something else instead")
	act := bb.String()
	r.True(strings.HasPrefix(act, "[DEPRECATED] Test_Deprecate_Message has been deprecated."))
	r.Contains(act, "deprecate_test.go:")
	r.Contains(act, "Use something else instead")
}
