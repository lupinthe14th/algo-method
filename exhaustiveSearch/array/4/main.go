package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, v int, a []int) int {
	for i := len(a) - 1; i >= 0; i-- {
		if v == a[i] {
			return i
		}
	}
	return -1
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n, v int
	fmt.Fscan(r, &n, &v)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	fmt.Println(solution(n, v, a))
}
