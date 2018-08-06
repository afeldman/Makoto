package main

import (
	"log"

	"github.com/afeldman/Makoto/cmd"
)

func main() {
	if err := cmd.Makoto.Execute(); err != nil {
		log.Fatal(err)
	}
}
