package fn_test

import (
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestReduceToSum(t *testing.T) {
	arr := []int{1, 2, 3}
	sum := fn.Reduce(arr, 0, func(sum, val int) int {
		return sum + val
	})

	require.Equal(t, sum, 6)
}

func TestReduceToMap(t *testing.T) {
	arr := []int{1, 2, 3}
	m := fn.Reduce(arr, map[int]bool{}, func(m map[int]bool, val int) map[int]bool {
		m[val] = true
		return m
	})

	require.Equal(t, m, map[int]bool{1: true, 2: true, 3: true})
}
