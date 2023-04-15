package main

type intParamFunction func(key int) (interface{}, error)

type FunctionResult struct {
	value interface{}
	err   error
}

// Data representantion for cache
type IntCacheMemory struct {
	cacheableFunction intParamFunction
	cache             map[int]FunctionResult
}

func NewCache(f intParamFunction) *IntCacheMemory {
	return &IntCacheMemory{
		cacheableFunction: f,
		cache:             make(map[int]FunctionResult),
	}
}

func (m *IntCacheMemory) Get(key int) (interface{}, error) {
	result, exists := m.cache[key]
	if !exists {
		result.value, result.err = m.cacheableFunction(key)
		m.cache[key] = result
	}
	return result.value, result.err
}
