package convert

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestEquals(t *testing.T) {
	table := []struct {
		line   string
		result *AssertResult
	}{
		{"c.Assert(stmtEvictedElement.beginTime, Equals, now)",
			&AssertResult{match: true, caller: "c", actual: "stmtEvictedElement.beginTime", checker: "Equals", expect: "now"},
		},
		{"c.Assert(stmtEvictedElement.beginTime, Greater, now)",
			&AssertResult{match: false, caller: "", actual: "", checker: "", expect: ""},
		},
		{"  c.Assert(stmtEvictedElement.beginTime(), Equals, now1) ",
			&AssertResult{match: true, caller: "c", actual: "stmtEvictedElement.beginTime()", checker: "Equals", expect: "now1"},
		},
		{`c.Assert(getAllEvicted(ssbde), Equals, "{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}")`,
			&AssertResult{match: true, caller: "c", actual: "getAllEvicted(ssbde)", checker: "Equals", expect: `"{begin: 5, end: 6, count: 1}, {begin: 1, end: 2, count: 2}"`},
		},
	}

	for _, v := range table {
		x, err := Equals(v.line)
		require.NoError(t, err)
		require.Equal(t, v.result, x)
	}
}

