package main

import (
	"log"
	"os"

	"github.com/xplaceholder/xplaceholder/config"

	"github.com/xplaceholder/xplaceholder/application"
)

func main() {

	log.SetFlags(0)

	logger := application.NewLogger(os.Stdout, os.Stdin)
	_ = application.NewLogger(os.Stderr, os.Stdin)

	globals, _, err := config.ParseArgs(os.Args)
	if err != nil {
		log.Fatalf("\n\n%s\n", err)
	}
	if globals.NoConfirm {
		logger.NoConfirm()
	}
}
