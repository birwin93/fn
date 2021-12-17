package fn_test

import (
	"strconv"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestMap(t *testing.T) {
	arr := []int{1, 2, 3}
	mappedArr := fn.Map(arr, func(i int) string {
		return strconv.Itoa(i)
	})

	require.Equal(t, mappedArr, []string{"1", "2", "3"}, "they should be equal")
}
