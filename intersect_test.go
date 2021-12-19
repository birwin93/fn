package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestIntersect(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 5, 6}
	require.Equal(t, fn.Intersect(arr1, arr2), []int{3, 4})
}

func ExampleIntersect(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	arr2 := []int{3, 4, 5, 6}
	fmt.Println(fn.Intersect(arr1, arr2))
	// Output: [3 4]
}
