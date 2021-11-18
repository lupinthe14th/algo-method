package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n int
		a []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 10, a: []int{1, 5, 2, 9, 6, 4, 9, 3, 4, 9}}, want: 9},
		{in: in{n: 49, a: []int{5, 1, 4, 3, 7, 8, 3, 2, 2, 4, 3, 2, 1, 9, 8, 7, 2, 3, 1, 9, 8, 4, 6, 7, 1, 8, 7, 6, 8, 6, 4, 3, 6, 9, 6, 9, 9, 2, 5, 4, 2, 8, 4, 9, 2, 5, 3, 2, 6}}, want: 2},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.a)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}

func Example_main() {
	fd, _ := os.Open(filepath.Join("testdata", "input.txt"))
	orgStdin := os.Stdin
	os.Stdin = fd
	defer func() {
		os.Stdin = orgStdin
		fd.Close()
	}()

	main()

	// Output:
	// 9
}
