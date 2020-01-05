package chapter2

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")
var step = flag.String("s", " ", "separator")

func Flag() {
	flag.Parse()
	fmt.Print(strings.Join(flag.Args(), *step))
	if !*n {
		fmt.Println()
	}
}
