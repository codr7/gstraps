package utils

type Set [T]struct {
	compare Compare
	items   []T
}

func NewSet[T any](compare Compare) *Set[T] {
	return new(Set[T]).Init(compare)
}

func (self *Set[T]) Init(compare Compare) *Set[T] {
	self.compare = compare
	return m
}

func (self Set[T]) Index(value T) (int, bool) {
	min, max := 0, len(self.items)

	for min < max {
		i := (min + max) / 2

		switch self.compare(key, self.items[i]) {
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
	_, ok := self.Index(key)
	return ok
}

func (self Set[T]) Clone() *Set[T] {
	dst := NewSet[T](self.compare)
	dst.items = make([]T, len(self.items))
	copy(dst.items, self.items)
	return dst
}

func (self *Set[T]) Add(value T) bool {
	if i, exists := self.Index(key); !exists {
		self.items = append(self.items, value)
		copy(self.items[i+1:], self.items[i:])
		self.items[i] = value
		return true
	}

	return false
}

func (self *Set[T]) Remove(value T) bool {
	if i, exists := self.Index(key); !exists {
		self.items = self.items[:i+copy(self.items[i:], self.items[i+1:])]
		return true
	}

	return false
}

func (self Set[T]) Items() []T {
	out := make([]interface{}, len(self.items))

	for i, it := range self.items {
		out[i] = it.value
	}

	return out
}

func (self Set[T]) Len() int {
	return len(self.items)
}

func (self *Set) AddAll(src Map) {
	AddAll(m, src)
}

func (self *Set) KeepAll(src Map) {
	newLen := len(self.items)
	keep := make([]bool, newLen)

	for i, it := range self.items {
		found := src.Find(it.key) != nil
		keep[i] = found

		if !found {
			newLen--
		}
	}

	newItems := make([]SetItem, newLen)
	i := 0

	for j, it := range self.items {
		if keep[j] {
			newItems[i] = it
			i++
		}
	}

	self.items = newItems
}
