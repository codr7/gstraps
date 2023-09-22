package tests

import (
	"github.com/codr7/gstraps/utils"
	"testing"
)

func TestUtilsSet(t *testing.T) {
	var s utils.Set[int]

	s.Init(utils.CompareInt)
	s.Add(5)
	s.Add(2)
	s.Add(4)
	s.Add(1)
	s.Add(3)

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
