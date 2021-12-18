package fn_test

import (
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestLast(t *testing.T) {
	arr := []int{1, 2, 3}
	val, err := fn.Last(arr)
	require.Equal(t, *val, 3)
	require.Nil(t, err)
}

func TestLastError(t *testing.T) {
	arr := []int{}
	val, err := fn.Last(arr)
	require.Nil(t, val)
	require.Equal(t, err, &fn.EmptyArrayError{})
}
