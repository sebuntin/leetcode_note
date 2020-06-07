package leetcode_279

import "testing"

func TestNumSquares(t *testing.T) {
	type testCase struct {
		in  int
		out int
	}
	cases := []testCase{{1, 1}, {2, 2}, {12, 3}, {13, 2}}
	for i := range cases {
		got := numSquares(cases[i].in)
		if got != cases[i].out {
			t.Fatalf("expected result %d, but got %d", cases[i].out, got)
		}
	}
}
