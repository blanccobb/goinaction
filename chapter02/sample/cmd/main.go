package main

import (
	"log"
	"os"
)

// init()는 main()보다 먼저 호출됨.
func init() {
	// change log output to stdout
	log.SetOutput(os.Stdout)
}
