package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) int {
	out := 0
	for i := range a {
		if a[i] > 0 {
			out++
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
