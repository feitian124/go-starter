package convert

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func Test_assert(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(stmtEvictedElement.beginTime, Equals, now)",
			&AssertResult{caller: "c", actual: "stmtEvictedElement.beginTime", checker: "Equals", expect: "now"},
		},
		{"  c.Assert(stmtEvictedElement.beginTime(), Equals, now1) ",
			&AssertResult{caller: "c", actual: "stmtEvictedElement.beginTime()", checker: "Equals", expect: "now1"},
		},
	}

	for _, v := range table {
		x, err := assert(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}
