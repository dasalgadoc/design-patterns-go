package main

import (
	"fmt"
	"log"
	"os"
	"sync"
	"text/tabwriter"
	"time"
)

func main() {
	// To print tables
	writer := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)

	entrySlice := []int{42, 40, 41, 42, 38}

	cache := NewCache(GetFibonacci)
	fmt.Println("Sequential cache get")
	sequentialGet(entrySlice, cache, writer)

	cacheWithRaceConditions := NewCache(GetFibonacci)
	fmt.Println("Race condition cache get")
	concurrentGet(entrySlice, cacheWithRaceConditions, writer)

	concurrentCache := NewConcurrentCache(GetFibonacci)
	fmt.Println("Concurrent cache get")
	concurrentGet(entrySlice, concurrentCache, writer)

	libraryCacheMemory := NewLibraryCacheMemory(GetFibonacci)
	fmt.Println("Library cache get")
	concurrentGet(entrySlice, libraryCacheMemory, writer)
}

func sequentialGet(
	entrySlice []int, cache Cache, writer *tabwriter.Writer) {

	fmt.Fprintln(writer, "number\ttime\tresult")
	for _, n := range entrySlice {
		calculateFunctionThroughCache(cache, n, writer)
	}

	writer.Flush()
}

func concurrentGet(
	entrySlice []int, cache Cache, writer *tabwriter.Writer) {
	var wg sync.WaitGroup
	fmt.Fprintln(writer, "number\ttime\tresult")

	for _, n := range entrySlice {
		wg.Add(1)
		go func(c Cache, index int, w *tabwriter.Writer) {
			defer wg.Done()
			calculateFunctionThroughCache(c, index, w)
		}(cache, n, writer)
	}

	wg.Wait()
	writer.Flush()
}

func calculateFunctionThroughCache(
	cache Cache, n int, writer *tabwriter.Writer) {
	start := time.Now()
	value, err := cache.Get(n)
	if err != nil {
		log.Println(err)
	}
	line := fmt.Sprintf("%d\t%s\t%d", n, time.Since(start), value)
	fmt.Fprintln(writer, line)
}
