package convert

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestWalk(t *testing.T) {
	Walk("/home/ming/learn/go/pingcap-projects/tidb/util/stmtsummary")
	require.Equal(t, "", "")
}

func TestProcess(t *testing.T) {
	require.Fail(t, "TestProcess")
}

func TestProcessLine(t *testing.T) {
	tests := []struct {
		in    string
		out    string
	}{
		{ "c.Assert(err, IsNil)", "require.Nil(t, err)" },
	}
		for _, tt := range tests {
			got, _ := ProcessLine(tt.in)
			require.Equal(t, tt.out, got)
	}
}