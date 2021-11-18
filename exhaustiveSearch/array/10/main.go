package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	memo := make([]int, 10)
	for i := range a {
		memo[a[i]]++
	}

	max := -1 << 31
	out := 0
	for i := range memo {
		if max < memo[i] {
			max = memo[i]
			out = i
		}
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
