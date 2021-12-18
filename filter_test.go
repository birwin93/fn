package fn_test

import (
	"fmt"
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

func ExampleFilter() {
	arr := []int{1, 2, 3}
	filteredArr := fn.Filter(arr, func(i int) bool {
		return i > 1
	})
	fmt.Println(filteredArr)
	// Output: [2 3]
}
