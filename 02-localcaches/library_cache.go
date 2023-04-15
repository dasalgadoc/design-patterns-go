package main

import (
	"github.com/patrickmn/go-cache"
	"time"
)

type LibraryCacheMemory struct {
	cacheableFunction intParamFunction
	cache             *cache.Cache
}

func NewLibraryCacheMemory(f intParamFunction) *LibraryCacheMemory {
	return &LibraryCacheMemory{
		cacheableFunction: f,
		cache:             cache.New(1*time.Minute, 10*time.Minute),
	}
}

func (m *LibraryCacheMemory) Get(key int) (interface{}, error) {
	var err error
	result, exists := m.cache.Get(string(key))

	if !exists {
		result, err = m.cacheableFunction(key)
		m.cache.Set(string(key), result, cache.DefaultExpiration)
	}

	return result, err
}
