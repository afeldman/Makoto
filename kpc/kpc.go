// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

type KPC struct {
	/********************             KPC          ********************/
	KPC_Version string `json:"kpc_version" yaml:"kpc_version"` // version of the kpc

	/******************** Package information data ********************/
	Name         string        `json:"name" yaml:"name"`         // project name
	Description  string        `json:"description,omitempty" yaml:"description,omitempty"`  // project discription
	Version      string        `json:"version" yaml:"version"`      // project version
	Homepage     string        `json:"url,omitempty" yaml:"url,omitempty"`          // project homepage
	Requirements []Requirement      `json:"requirements,omitempty" yaml:"requirements,omitempty"` // dependnency list
	Conflicts    []Conflict      `json:"conflicts,omitempty" yaml:"conflicts,omitempty"`    // known conflicts
	Authors      []Author     `json:"author" yaml:"author"`       // authorname
	Repository   Repo          `json:"source" yaml:"source"`       // sorce code repository url
	Issue        string        `json:"issues,omitempty" yaml:"issues,omitempty"`       // project issue homepage
	Keywords    []string      `json:"keywords,omitempty" yaml:"keywords,omitempty"`       // project issue homepage

	/********************* Package path settings***********************/
	/*Package path settings*/
	Prefix     string `json:"prefix" yaml:"prefix"`     // root path where the packge is installed
	SrcDir     string `json:"srcdir,omitempty" yaml:"srcdir,omitempty"`     // installation path to project sourcefiles (*.kl)
	TypeDir    string `json:"typedir,omitempty" yaml:"typedir,omitempty"`    // installation path of project typefiles (*.t.kl)
	IncludeDir string `json:"includedir,omitempty" yaml:"includedir,omitempty"` // installation path to project headerfiles (*.h.kl)
	ConstDir   string `json:"constdir,omitempty" yaml:"constdir,omitempty"`   // installation path of the constant diclaraion files (*.c.kl)
	FormDir    string `json:"formdir,omitempty" yaml:"formdir,omitempty"`    // if the package has a form to display content. the files might be in an directory
	DictDir    string `json:"dictdir,omitempty" yaml:"dictdir,omitempty"`    // dictionaries are available for the

	/*************** specific file includes ***************************/
	Main     string   `json:"main" yaml:"main"`     // the source file to compile
	Dicts    []string `json:"dict,omitempty" yaml:"dict,omitempty"`     // dictionary file
	Forms    []string `json:"form,omitempty" yaml:"form,omitempty"`     // form file
	Types    []string `json:"types,omitempty" yaml:"types,omitempty"`    // the library for
	Includes []string `json:"includes,omitempty" yaml:"includes,omitempty"` // specific header files for comilation
	Consts   []string `json:"consts,omitempty" yaml:"consts,omitempty"`   // the const files of this project
}

func KPC_Init(name string) (*KPC) {
	return &KPC{
		KPC_Version:  KPC_VERSION,
		Name:         name,
		Description:  "",
		Version:      "",
		Homepage:     "",
		Repository:   Repo{Type: "", URL: ""},
		Issue:        "",
		Prefix:       "",
		SrcDir:       "",
		TypeDir:      "",
		IncludeDir:   "",
		ConstDir:     "",
		Main:         "",
		Requirements: []Requirement{},
		Conflicts:    []Conflict{},
		Authors:      []Author{},
		Dicts:        []string{},
		Includes:     []string{},
		Forms:        []string{},
		Types:        []string{},
		Consts:       []string{},
		Keywords:     []string{},
	}
}

func (this *KPC)GetKPCVersion() (string){
	return this.KPC_Version;
}
