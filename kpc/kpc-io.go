package kpc

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)


func (this *KPC) ToJSON() (error,[]byte){
	b, err := json.Marshal(this)
	if err != nil {
		return err, []byte{}
	}
	return nil, b
}

func FromJSON(json_str []byte) (error,*KPC) {
	var kpc KPC
	err := json.Unmarshal(json_str, &kpc)
	if err != nil {
		return err, nil
	}
	return nil, &kpc
}

func (this *KPC) ToYAML() (error, []byte){
	y, err := yaml.Marshal(this)
	if err != nil {
		return err, []byte{}
	}
	return nil, y
}

func FromYAML(yml_str []byte) (error,*KPC) {
	var kpc KPC
	err := yaml.Unmarshal(yml_str, &kpc)
	if err != nil {
		return err, nil
	}
	return nil, &kpc
}
