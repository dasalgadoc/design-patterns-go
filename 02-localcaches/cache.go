package main

type Cache interface {
	Get(key int) (interface{}, error)
}
