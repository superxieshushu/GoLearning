package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

//本打印标准输入中多次出现的行，以重复次数开头
func Dup1() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		//line := input.Text()
		//counts[line] = counts[line] + 1
		counts[input.Text()]++
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

//读取标准输入或是使用 os.Open 打开各个具名文件，并操作它们
func Dup2() {
	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
	}
}

//把全部输入数据读到内存中，一次分割为多行，然后处理它们
func Dup3() {
	counts := make(map[string]int)

	for _, file := range os.Args[1:] {
		data, err := ioutil.ReadFile(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}

}

//修改 dup2 ，出现重复的行时打印文件名称
func P14() {
	counts := make(map[string]int)
	fileName := make(map[string]string)
	files := os.Args[1:]

	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2:%v\n", err)
				continue
			}
			fileName = countLinesAndName(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\t%s\n", n, line, fileName[line])
		}
	}
}

func countLinesAndName(f *os.File, counts map[string]int) map[string]string {
	input := bufio.NewScanner(f)
	fileName := make(map[string]string)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			fileName[input.Text()] = f.Name()
		}
	}
	return fileName
}
