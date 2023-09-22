package tests

import (
	"github.com/codr7/gstraps/utils"
	"testing"
)

func TestUtilsMap(t *testing.T) {
	var m utils.Map[int, string]
	m.Init(utils.CompareInt)

	m.Upsert(5, "five")
	m.Upsert(2, "two")
	m.Upsert(4, "?")
	m.Upsert(4, "four")
	m.Upsert(1, "one")
	m.Upsert(3, "three")

	if l := m.Len(); l != 5 {
		t.Fatalf("Wrong length: %v", l)
	}

	m.Remove(3)
	m.Remove(5)
	m.Remove(1)

	if l := m.Len(); l != 2 {
		t.Fatalf("Wrong length: %v", l)
	}

	if v, ok := m.Find(2); !ok || v != "two" {
		t.Fatalf("Wrong value: %v", v)
	}

	if v, ok := m.Find(4); !ok || v != "four" {
		t.Fatalf("Wrong value: %v", v)
	}
}

func TestUtilsSet(t *testing.T) {
	var s utils.Set[int, int]
	s.Init(utils.CompareInt)

	s.Add(5, 5)
	s.Add(2, 2)
	s.Add(4, 4)
	s.Add(1, 1)
	s.Add(3, 3)

	if l := s.Len(); l != 5 {
		t.Fatalf("Wrong length: %v", l)
	}

	s.Remove(3)
	s.Remove(5)
	s.Remove(1)

	if l := s.Len(); l != 2 {
		t.Fatalf("Wrong length: %v", l)
	}
}
