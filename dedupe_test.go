package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestDedupe(t *testing.T) {
	arr := []int{1, 2, 3, 4, 1, 3}
	require.Equal(t, fn.Dedupe(arr), []int{1, 2, 3, 4})
}

func TestDedupeStruct(t *testing.T) {
	type person struct {
		name string
		age  int
	}

	billy := person{name: "Billy", age: 20}
	sameBilly := person{name: "Billy", age: 20}
	diffBilly := person{name: "Billy", age: 21}

	billys := []person{billy, sameBilly, diffBilly}
	require.Equal(t, fn.Dedupe(billys), []person{billy, diffBilly})
}

func ExampleDedupe() {
	arr := []int{1, 2, 3, 4, 1, 3}
	fmt.Println(fn.Dedupe(arr))
	// Output: [1 2 3 4]
}
