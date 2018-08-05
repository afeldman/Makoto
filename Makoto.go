package main

import (
	"log"

	"github.com/afeldman/Rossum/src"
)

func main() {
	if err := cmd.Rossum.Execute(); err != nil {
		log.Fatal(err)
	}
}
