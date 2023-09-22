package tests

import (
	"github.com/codr7/gstraps/utils"
	"testing"
)

func TestUtilsSet(t *testing.T) {
	var s utils.Set

	s.Init(utils.CompareInt)
	s.Add(5)
	s.Add(2)
	s.Add(4)
	s.Add(1)
	s.Add(3)

	s.Remove(3)
	s.Remove(5)
	s.Remove(1)
}
