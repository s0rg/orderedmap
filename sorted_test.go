package orderedmap_test

import (
	"slices"
	"testing"

	"github.com/s0rg/orderedmap"
)

func TestSortedSet(t *testing.T) {
	t.Parallel()

	var m orderedmap.Sorted[string, string]

	if m.Len() > 0 {
		t.Fail()
	}

	m.Set("1", "foo")

	if m.Len() != 1 {
		t.Fail()
	}

	m.Set("3", "baz")

	if m.Len() != 2 {
		t.Fail()
	}

	m.Set("2", "bar")

	if m.Len() != 3 {
		t.Fail()
	}

	{
		want := []string{"1", "2", "3"}
		have := make([]string, 0, 3)

		for k := range m.Iter {
			have = append(have, k)
		}

		if !slices.Equal(want, have) {
			t.Fail()
		}
	}

	m.Del("1")

	{
		want := []string{"2", "3"}
		have := make([]string, 0, 2)

		for k := range m.Iter {
			have = append(have, k)
		}

		if !slices.Equal(want, have) {
			t.Fail()
		}
	}

	m.Del("3")

	{
		want := []string{"2"}
		have := make([]string, 0, 1)

		for k := range m.Iter {
			have = append(have, k)
		}

		if !slices.Equal(want, have) {
			t.Fail()
		}
	}

	m.Del("2")

	have := make([]string, 0, 1)

	for k := range m.Iter {
		have = append(have, k)
	}

	if len(have) > 0 {
		t.Fail()
	}
}

func TestSortedGet(t *testing.T) {
	t.Parallel()

	var m orderedmap.Sorted[string, string]

	if m.Has("1") {
		t.Fail()
	}

	m.Set("1", "foo")

	val, ok := m.Get("1")
	if !ok {
		t.Fail()
	}

	if val != "foo" {
		t.Fail()
	}

	if m.Has("2") {
		t.Fail()
	}
}

func TestSortedDel(t *testing.T) {
	t.Parallel()

	var m orderedmap.Sorted[string, string]

	m.Del("1")
	m.Del("2")

	m.Set("1", "foo")
	m.Set("2", "bar")

	if !m.Has("1") {
		t.Fail()
	}

	m.Del("1")

	if m.Has("1") {
		t.Fail()
	}

	m.Del("1") // double-delete to make sure its ok

	if m.Len() != 1 {
		t.Log("1", m.Len())
		t.Fail()
	}

	val, ok := m.Get("2")
	if !ok {
		t.Log("2")
		t.Fail()
	}

	if val != "bar" {

		t.Log("3")
		t.Fail()
	}
}

func TestSortedIter(t *testing.T) {
	t.Parallel()

	var m orderedmap.Sorted[string, string]

	for k := range m.Iter {
		_ = k
	}

	m.Set("1", "1")
	m.Set("2", "2")
	m.Set("3", "3")

	for k := range m.Iter {
		if k == "2" {
			m.Del(k)

			break
		}
	}

	if m.Len() != 2 {
		t.Fail()
	}
}

func TestSortedClone(t *testing.T) {
	t.Parallel()

	var m orderedmap.Sorted[string, string]

	c := m.Clone()

	if c.Len() > 0 {
		t.Fail()
	}

	m.Set("1", "1")
	m.Set("2", "2")
	m.Set("3", "3")

	c = m.Clone()

	if c.Len() != 3 {
		t.Fail()
	}

	if !c.Has("2") {
		t.Fail()
	}
}

func TestSortedNew(t *testing.T) {
	t.Parallel()

	m := orderedmap.NewSorted[int, string]()

	m.Set(10, "10")
	m.Set(1, "1")
	m.Set(5, "5")

	want := []int{1, 5, 10}
	have := make([]int, 0, 3)

	for k := range m.Iter {
		have = append(have, k)
	}

	if !slices.Equal(want, have) {
		t.Fail()
	}
}
