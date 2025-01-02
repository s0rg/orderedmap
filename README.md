# orderedmap

Package orderedmap implements two types of generic maps, both iterates keys in ordered manner

# features

- two types of maps:
  - Ordered - iterates in same order as items inserted
  - Sorted - iterates in sorted order
- zero-dependency
- generic
- 100% test coverage

# example

```go
import (
    "fmt"

    "github.com/s0rg/orderedmap"
)

func exampleSorted() {
	var m orderedmap.Sorted[string, string]

	m.Set("1", "foo")
	m.Set("3", "baz")
	m.Set("2", "bar")

	for k, v := range m.Iter {
		fmt.Printf("%s: %s\n", k, v) // will print in order: 1, 2, 3 (as its sorted)
	}
}

func exampleOrdered() {
	var m orderedmap.Ordered[string, string]

	m.Set("1", "foo")
	m.Set("3", "baz")
	m.Set("2", "bar")

	for k, v := range m.Iter {
		fmt.Printf("%s: %s\n", k, v) // will print in order: 1, 3, 2 (in order of insertion)
	}
}
```
