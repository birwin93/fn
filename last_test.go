package fn_test

import (
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestLast(t *testing.T) {
	arr := []int{1, 2, 3}
	val, ok := fn.Last(arr)
	require.Equal(t, val, 3)
	require.True(t, ok)
}

func TestLastEmpty(t *testing.T) {
	arr := []int{}
	val, ok := fn.Last(arr)
	require.False(t, ok)
	require.Equal(t, val, 0)
}
