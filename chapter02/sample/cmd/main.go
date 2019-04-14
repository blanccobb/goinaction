package main

import (
	"log"
	"os"

	_ "github.com/blanccobb/goinaction/chapter02/sample/matchers"
	"github.com/blanccobb/goinaction/chapter02/sample/search"
)

// init()는 main()보다 먼저 호출됨.
func init() {
	// change log output to stdout
	log.SetOutput(os.Stdout)
}

func main() {
	search.Run("Sherlock Holmes")
}
