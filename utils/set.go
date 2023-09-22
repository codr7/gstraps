package utils

type Set[K, V any] struct {
	compare Compare[K, V]
	items   []V
}

func NewSet[K, V any](compare Compare[K, V]) *Set[K, V] {
	return new(Set[K, V]).Init(compare)
}

func (self *Set[K, V]) Init(compare Compare[K, V]) *Set[K, V] {
	self.compare = compare
	return self
}

func (self Set[K, V]) Index(key K) (int, bool) {
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

func (self Set[K, V]) Member(key K) bool {
	_, ok := self.Index(key)
	return ok
}

func (self Set[K, V]) Clone() *Set[K, V] {
	dst := NewSet[K, V](self.compare)
	dst.items = make([]V, len(self.items))
	copy(dst.items, self.items)
	return dst
}

func (self *Set[K, V]) Insert(index int, value V) {
	self.items = append(self.items, value)
	copy(self.items[index+1:], self.items[index:])
	self.items[index] = value
}

func (self *Set[K, V]) Add(key K, value V) bool {
	if i, found := self.Index(key); !found {
		self.Insert(i, value)
		return true
	}

	return false
}

func (self *Set[K, V]) Remove(key K) bool {
	if i, found := self.Index(key); found {
		self.items = self.items[:i+copy(self.items[i:], self.items[i+1:])]
		return true
	}

	return false
}

func (self Set[K, V]) Items() []V {
	out := make([]V, len(self.items))
	copy(out, self.items)
	return out
}

func (self Set[K, V]) Len() int {
	return len(self.items)
}
