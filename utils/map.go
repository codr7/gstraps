package utils

type MapItem[K, V any] struct {
	key   K
	value V
}

type Map[K, V any] struct {
	Set[K, MapItem[K, V]]
}

func NewMap[K, V any](compare Compare[K, K]) *Map[K, V] {
	return new(Map[K, V]).Init(compare)
}

func (self *Map[K, V]) Init(compare Compare[K, K]) *Map[K, V] {
	self.Set.Init(func(l K, r MapItem[K, V]) Order {
		return compare(l, r.key)
	})

	return self
}

func (self *Map[K, V]) Find(key K) (V, bool) {
	if i, ok := self.Index(key); ok {
		return self.items[i].value, true
	}

	var v V
	return v, false
}

func (self *Map[K, V]) Upsert(key K, value V) {
	if i, ok := self.Index(key); ok {
		self.items[i].value = value
	} else {
		self.Insert(i, MapItem[K, V]{key, value})
	}
}
