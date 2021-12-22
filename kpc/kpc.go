// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
//
package kpc

type KPC struct {
	/********************             KPC          ********************/
	KPC_Version string `toml:"kpc_version"` // version of the kpc

	/******************** Package information data ********************/
	Name         string        `toml:"name"`                   // project name
	Description  string        `toml:"description,omitempty"`  // project discription
	Version      string        `toml:"version"`                // project version
	Homepage     string        `toml:"url,omitempty"`          // project homepage
	Requirements []Requirement `toml:"requirements,omitempty"` // dependnency list
	Conflicts    []Conflict    `toml:"conflicts"`              // known conflicts
	Authors      []Author      `toml:"authors"`                // authorname
	Repository   Repository    `toml:"source"`                 // sorce code repository url
	Issue        string        `toml:"issues,omitempty"`       // project issue homepage
	Keywords     []string      `toml:"keywords,omitempty"`     // project issue homepage

	/********************* Package path settings***********************/
	/*Package path settings*/
	Prefix     string `toml:"prefix"`               // root path where the packge is installed
	SrcDir     string `toml:"srcdir,omitempty"`     // installation path to project sourcefiles (*.kl)
	TypeDir    string `toml:"typedir,omitempty"`    // installation path of project typefiles (*.t.kl)
	IncludeDir string `toml:"includedir,omitempty"` // installation path to project headerfiles (*.h.kl)
	ConstDir   string `toml:"constdir,omitempty"`   // installation path of the constant diclaraion files (*.c.kl)
	FormDir    string `toml:"formdir,omitempty"`    // the files might be in an directory (*.ftx)
	DictDir    string `toml:"dictdir,omitempty"`    // dictionaries are available for the (*.utx)

	/*************** specific file includes ***************************/
	Main     string   `toml:"main,omitempty"`     // the source file to compile
	Dicts    []string `toml:"dict,omitempty"`     // dictionary file
	Forms    []string `toml:"form,omitempty"`     // form file
	Types    []string `toml:"types,omitempty"`    // the library for
	Includes []string `toml:"includes,omitempty"` // specific header files for comilation
	Consts   []string `toml:"consts,omitempty"`   // the const files of this project
}

// initilize a KPC object
func InitKPC(name string) *KPC {
	return &KPC{
		KPC_Version:  KpcVersion,
		Name:         name,
		Description:  "",
		Version:      "",
		Homepage:     "",
		Repository:   *(InitRepository()),
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

// return KPC version version
func (kpc_obj *KPC) GetKPCVersion() string {
	return kpc_obj.KPC_Version
}
