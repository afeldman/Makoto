package main

import (
	"github.com/afeldman/Makoto/cmd"

	log "github.com/sirupsen/logrus"
)

func main() {
	if err := cmd.Makoto.Execute(); err != nil {
		log.Fatalln(err)
	}

}
