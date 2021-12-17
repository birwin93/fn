package fn_test

import (
	"strconv"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestFlatMap(t *testing.T) {
	arr := []int{1, 2, 3}
	mappedArr := fn.FlatMap(arr, func(i int) []string {
		return []string{strconv.Itoa(i), strconv.Itoa(i * 2)}
	})
	require.Equal(t, mappedArr, []string{"1", "2", "2", "4", "3", "6"}, "they should be equal")
}
