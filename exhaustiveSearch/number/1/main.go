package main

import (
	"bufio"
	"fmt"
	"os"
)

func solution(n int) int {
	out := 0
	for i := 1; i <= n; i++ {
		if i%2 != 0 && i%3 != 0 && i%5 != 0 {
			out++
		}
	}
	return out
}

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	fmt.Println(solution(n))
}
