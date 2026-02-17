package main

import (
	"fmt"
	"math/big"
	"sync"
)

var (
	cache   = make(map[int]*big.Int)
	pending = make(map[int][]chan *big.Int)
	mu      sync.Mutex
)

func main() {
	fmt.Println("  n   fib(n)")
	fmt.Println("------------")
	for i := 0; i <= 100; i++ {
		fmt.Printf("%3d   %s\n", i, fib(i).String())
	}
}

func fib(n int) *big.Int {
	if n < 3 {
		return big.NewInt(int64(n))
	}

	mu.Lock()
	if v, ok := cache[n]; ok {
		mu.Unlock()
		return v
	}
	if waiters, ok := pending[n]; ok {
		ch := make(chan *big.Int, 1)
		pending[n] = append(waiters, ch)
		mu.Unlock()
		return <-ch
	}
	ch := make(chan *big.Int, 1)
	pending[n] = []chan *big.Int{ch}
	mu.Unlock()

	var a, b *big.Int
	var wg sync.WaitGroup
	wg.Add(2)
	go func() { a = fib(n - 1); wg.Done() }()
	go func() { b = fib(n - 2); wg.Done() }()
	wg.Wait()
	result := new(big.Int).Add(a, b)

	mu.Lock()
	cache[n] = result
	for _, c := range pending[n] {
		c <- result
		close(c)
	}
	delete(pending, n)
	mu.Unlock()
	return result
}
