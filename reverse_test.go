package main

import (
	"testing"
)

var (
	testString = "Hello Go! 🔥🌐 🔥 こんにちは！你好！"
	reverseString = "！好你！はちにんこ 🔥 🌐🔥 !oG olleH"
)

func TestReverseStr(t *testing.T) {
	s := Reverse(testString)
	assertEqualString(t, s, reverseString)

	s = ReverseV1(testString)
	assertEqualString(t, s, reverseString)

	s = ReverseV2(testString)
	assertEqualString(t, s,reverseString)

	s = ReverseFast(testString)
	assertEqualString(t, s,reverseString)

	s = ReverseClosure(testString)
	assertEqualString(t, s,reverseString)
}

func assertEqualString(t *testing.T, a, b string) {
	if a != b {
		t.Errorf("Strings not equal, expected `%s` and got `%s`", a, b)
	}
}
