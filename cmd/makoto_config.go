package cmd

import (
	"io/ioutil"
	"log"

	"encoding/json"

	kpc "github.com/afeldman/Makoto/kpc"
)

type MakotoConfig struct {
	RootDir []string `json:"makoto_root"`
	Version string   `json:"makoto_version"`
}

var rfg MakotoConfig

func (r *MakotoConfig) init(makotopath []string) {
	r.RootDir = makotopath
	r.Version = MAKOTO_VERSION

	kpc.InitKPCs(r.RootDir)
}

func (r *MakotoConfig) save(path string) {
	d, err := json.Marshal(r)
	if err != nil {
		log.Fatal("cannot jsonize config")
	}

	if err := ioutil.WriteFile(path, d, 0640); err != nil {
		log.Fatal("cannot write configuration into file")
	}
}
