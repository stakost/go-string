// Package main base for reverse string
package main

import "unicode/utf8"

// Reverse fast string reverse
func Reverse(s string) string {
	return ReverseFast(s)
}

// ReverseV1 old school method
func ReverseV1(s string) string {
	n := 0
	rune := make([]rune, len(s))
	for _, r := range s {
		rune[n] = r
		n++
	}
	rune = rune[0:n]
	// Reverse
	for i := 0; i < n/2; i++ {
		rune[i], rune[n-1-i] = rune[n-1-i], rune[i]
	}
	// Convert back to UTF-8.
	return string(rune)
}

// ReverseV2 more simple method
func ReverseV2(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

// ReverseFast more fast method
func ReverseFast(s string) string {
	size := len(s)
	buf := make([]byte, size)
	for start := 0; start < size; {
		r, n := utf8.DecodeRuneInString(s[start:])
		start += n
		utf8.EncodeRune(buf[size-start:], r)
	}
	return string(buf)
}

// ReverseClosure experimental and slow method
func ReverseClosure(s string) (result string) {
	for _, v := range s {
		defer func(r rune) { result += string(r) }(v)
	}
	return
}
