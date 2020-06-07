package leetcode_91

import "testing"

func TestNumDecodings(t *testing.T) {
	in := []string{"226", "12"}
	out := []int{3, 2}
	for i:=0; i<len(in); i++ {
		actual := numDecodings(in[i])
		if actual != out[i] {
			t.Fatalf("result wrong, expected %d, but got %d", out[i], actual)
		}
	}
}
