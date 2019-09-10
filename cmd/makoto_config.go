package cmd

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path"

	kpc "github.com/afeldman/kpc"
)

type MakotoConfig struct {
	RootDir string `json:"kpc_root"`
	Version string `json:"makoto_version"`
}

var rfg MakotoConfig

func (r *MakotoConfig) init(makotopath string) {
	r.RootDir = makotopath
	r.Version = MakotoVersion

	kpc.InitKPCs(r.RootDir)
}

func (r *MakotoConfig) save(filepath, filename string) {
	d, err := json.Marshal(r)
	if err != nil {
		log.Fatal("cannot jsonize config")
	}

	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		os.Mkdir(filepath, os.ModePerm)
	}

	if err := ioutil.WriteFile(path.Join(filepath, filename), d, 0640); err != nil {
		log.Fatal("cannot write configuration into file")
	}
}
