package utils

type Set[T any] struct {
	compare Compare[T]
	items   []T
}

func NewSet[T any](compare Compare[T]) *Set[T] {
	return new(Set[T]).Init(compare)
}

func (self *Set[T]) Init(compare Compare[T]) *Set[T] {
	self.compare = compare
	return self
}

func (self Set[T]) Index(value T) (int, bool) {
	min, max := 0, len(self.items)

	for min < max {
		i := (min + max) / 2

		switch self.compare(value, self.items[i]) {
		case Lt:
			max = i
		case Eq:
			return i, true
		case Gt:
			min = i + 1
		}
	}

	return min, false
}

func (self Set[T]) Exists(value T) bool {
	_, ok := self.Index(value)
	return ok
}

func (self Set[T]) Clone() *Set[T] {
	dst := NewSet[T](self.compare)
	dst.items = make([]T, len(self.items))
	copy(dst.items, self.items)
	return dst
}

func (self *Set[T]) Add(value T) bool {
	if i, exists := self.Index(value); !exists {
		self.items = append(self.items, value)
		copy(self.items[i+1:], self.items[i:])
		self.items[i] = value
		return true
	}

	return false
}

func (self *Set[T]) Remove(value T) bool {
	if i, exists := self.Index(value); exists {
		self.items = self.items[:i+copy(self.items[i:], self.items[i+1:])]
		return true
	}

	return false
}

func (self Set[T]) Items() []T {
	out := make([]T, len(self.items))
	copy(out, self.items)
	return out
}

func (self Set[T]) Len() int {
	return len(self.items)
}
