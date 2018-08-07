package kpc

import (
	"encoding/json"
	"log"
)


func (this *KPC) ToJSON() ([]byte){
	b, err := json.Marshal(this)
	if err != nil {
		log.Fatal(err);
	}
	return b
}

func FromJSON(json_str string) *KPC {


	return &KPC{}
}
