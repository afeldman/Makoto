package kpc

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)


func (this *KPC) ToYAML(path string) {

	if path == "" {
		path = os.Getenv("KPC_HOME")
	}

	d, e := yaml.Marshal(this)
	if e != nil {
		panic(e)
	}
	err := ioutil.WriteFile(path, d, 0644)
	if err != nil {
		panic(err)
	}
}

func FromYAML(path string) *KPC {

	if path == "" {
		log.Fatal("No path, you have to read the file");
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	kpc := KPC{}

	err = yaml.Unmarshal([]byte(file), &kpc)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return &kpc
}
