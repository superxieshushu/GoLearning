package main

import (
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().Unix())
	Lissajous(os.Stdout)
}
