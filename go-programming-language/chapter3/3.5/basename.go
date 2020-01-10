package main

import (
	"strings"
)

func BaseName1(s string) string {

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '/' {
			s = s[i+i:]
			break
		}
	}

	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == '.' {
			s = s[:i]
		}
	}

	return s
}

func BaseName2(s string) string {
	slash := strings.LastIndex(s, "/")
	s = s[slash+1:]
	if dot := strings.LastIndex(s, "."); dot > 0 {
		s = s[:dot]
	}
	return s
}
