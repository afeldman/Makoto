// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

type KPC struct {
	/********************             KPC          ********************/
	KPC_Version string `json:"kpc_version"` // version of the kpc

	/******************** Package information data ********************/
	Name         string        `json:"name"`         // project name
	Description  string        `json:"description"`  // project discription
	Version      string        `json:"version"`      // project version
	Homepage     string        `json:"url"`          // project homepage
	Requirements []Requirement      `json:"requirements,omitempty"` // dependnency list
	Conflicts    []Conflict      `json:"conflicts,omitempty"`    // known conflicts
	Authors      []Author     `json:"author,omitempty"`       // authorname
	Repository   Repo          `json:"source,omitempty"`       // sorce code repository url
	Issue        string        `json:"issues"`       // project issue homepage

	/********************* Package path settings***********************/
	/*Package path settings*/
	Prefix     string `json:"prefix"`     // root path where the packge is installed
	SrcDir     string `json:"srcdir"`     // installation path to project sourcefiles (*.kl)
	TypeDir    string `json:"typedir"`    // installation path of project typefiles (*.t.kl)
	IncludeDir string `json:"includedir"` // installation path to project headerfiles (*.h.kl)
	ConstDir   string `json:"constdir"`   // installation path of the constant diclaraion files (*.c.kl)
	FormDir    string `json:"formdir"`    // if the package has a form to display content. the files might be in an directory
	DictDir    string `json:"dictdir"`    // dictionaries are available for the

	/*************** specific file includes ***************************/
	Main     string   `json:"main"`     // the source file to compile
	Dicts    []string `json:"dict,omitempty"`     // dictionary file
	Forms    []string `json:"form,omitempty"`     // form file
	Types    []string `json:"types,omitempty"`    // the library for
	Includes []string `json:"includes,omitempty"` // specific header files for comilation
	Consts   []string `json:"consts,omitempty"`   // the const files of this project
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
	}
}

func (this *KPC)GetKPCVersion() (string){
	return this.KPC_Version;
}
