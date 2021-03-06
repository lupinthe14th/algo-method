package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	out := 1<<31 - 1
	min := func(x, y int) int {
		if x < y {
			return x
		}
		return y
	}
	for i := range a {
		out = min(out, a[i])
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	fmt.Println(solution(n, a))
}
