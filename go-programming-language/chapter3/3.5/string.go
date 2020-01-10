package main

import (
	"bytes"
	"fmt"
	"unicode/utf8"
)

func String() {
	var s = "Hello, 世界"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))

	for i := 0; i < len(s); {
		r, size := utf8.DecodeRuneInString(s[i:])
		fmt.Printf("%d\t%c\n", i, r)
		i += size
	}
}

func Rune() {
	s := "プログラム"
	fmt.Printf("%x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)

	//进行 UTF-8 编码
	fmt.Println(string(r))
	fmt.Println(string(65))
	fmt.Printf(string(0x4eac))
}

func byte2String() {
	s := "世界"
	b := []byte(s)
	for _, x := range b {
		fmt.Printf("%x ", x)
	}
	fmt.Println(string(b))
}

func ints2String(value []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range value {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}
