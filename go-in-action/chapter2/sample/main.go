package main

import (
	"log"
	"os"

	_ "learning/go-in-action/chapter2/sample/matchers"
	"learning/go-in-action/chapter2/sample/search"
)

// init is called prior to main.
func init() {
	// Change the device for logging to stdout.
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("president")
}
