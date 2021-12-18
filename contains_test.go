package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestContains(t *testing.T) {
	arr := []int{1, 2, 3}
	require.True(t, fn.Contains(arr, 1))
	require.False(t, fn.Contains(arr, 4))
}

func ExampleContains() {
	fmt.Println(fn.Contains([]int{1, 2, 3}, 1))
	// Output: true
}
