package utils

type MapItem[K, V any] struct {
	key   K
	value V
}

type Map[K, V any] struct {
	Set[MapItem[K, V]]
}

func NewMap[K, V any](compare Compare[K]) *Map[K, V] {
	return new(Map[K, V]).Init(compare)
}

func (self *Map[K, V]) Init(compare Compare[K]) *Map[K, V] {
	self.Set.Init(func(l, r MapItem[K, V]) Order {
		return compare(l.key, r.key)
	})

	return self
}

func (self *Map[K, V]) Remove(key K) bool {
	var v V
	return self.Set.Remove(MapItem[K, V]{key, v})
}

func (self *Map[K, V]) Upsert(key K, value V) {
	item := MapItem[K, V]{key, value}

	if i, ok := self.Index(item); ok {
		self.items[i] = item
	} else {
		self.Insert(i, item)
	}
}
