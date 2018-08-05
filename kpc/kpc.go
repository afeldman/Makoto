// Copyright Anton Feldmann
//
// KPC is a Karel package managemant file.
// This structure is used to inform the tools about the structure of the project
//
package kpc

type Repo struct {
	Type string `yaml:"repo"`
	URL  string `yaml:"repo"`
}

type KPC struct {
	/********************             KPC          ********************/
	KPC_Version string `yaml:"kpc_version"` // version of the kpc

	/******************** Package information data ********************/
	Name         string        `yaml:"name"`         // project name
	Description  string        `yaml:"description"`  // project discription
	Version      string        `yaml:"version"`      // project version
	Homepage     string        `yaml:"url"`          // project homepage
	Requirements []Requirement `yaml:"requirements"` // dependnency list
	Conflicts    []Conflict    `yaml:"conflicts"`    // known conflicts
	Authors      []Author      `yaml:"author"`       // authorname
	Repo         Repo          `yaml:"source"`       // sorce code repository url
	Issues       string        `yaml:"issues"`       // project issue homepage

	/********************* Package path settings***********************/
	/*Package path settings*/
	Prefix     string `yaml:"prefix"`     // root path where the packge is installed
	SrcDir     string `yaml:"srcdir"`     // installation path to project sourcefiles (*.kl)
	TypeDir    string `yaml:"typedir"`    // installation path of project typefiles (*.t.kl)
	IncludeDir string `yaml:"includedir"` // installation path to project headerfiles (*.h.kl)
	ConstDir   string `yaml:"constdir"`   // installation path of the constant diclaraion files (*.c.kl)
	FormDir    string `yaml:"formdir"`    // if the package has a form to display content. the files might be in an directory
	DictDir    string `yaml:"dictdir"`    // dictionaries are available for the

	/*************** specific file includes ***************************/
	Main     string   `yaml:"main"`     // the source file to compile
	Dicts    []string `yaml:"dict"`     // dictionary file
	Forms    []string `yaml:"form"`     // form file
	Types    []string `yaml:"types"`    // the library for
	Includes []string `yaml:"includes"` // specific header files for comilation
	Consts   []string `yaml:"consts"`   // the const files of this project
}

func KPC_INIT(name string) *KPC {
	return &KPC{
		KPC_Version:  KPC_VERSION,
		Name:         name,
		Description:  "",
		Version:      "",
		Homepage:     "",
		Requirements: []Requirement{},
		Conflicts:    []Conflicts{},
		Authors:      []Author{},
		Repo:         {Type: "", URL: ""},
		Issus:        "",
		Prefix:       "",
		SrcDir:       "",
		TypeDir:      "",
		IncludeDir:   "",
		ConstDir:     "",
		Main:         "",
		Dists:        []string{},
		Includes:     []string{},
		Forms:        []string{},
		Types:        []string{},
		Consts:       []string{},
	}
}
