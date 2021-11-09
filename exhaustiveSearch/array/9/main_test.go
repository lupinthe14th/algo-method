package main

import (
	"fmt"
	"os"
	"path/filepath"
	"reflect"
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
		want map[int]int
	}{
		{in: in{n: 10, a: []int{1, 5, 2, 9, 6, 4, 9, 3, 4, 9}}, want: map[int]int{1: 1, 2: 1, 3: 1, 4: 2, 5: 1, 6: 1, 9: 3}},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.a)
			if !reflect.DeepEqual(got, tt.want) {
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
	// 1
	// 1
	// 1
	// 2
	// 1
	// 1
	// 0
	// 0
	// 3
}
