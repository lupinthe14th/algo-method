package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n, v int, a []int) string {
	for i := range a {
		if a[i] == v {
			return "Yes"
		}
	}
	return "No"
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
