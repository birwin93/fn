package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestGetOr(t *testing.T) {
	dict := map[int]string{
		1: "1",
		2: "2",
	}

	require.Equal(t, fn.GetOr(dict, 1, "3"), "1")
	require.Equal(t, fn.GetOr(dict, 3, "3"), "3")
}

func ExampleGetOr() {
	dict := map[int]string{
		1: "1",
		2: "2",
	}
	fmt.Println(fn.GetOr(dict, 3, "3"))
	// Output: 3
}
