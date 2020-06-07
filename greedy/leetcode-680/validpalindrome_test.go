package main

import "testing"

func TestvalidPalindrome(t testing.T) {
	in := "cabc"
	target := true
	out := validPalindrome(in)
	if target != out {
		t.Fail()
	}
}