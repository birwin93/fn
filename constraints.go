package fn

import "golang.org/x/exp/constraints"

type Num interface {
	constraints.Integer | constraints.Float | constraints.Complex
}
