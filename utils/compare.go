package utils

import (
	"strings"
)

type Order = int

const (
	Lt = Order(-1)
	Eq = Order(0)
	Gt = Order(1)
)

type Compare[T any] = func(l, r T) Order

func CompareInt(l, r int) Order {
	if l < r {
		return Lt
	}

	if l > r {
		return Gt
	}

	return Eq
}

func CompareString(l, r string) Order {
	return Order(strings.Compare(l, r))
}
