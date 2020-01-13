package main

import (
	"crypto/sha256"
	"fmt"
)

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func Array() {
	symbol := [...]string{USD: "$", EUR: "€", GBP: "￡", RMB: "￥"}
	fmt.Println(RMB, symbol[RMB])
}

func Sha() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("x"))
	fmt.Printf("%x\n %x\n %t\n %T\n", c1, c2, c1 == c2, c1)
}

func CompareSha256(s1 string, s2 string) {
	a1 := sha256.Sum256([]byte(s1))
	a2 := sha256.Sum256([]byte(s2))
	num := 0
	for i := 0; i < len(a1); i++ {
		for m := 1; m <= 8; m++ {
			if a1[i]>>uint(m) != a2[i]>>uint(m) {
				num++
			}
		}
	}
	fmt.Println(num)
}

func Slice() {
	var y []int
	for i := 0; i < 10; i++ {
		y = append(y, i)
		fmt.Printf("%d cap = %d\t %v\n", i, cap(y), y)
	}
}
