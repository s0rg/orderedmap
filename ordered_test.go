package orderedmap_test

import (
	"slices"
	"testing"

	"github.com/s0rg/orderedmap"
)

func TestOrderedSet(t *testing.T) {
	t.Parallel()

	var m orderedmap.Ordered[string, string]

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
		want := []string{"1", "3", "2"}
		have := make([]string, 0, 3)

		for k := range m.Iter {
			have = append(have, k)
		}

		if !slices.Equal(want, have) {
			t.Fail()
		}
	}
}

func TestOrderedGet(t *testing.T) {
	t.Parallel()

	var m orderedmap.Ordered[string, string]

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

func TestOrderedDel(t *testing.T) {
	t.Parallel()

	var m orderedmap.Ordered[string, string]

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
		t.Fail()
	}

	val, ok := m.Get("2")
	if !ok {
		t.Fail()
	}

	if val != "bar" {
		t.Fail()
	}
}

func TestOrderedIter(t *testing.T) {
	t.Parallel()

	var m orderedmap.Ordered[string, string]

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

func TestOrderedClone(t *testing.T) {
	t.Parallel()

	var m orderedmap.Ordered[string, string]

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
