package cmd

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"

	kpc "github.com/afeldman/Makoto"
)

type RossumConfig struct {
	RootDir []string `yaml:"rossum_root"`
	Version string   `yaml:"rossum_version"`
}

var rfg RossumConfig

func (r *RossumConfig) init(rossumpath []string) {
	r.RootDir = rossumpath
	r.Version = ROSSUM_VERSION

	kpc.InitKPCs(r.RootDir)
}

func (r *RossumConfig) save(path string) {
	d, err := yaml.Marshal(r)
	if err != nil {
		log.Fatal("cannot yamalize config")
	}

	if err := ioutil.WriteFile(path, d, 0640); err != nil {
		log.Fatal("cannot write configuration into file")
	}
}
