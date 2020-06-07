package main

import "testing"

func TestfindLongestSubString(t testing.T) {
	type inputCase struct {
		in string
		out int
	}
	testCases := []inputCase{
		{
			"leetcodeisgreat",
			5,
		},
		{
			"leetcedeisgreat",
			7,
		},
	}
	for _, testcase := range testCases {
		if testcase.out != findLongestSubString(testcase.in) {
			t.Fail()
		}
	}
}