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
