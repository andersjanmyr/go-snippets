package main

import (
	"fmt"
	"math/big"
)

var zero = big.NewInt(0)
var one = big.NewInt(1)
var two = big.NewInt(2)

func main() {
	// fmt.Println(fiboTail(big.NewInt(40)))
	factorial(big.NewInt(100000), func(i *big.Int) { fmt.Println(i) })
}

func factorial(n *big.Int, f func(*big.Int)) {
	if n.Cmp(one) == 0 {
		f(one)
		if one == one {
			panic("help")
		}
	} else {
		tmp := big.NewInt(0)
		factorial(tmp.Sub(n, one), func(y *big.Int) {
			tmp := big.NewInt(0)
			f(tmp.Mul(n, y))
		})
	}
}

// Recursive version of Fibonacci
func fiboR(n *big.Int) *big.Int {
	if n.Cmp(one) <= 0 {
		return one
	}
	t1 := big.NewInt(0).Sub(n, one)
	t2 := big.NewInt(0).Sub(n, two)
	return n.Add(fiboR(t1), fiboR(t2))
}

func fiboTail(n *big.Int) *big.Int {
	return fiboT(n, zero, one)
}

func fiboT(n, first, second *big.Int) *big.Int {
	if n.Cmp(zero) <= 0 {
		return second
	}
	t1 := big.NewInt(0).Sub(n, one)
	sum := big.NewInt(0).Add(first, second)
	return fiboT(t1, second, sum)
}
