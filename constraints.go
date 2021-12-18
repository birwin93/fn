package fn

import "constraints"

type Num interface {
	constraints.Integer | constraints.Float | constraints.Complex
}
