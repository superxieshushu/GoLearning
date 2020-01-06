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
func P1() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}

//练习 1.2： 修改 echo 程序，使其打印每个参数的索引和值，每个一行
func P2() {
	for i, arg := range os.Args {
		fmt.Println(i)
		fmt.Println(arg)
	}
}

//做实验测量潜在低效的版本和使用了 strings.Join 的版本的运行时间差异。（1.6
//节讲解了部分 time 包，11.4节展示了如何写标准测试程序，以得到系统性的性能评测。）
//todo
