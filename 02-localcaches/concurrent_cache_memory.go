package main

import "sync"

type ConcurrentIntCacheMemory struct {
	cacheableFunction intParamFunction
	cache             map[int]FunctionResult
	lock              sync.RWMutex
}

func NewConcurrentCache(f intParamFunction) *ConcurrentIntCacheMemory {
	return &ConcurrentIntCacheMemory{
		cacheableFunction: f,
		cache:             make(map[int]FunctionResult),
	}
}

func (m *ConcurrentIntCacheMemory) Get(key int) (interface{}, error) {
	m.lock.RLock()
	result, exists := m.cache[key]
	m.lock.RUnlock()

	if !exists {
		m.lock.Lock()
		result.value, result.err = m.cacheableFunction(key)
		m.cache[key] = result
		m.lock.Unlock()
	}
	return result.value, result.err
}
