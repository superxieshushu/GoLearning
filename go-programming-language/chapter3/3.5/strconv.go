package main

import (
	"fmt"
	"strconv"
)

func Strconv() {
	x := 130
	y := fmt.Sprintf("%d", x)
	fmt.Println(y, strconv.Itoa(x))

	fmt.Println(strconv.FormatInt(int64(x), 2))

	z, _ := strconv.Atoi(y)
	zz, _ := strconv.ParseInt(y, 10, 64)

	fmt.Println(z)
	fmt.Println(zz)
}
