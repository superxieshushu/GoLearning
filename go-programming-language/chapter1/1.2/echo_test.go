package main

import (
	"strings"
	"testing"
)

//做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6
//节讲解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）

func BenchmarkStringPlus(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringPlus([]string{"a", "b", "c"})
	}
}

func BenchmarkStringJoin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		stringJoin([]string{"a", "b", "c"})
	}
}

func stringPlus(s []string) {
	s1, sep := "", ""
	for i := 0; i < len(s); i++ {
		s1 += sep + s[i]
		sep = " "
	}
}

func stringJoin(s []string) {
	strings.Join(s, "")
}
