[![PkgGoDev](https://pkg.go.dev/badge/github.com/s0rg/orderedmap)](https://pkg.go.dev/github.com/s0rg/orderedmap)
[![License](https://img.shields.io/github/license/s0rg/orderedmap)](https://github.com/s0rg/orderedmap/blob/master/LICENSE)
[![Go Version](https://img.shields.io/github/go-mod/go-version/s0rg/orderedmap)](go.mod)
[![Tag](https://img.shields.io/github/v/tag/s0rg/orderedmap?sort=semver)](https://github.com/s0rg/orderedmap/tags)

[![CI](https://github.com/s0rg/grid/workflows/ci/badge.svg)](https://github.com/s0rg/grid/actions?query=workflow%3Aci)
[![Go Report Card](https://goreportcard.com/badge/github.com/s0rg/orderedmap)](https://goreportcard.com/report/github.com/s0rg/orderedmap)
[![Maintainability](https://api.codeclimate.com/v1/badges/87648afa2a219e5ca0c8/maintainability)](https://codeclimate.com/github/s0rg/orderedmap/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/87648afa2a219e5ca0c8/test_coverage)](https://codeclimate.com/github/s0rg/orderedmap/test_coverage)
![Issues](https://img.shields.io/github/issues/s0rg/orderedmap)

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
