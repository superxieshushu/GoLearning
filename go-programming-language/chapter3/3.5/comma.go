package main

import (
	"bytes"
	"strings"
)

func Comma1(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma1(s[:n-3]) + "," + s[n-3:]
}

func Comma2(s string) string {
	var buf bytes.Buffer
	if len(s)%3 != 0 {
		for i := 0; i < len(s)%3; i++ {
			s = "0" + s
		}
	}
	for i := 0; i < len(s); i++ {
		if i != 0 && (i)%3 == 0 {
			buf.WriteByte(',')
		}
		if s[i] != '0' {
			buf.WriteByte(s[i])
		}
	}
	return buf.String()
}

func Comma3(s string) string {
	var buf bytes.Buffer
	for i := 0; i < len(s); i++ {
		buf.WriteByte(s[i])
		if len(s)%3 == (i+1)%3 && len(s) != i+1 {
			buf.WriteByte(',')
		}
	}
	return buf.String()
}

func Comma4(s string) string {
	var buf bytes.Buffer
	if s[0] == '+' || s[0] == '-' {
		s = s[1:]
	}
	split := strings.Split(s, ".")
	for i := 0; i < len(split[0]); i++ {
		buf.WriteByte(split[0][i])
		if len(s)%3 == (i+1)%3 && len(split[0]) != i+1 {
			buf.WriteByte(',')
		}
	}
	if len(split) > 1 {
		buf.WriteByte('.')
		buf.WriteString(split[1])
	}
	return buf.String()
}

func isReverse(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	m := make(map[rune]int)
	n := make(map[rune]int)

	for _, v := range a {
		m[v]++
	}

	for _, v := range b {
		n[v]++
	}

	for k, v := range m {
		if v != n[k] {
			return false
		}
	}
	return true
}
