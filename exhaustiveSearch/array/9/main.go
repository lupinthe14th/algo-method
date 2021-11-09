package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int, a []int) map[int]int {
	out := make(map[int]int, 9)
	for i := range a {
		out[a[i]]++
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
	out := solution(n, a)
	for i := 1; i <= 9; i++ {
		fmt.Println(out[i])
	}
}
