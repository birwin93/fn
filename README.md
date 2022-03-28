[![test](https://github.com/birwin93/fn/actions/workflows/go_test.yaml/badge.svg)](https://github.com/birwin93/fn/actions/workflows/go_test.yaml/badge.svg)
[![codecov](https://codecov.io/gh/birwin93/fn/branch/main/graph/badge.svg?token=KNI2VGZT6L)](https://codecov.io/gh/birwin93/fn)

# fn
fn is a collection of go functional operators with generics

## Getting Started

### Prerequisites

Go 1.18

```sh
go install golang.org/dl/go1.181@latest
go1.18 download
```

### Installation
```sh
go get github.com/birwin93/fn
```

## Usage

```go
import "github.com/birwin93/fn"

func main() {
  nums := []int{1, 2, 3}
  
  sum := fn.Sum(nums)
  
  filteredNums := fn.Filter(nums, func(i int) bool {
    return i > 1
  })
  
  doubleNums := fn.Map(nums, func(i int) int {
    return i * 2
  })
  
  if fn.Contains(nums, 2) {
    fmt.Println("nums contains 2!")
  }
}
```

## Supported Functions

### Arrays / Slices

- [Contains](https://github.com/birwin93/fn/blob/main/contains.go)
- [Dedupe](https://github.com/birwin93/fn/blob/main/dedupe.go)
- [Equal](https://github.com/birwin93/fn/blob/main/equal.go))
- [Filter](https://github.com/birwin93/fn/blob/main/filter.go)
- [FlatMap](https://github.com/birwin93/fn/blob/main/flatmap.go)
- [Group](https://github.com/birwin93/fn/blob/main/group.go)
- [Intersect](https://github.com/birwin93/fn/blob/main/intersect.go)
- [Map](https://github.com/birwin93/fn/blob/main/map.go)
- [Last](https://github.com/birwin93/fn/blob/main/last.go)
- [Reduce](https://github.com/birwin93/fn/blob/main/reduce.go)
- [Sum](https://github.com/birwin93/fn/blob/main/sum.go)

### Maps

- [GetOr](https://github.com/birwin93/fn/blob/main/get_or.go)

See the [open issues](https://github.com/birwin93/fn/issues) for a full list of proposed features (and known issues).

## Contributing

If you have a suggestion that would make this better, please fork the repo and create a pull request. You can also simply open an issue with the tag "enhancement". 

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/newfunction`)
3. Commit your Changes (`git commit -m 'Adds fn.NewFunction'`)
4. Push to the Branch (`git push origin feature/newfunction`)
5. Open a Pull Request

## License

Distributed under the MIT License. See `LICENSE` for more information.

## Contact
Twitter - [@billy_the_kid](https://twitter.com/billy_the_kid)
Project Link: [https://github.com/birwin93/fn](https://github.com/birwin93/fn)
