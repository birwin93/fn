package fn_test

import (
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3}
	require.True(t, fn.Contains(arr, 1))
	require.False(t, fn.Contains(arr, 4))
}
