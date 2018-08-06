package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	kpc "github.com/afeldman/Makoto/kpc"
)

type MakotoConfig struct {
	RootDir []string `yaml:"makoto_root"`
	Version string   `yaml:"makoto_version"`
}

var rfg MakotoConfig

func (r *MakotoConfig) init(makotopath []string) {
	r.RootDir = makotopath
	r.Version = MAKOTO_VERSION

	kpc.InitKPCs(r.RootDir)
}

func (r *MakotoConfig) save(path string) {
	d, err := yaml.Marshal(r)
	if err != nil {
		log.Fatal("cannot yamalize config")
	}

	if err := ioutil.WriteFile(path, d, 0640); err != nil {
		log.Fatal("cannot write configuration into file")
	}
}
