package knapsack01

import (
	"testing"
)

func TestKnapsack(t *testing.T) {
	weight := []int{2, 3, 5, 6, 9, 5}
	value := []int{4, 1, 5, 3, 10, 4}
	c := 20
	target := knapSack(weight, value, c)
	actual := knapSack01Dp(weight, value, c)
	if target != actual {
		t.Fatalf("Wrong Result: target result %d, "+
			"but got %d\n", target, actual)
	}
}
