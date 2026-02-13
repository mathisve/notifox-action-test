package main

import (
	"fmt"
	"math/big"
	"testing"
)

// fibIter returns the nth Fibonacci number (iterative, for test oracle).
func fibIter(n int) *big.Int {
	if n < 2 {
		return big.NewInt(int64(n))
	}
	a := big.NewInt(0)
	b := big.NewInt(1)
	for i := 2; i <= n; i++ {
		a, b = b, new(big.Int).Add(a, b)
	}
	return b
}

func TestFib(t *testing.T) {
	for n := 0; n <= 50; n++ {
		n := n
		want := fibIter(n)
		t.Run(fmt.Sprintf("fib(%d)", n), func(t *testing.T) {
			got := fib(n)
			if got.Cmp(want) != 0 {
				t.Errorf("fib(%d) = %s, want %s", n, got.String(), want.String())
			}
		})
	}
}
