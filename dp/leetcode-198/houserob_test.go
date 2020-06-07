package leetcode_198

import "testing"

func TestRobHelper(t *testing.T) {
	inputs := [][]int{
		{2, 3, 4, 5, 6},
		{4, 3, 2, 1},
		{6, 7, 4, 2, 5},
	}
	for _, input := range inputs {
		target := rob(input)
		actual := robHelper(input)
		if actual != target {
			t.Fatalf("Result wrong: input %v, target %v ,but got %v\n", input, target, actual)
		}
	}
}
