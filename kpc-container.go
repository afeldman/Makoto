package kpc

import (
	"fmt"
	"log"
	"path/filepath"
)

var KPCs map[string]KPC

func InitKPCs(basic_path []string) {

	KPCs = make(map[string]KPC)

	for _, dir := range basic_path {
		files, err := filepath.Glob(dir + "*.kpc")
		if err != nil {
			log.Fatalln(err)
		}

		for _, kpc_file := range files {
			kcpt := FromYAML(kpc_file)
			KPCs[kcpt.Name] = *kcpt
		}
	}
}

func KPC_Size() int {
	return len(KPCs)
}

func All() string {
	var names string

	for kpc_key, _ := range KPCs {
		names += fmt.Sprintf("%s\n", kpc_key)
	}

	return names
}

func ContainsPackage(comp string) bool {
	if _, ok := KPCs[comp]; ok {
		return true
	}
	return false
}

func GetKPC(kpg_name string) *KPC {
	if ContainsPackage(kpg_name) {
		kpc := KPCs[kpg_name]
		return (&kpc)
	} else {
		return nil
	}
}
