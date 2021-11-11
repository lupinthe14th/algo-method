package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	m := -1 >> 31
	memo := make(map[int]int, n)
	max := func(x, y int) int {
		if x > y {
			return x
		}
		return y
	}
	for i := range a {
		m = max(m, a[i])
		memo[a[i]] = i
	}
	return memo[m]
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
