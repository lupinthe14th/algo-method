package main

import (
	"fmt"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Parallel()
	type in struct {
		n, v int
		a    []int
	}
	tests := []struct {
		in   in
		want int
	}{
		{in: in{n: 5, v: 1, a: []int{3, 5, 1, 9, 1}}, want: 2},
	}
	for i, tt := range tests {
		i, tt := i, tt
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			t.Parallel()
			got := solution(tt.in.n, tt.in.v, tt.in.a)
			if got != tt.want {
				t.Fatalf("in: %v got: %v want: %v", tt.in, got, tt.want)
			}
		})
	}
}
