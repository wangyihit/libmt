package safe_map

import "sync"

type SafeMap struct {
	sync.Map
}

func NewSafeMap() *SafeMap {
	m := &SafeMap{}
	return m
}

func (m *SafeMap) Contains(i interface{}) bool {
	_, ok := m.Load(i)
	return ok
}

func (m *SafeMap) Add(key interface{}, value interface{}) {
	m.Store(key, value)
}

func (m *SafeMap) Keys() []interface{} {
	keys := make([]interface{}, 0)
	m.Range(func(key, _ interface{}) bool {
		keys = append(keys, key)
		return true
	})
	return keys
}
