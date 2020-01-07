package main

import (
	"fmt"
	"os"
	"strings"
)

func EchoV1() {
	var s, sep string
	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}

func EchoV2() {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func EchoV3() {
	fmt.Println(strings.Join(os.Args[1:], ""))
}

//练习 1.1： 修改 echo 程序，使其能够打印 os.Args[0] ，即被执行命令本身的名字
func P11() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

//练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行
func P12() {
	for i, arg := range os.Args {
		fmt.Println(i)
		fmt.Println(arg)
	}
}
