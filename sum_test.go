package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestSum(t *testing.T) {
	arr := []int{1, 2, 3}
	sum := fn.Sum(arr)

	require.Equal(t, sum, 6, "they should be equal")
}

func ExampleSum() {
	arr := []int{1, 2, 3}
	fmt.Println(fn.Sum(arr))
	// Output: 6
}
