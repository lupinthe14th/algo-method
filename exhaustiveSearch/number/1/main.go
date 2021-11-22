package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		if i%2 == 0 {
			continue
		}
		if i%3 == 0 {
			continue
		}
		if i%5 == 0 {
			continue
		}
		out++
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	fmt.Println(solution(n))
}
