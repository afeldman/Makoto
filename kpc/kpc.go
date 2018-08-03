// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

type KPC struct {
	/********************             KPC          ********************/
	KPC_VERSION  string        `yaml:"kpc_version"`  // version of the kpc

	/******************** Package information data ********************/
	Name         string        `yaml:"name"`         // project name
	Description  string        `yaml:"description"`  // project discription
	Version      string        `yaml:"version"`      // project version
	Url          string        `yaml:"url"`          // project url
	Requirements []Requirement `yaml:"requirements"` // dependnency list
	Conflicts    []Conflict    `yaml:"conflicts"`    // known conflicts
	Authors      []Author      `yaml:"author"`       // authorname
	Source       string        `yaml:"source"`       // sorce url

	/********************* Package path settings***********************/
	/*Package path settings*/
	Prefix     string `yaml:"prefix"`     // root path where the packge is installed
	SrcDir     string `yaml:"srcdir"`     // installation path to project sourcefiles (*.kl)
	TypeDir    string `yaml:"typedir"`    // installation path of project typefiles (*.klt)
	IncludeDir string `yaml:"includedir"` // installation path to project headerfiles (*klh)
	ConstDir   string `yaml:"constdir"`   // installation path of the constant diclaraion files (*klc)

	/*************** specific file includes ***************************/
	main     string   `yaml:"main"`     // the source file to compile
	Dict     []string `yaml:"dict"`     // dictionary file
	Form     []string `yaml:"form"`     // form file
	Types    []string `yaml:"types"`    // the library for
	Includes []string `yaml:"includes"` // specific header files for comilation
	Consts   []string `yaml:"consts"`   // the const files of this project
}

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
		path = os.Getenv("KPC_HOME")
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
