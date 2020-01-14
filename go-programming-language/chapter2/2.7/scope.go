package main

import "fmt"

func f() {

}

var g = "g"

func Scope() {
	f := "f"
	fmt.Println(f)
	fmt.Println(g)
}
