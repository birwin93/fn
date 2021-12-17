package fn_test

import (
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestFilter(t *testing.T) {
	arr := []int{1, 2, 3}
	newArr := fn.Filter(arr, func(i int) bool {
		return i > 1
	})
	require.Equal(t, newArr, []int{2, 3}, "they should be equal")
}
