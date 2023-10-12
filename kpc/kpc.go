// Copyright Anton Feldmann
//
// KPC is a Karel package management structure.
// This structure is used to inform the tools about the structure of the project
package kpc

import "fmt"

// kpc datastructurwe
type KPC struct {
	/********************             KPC          ********************/
	KPCVersion string `toml:"kpc_version"` // version of the kpc

	/******************** Package information data ********************/
	Name         string        `toml:"name"`                   // project name
	Description  string        `toml:"description"`            // project discription
	Version      string        `toml:"version"`                // project version
	Homepage     string        `toml:"url,omitempty"`          // project homepage
	Requirements []Requirement `toml:"requirements,omitempty"` // dependnency list
	Conflicts    []Conflict    `toml:"conflicts,omitempty"`    // known conflicts
	Authors      []Author      `toml:"authors,omitempty"`      // authorname
	Repository   Repository    `toml:"source,omitempty"`       // sorce code repository url
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
	Main     string   `toml:"main"`               // the source file to compile
	Dicts    []string `toml:"dicts,omitempty"`    // dictionary file
	Forms    []string `toml:"forms,omitempty"`    // form file
	Types    []string `toml:"types,omitempty"`    // the library for
	Includes []string `toml:"includes,omitempty"` // specific header files for comilation
	Consts   []string `toml:"consts,omitempty"`   // the const files of this project
}

// initilize a KPC object
func InitKPC(name string) *KPC {
	return &KPC{
		KPCVersion:   KpcVersion,
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
//
// @return (string) version
func (kpc_obj *KPC) GetKPCVersion() string {
	return kpc_obj.KPCVersion
}

// RejectConflict you did well and fixed an issue. then delete the conflict
//
// @param (string) name of conflivt package
func (kpc_obj *KPC) RejectConflict(name string) {
	var tmp []Conflict
	for _, v := range kpc_obj.Conflicts {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Conflicts = tmp
}

// ConflictSize how many conflicts are known
func (kpc_obj *KPC) ConflictSize() int {
	return len(kpc_obj.Conflicts)
}

// AddConflict you found a new conflict so add the conflict to the list
func (kpc_obj *KPC) AddConflict(con Conflict) {
	kpc_obj.Conflicts = append(kpc_obj.Conflicts, con)
}

// GetConflicts get conflicts
func (kpc_obj *KPC) GetConflicts() []Conflict {
	return kpc_obj.Conflicts
}

// GetConflict get a specific conflict
func (kpc_obj *KPC) GetConflict(name string) *Conflict {
	for _, val := range kpc_obj.Conflicts {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

// GetAuthor is the author structure of the KPC struct
//
// @params (strung) the authors name
//
// @return (*Author) the author structure
// @return (error) if the author is unknown throws error
func (kpc_obj *KPC) GetAuthor(authorsName string) (*Author, error) {
	for _, val := range kpc_obj.Authors {
		if val.Name == authorsName {
			return &val, nil
		}
	}
	return nil, fmt.Errorf("can not find author %s", authorsName)
}

// AddAuthor add a new Author to the list
//
// @param (Author) an author struct
func (kpc_obj *KPC) AddAuthor(aut Author) {
	_, err := kpc_obj.GetAuthor(*aut.GetName())
	if err != nil {
		// author do not exists in list
		kpc_obj.Authors = append(kpc_obj.Authors, aut)
	}
}

// NumberOfAuthors is the number of authors in the author list
//
// @return (int) the number of authors in kpc
func (kpc_obj *KPC) NumberOfAuthors() int {
	return len(kpc_obj.Authors)
}

// RejectAuthor deletes an author in the list of authors
//
// @param (string) authors name
func (kpc_obj *KPC) RejectAuthor(name string) {
	var tmp []Author
	for _, v := range kpc_obj.Authors {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Authors = tmp
}

// get all rewuirements
func (kpc_obj *KPC) GetRequirement(name string) *Requirement {
	for _, val := range kpc_obj.Requirements {
		if val.Name == name {
			return &val
		}
	}
	return nil
}

// add a requirement
func (kpc_obj *KPC) AddRequirement(req Requirement) {
	kpc_obj.Requirements = append(kpc_obj.Requirements, req)
}

// how many requirements are stored
func (kpc_obj *KPC) RequirementSize() int {
	return len(kpc_obj.Requirements)
}

// delete a requirement
func (kpc_obj *KPC) RejectRequirement(name string) {
	var tmp []Requirement
	for _, v := range kpc_obj.Requirements {
		if v.Name == name {
			continue
		} else {
			tmp = append(tmp, v)
		}
	}
	kpc_obj.Requirements = tmp
}
