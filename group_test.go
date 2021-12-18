package fn_test

import (
	"fmt"
	"testing"

	"github.com/birwin93/fn"
	"github.com/stretchr/testify/require"
)

func TestGroup(t *testing.T) {
	type member struct {
		name string
		age  int
	}

	billy := member{name: "Billy", age: 1}
	bob := member{name: "Bob", age: 1}
	chris := member{name: "Chris", age: 2}
	dennis := member{name: "Dennis", age: 3}
	members := []member{billy, bob, chris, dennis}

	grouped := fn.Group(members, func(m member) rune {
		return rune(m.name[0])
	})

	require.Equal(t, grouped['B'], []member{billy, bob})
	require.Equal(t, grouped['C'], []member{chris})
	require.Equal(t, grouped['D'], []member{dennis})
}

func ExampleGroup() {
	type member struct {
		name string
		age  int
	}

	billy := member{name: "Billy", age: 1}
	bob := member{name: "Bob", age: 1}
	chris := member{name: "Chris", age: 2}
	members := []member{billy, bob, chris}

	groupedMembers := fn.Group(members, func(m member) string {
		return string(m.name[0])
	})
	fmt.Println(groupedMembers)
	// Output: map[B:[{Billy 1} {Bob 1}] C:[{Chris 2}]]
}
