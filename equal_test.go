package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestEqual(t *testing.T) {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	c := []int{1, 2, 4}
	d := []int{1, 2, 3, 4}

	require.True(t, fn.Equal(a, b))
	require.False(t, fn.Equal(a, c))
	require.False(t, fn.Equal(a, d))
}

func ExampleEqual() {
	a := []int{1, 2, 3}
	b := []int{1, 2, 3}
	fmt.Println(fn.Equal(a, b))
	// Output: true
}
