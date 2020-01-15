package main

import (
	"fmt"
	"unicode"
)

func Slice() {
	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i)
		fmt.Printf("%d cap = %d\t %v\n", i, cap(y), y)
	}
}

func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

//重写reverse函数，使用数组指针代替slice。
func P43(s *[6]int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

var rotate = func(s []int, r int) []int {
	lens := len(s)
	arr := make([]int, lens)
	for k := range s {
		index := r + k
		if index >= lens {
			index -= lens
		}
		arr[k] = s[index]
	}
	return arr
}

//编写一个rotate函数，通过一次循环完成旋转。
func P44() {
	a := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(a)

	fmt.Println(rotate(a, 3))
}

var dedupe = func(s []string) []string {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

//写一个函数在原地完成消除[]string中相邻重复的字符串的操作
func P45() {
	s := []string{"a", "a", "a", "b", "c", "c", "b"}
	fmt.Println(dedupe(s))
}

var dedupeSpace = func(s []byte) []byte {
	for i := 0; i < len(s)-1; i++ {
		if unicode.IsSpace(rune(s[i])) && unicode.IsSpace(rune(s[i+1])) {
			copy(s[i:], s[i+1:])
			s = s[:len(s)-1]
			i--
		}
	}
	return s
}

//编写一个函数，原地将一个UTF-8编码的[]byte类型的slice中相邻的空格（参考 unicode.IsSpace）替换成一个空格返回
func P46() {
	s := []byte("a    a   a bc   c  你好     b")
	fmt.Println(string(dedupeSpace(s)))
}

var r = func(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

//修改reverse函数用于原地反转UTF-8编码的[]byte。是否可以不用分配额外的内存
func P47() {
	s := []byte("abc def   ghi jk")
	fmt.Println(string(r(s)))
}
