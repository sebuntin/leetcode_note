package leetcode_78

import (
	"testing"
)
/*
[
  [3],
  [1],
  [2],
  [1,2,3],
  [1,3],
  [2,3],
  [1,2],
  []
]
*/

func TestSubsets(t *testing.T) {
	type test struct {
		input []int
	}

	tests := []test{
		{input:[]int{1,2,3}},
	}
	for _, testCase := range tests {
		got := Subsets(testCase.input)
		t.Logf("input: %v\nresult: %v\n", testCase.input, got)
	}

}
