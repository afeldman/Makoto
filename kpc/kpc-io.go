package kpc

import (
	"encoding/json"
	"log"
	"gopkg.in/yaml.v2"
)


func (this *KPC) ToJSON() ([]byte){
	b, err := json.Marshal(this)
	if err != nil {
		log.Fatal(err)
	}
	return b
}

func FromJSON(json_str []byte) *KPC {
	var kpc KPC
	err := json.Unmarshal(json_str, &kpc)
	if err != nil {
		log.Fatal(err)
	}
	return &kpc
}

func (this *KPC) ToYAML() ([]byte){
	y, err := yaml.Marshal(this)
	if err != nil {
		log.Fatal(err)
	}
	return y
}

func FromYAML(yml_str []byte) *KPC {
	var kpc KPC
	err := yaml.Unmarshal(yml_str, &kpc)
	if err != nil {
		log.Fatal(err)
	}
	return &kpc
}
